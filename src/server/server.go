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

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
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

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowCredentials: true,
		Debug:            false,
	})

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
	handler := cors.Handler(router)
	fmt.Println("Listening on port", PORT)

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), handler)
	if err != nil {
		log.Fatal(err)
	}
}

var stats = make(map[string]interface{})

func update() {
	platfrom, _, _, _ := host.PlatformInformation()
	kernel_version, _ := host.KernelVersion()
	kernel_arch, _ := host.KernelArch()
	platfrominfo := fmt.Sprintf("%v %v %v", platfrom, kernel_version, kernel_arch)

	stats["kernel"] = platfrominfo
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
		load, _ := load.Avg()
		loadstatus := fmt.Sprintf("%.2f %.2f %.2f", load.Load1, load.Load5, load.Load15)

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
		stats["loadaverage"] = loadstatus
		time.Sleep(1 * time.Second)
	}
}
