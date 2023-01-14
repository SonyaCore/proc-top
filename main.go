package main

import (
	"flag"
	"fmt"
	"os"
	"proc-top/colors"
	"proc-top/src/host"
	"proc-top/utils"
	"runtime"
	"strings"
	"time"

	cpu "proc-top/src/cpu"
	disk "proc-top/src/disk"
	memory "proc-top/src/memory"
	procserver "proc-top/src/server"
)

var (
	name     = "ProcTop"
	version  = "0.6.0"
	build    = "Custom"
	codename = "ProcTop , System monitor tool."
)

var (
	Purple = colors.Purple
	res    = colors.Reset
)

const (
	banner = `     

  ______                 _____             
  | ___ \               |_   _|            
  | |_/ / __ ___   ___    | | ___  _ __    
  |  __/ '__/ _ \ / __|   | |/ _ \| '_ \   
  | |  | | | (_) | (__    | | (_) | |_) |  
  \_|  |_|  \___/ \___|   \_/\___/| .__/   
                                  | |      
                                  |_|      

`
)

func Header() {
	fmt.Println(VersionStatement())
	fmt.Println(Purple, banner, res)
}

func Version() string {
	return version
}

// VersionStatement returns a list of strings representing the full version info.
func VersionStatement() string {
	return strings.Join([]string{
		"ProcTop ", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")",
	}, "")
}

func main() {
	server := flag.Bool("server", false, "start web mode (default mode)")
	cli := flag.Bool("cli", false, "start cli mode")
	interval := flag.Int("interval", 1, "refresh screen per second")
	port := flag.Int("port", 8080, "webserver port. ")
	versionflag := flag.Bool("version", false, "Show version & exit")

	flag.Parse()

	if *versionflag {
		fmt.Println(name, version)
		os.Exit(0)
	}
	if *server {
		Header()
		procserver.Start(*port)

	} else if *cli {
		for {
			utils.CallClear()
			Header()
			host.KernelInfo()
			memory.Memory()
			memory.Swap()
			cpu.Cpu()
			host.Sensors()
			host.Loadaverage()
			disk.Disk()
			time.Sleep(time.Duration(*interval) * time.Second)

		}
	} else {
		Header()
		procserver.Start(*port)
	}
}
