package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

//go:embed public
var content embed.FS

func getFileSystem(OSDir string) http.FileSystem {
	if OSDir == "live" {
		log.Print("using live mode")
		return http.FS(os.DirFS("public"))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(content, "public")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func Start(port int) {
	go update()

	router := mux.NewRouter()

	router.HandleFunc("/raw", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(stats)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}).Methods("GET")

	OSDir := "embed"
	router.PathPrefix("/").Handler(http.FileServer(getFileSystem(OSDir)))

	PORT := port
	fmt.Println("Listening on port", PORT)
	http.ListenAndServe(":"+strconv.Itoa(PORT), handlers.CORS()(router))
}

var stats = make(map[string]interface{})

func update() {
	for {
		diskused := uint64(0)
		disktotal := uint64(0)
		diskinfo, _ := disk.Usage("/")

		diskused += diskinfo.Used
		disktotal += diskinfo.Total

		uptime, _ := host.Uptime()
		cpuusage, _ := cpu.Percent(0, true)
		cpudict := cpuusage[0]
		vm, _ := mem.VirtualMemory()
		dio, _ := disk.IOCounters()
		diskdict := dio["disk0"]
		nio, _ := net.IOCounters(false)
		niodict := nio[0]
		swap, _ := mem.SwapMemory()

		stats["uptime"] = uptime
		stats["fqdn"], _ = os.Hostname()
		stats["cpuusage"] = cpudict
		stats["ramusage"] = [6]uint64{vm.Total, vm.Available,
			uint64(vm.UsedPercent),
			vm.Used, vm.Free}
		stats["diskio"] = [6]uint64{diskdict.ReadCount, diskdict.WriteCount,
			diskdict.ReadBytes, diskdict.WriteBytes,
			diskdict.ReadTime, diskdict.WriteTime}
		stats["diskusage"] = [2]uint64{diskused, disktotal}
		stats["netio"] = [8]uint64{niodict.BytesSent,
			niodict.BytesRecv, niodict.PacketsSent,
			niodict.PacketsRecv, niodict.Errin,
			niodict.Errout, niodict.Dropout, niodict.Dropout}
		stats["swapusage"] = [6]uint64{swap.Total, swap.Used,
			swap.Free, uint64(swap.UsedPercent),
			swap.Sin, swap.Sout}
		time.Sleep(1 * time.Second)
	}
}
