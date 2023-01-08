package main

import (
	"flag"
	"fmt"
	cpu "proc-top/src/cpu"
	disk "proc-top/src/disk"
	"proc-top/src/host"
	memory "proc-top/src/memory"
	"proc-top/utils"
	"runtime"
	"strings"
	"time"
)

var (
	name     = "ProcTop"
	version  = "0.4.0"
	build    = "Custom"
	codename = "ProcTop , CLI monitor tool."
)

func Version() string {
	return version
}

// VersionStatement returns a list of strings representing the full version info.
func VersionStatement() string {
	return strings.Join([]string{
		"ProcTop ", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")",
	}, "")
}

func banner() string {
	return "PROC TOP"
}

func main() {
	// flags declaration using flag package
	interval := flag.Int("interval", 1, "refresh screen per second")
	full := flag.Bool("full", false, "Show all information")
	kernel := flag.Bool("kernel", true, "Show kernel info & uptime")
	memoryflag := flag.Bool("memory", false, "Show memory usage")
	swapflag := flag.Bool("swap", false, "Show swap usage")
	loadaverageflag := flag.Bool("load", false, "Show load average")
	sensorsflag := flag.Bool("sensors", false, "Show sensors")
	cpuflag := flag.Bool("cpu", true, "Show Cpu info")
	diskflag := flag.Bool("disk", false, "Show disk usage")
	versionflag := flag.Bool("version", false, "Show version & exit")

	flag.Parse()
	for {
		if *versionflag {
			fmt.Println(name, version)
			break
		}
		utils.CallClear()
		fmt.Println(VersionStatement())
		if *full {
			host.KernelInfo()
			memory.Memory()
			memory.Swap()
			cpu.Cpu()
			host.Sensors()
			host.Loadaverage()
			disk.Disk()
		} else {
			if *memoryflag {
				memory.Memory()
			}
			if *swapflag {
				memory.Swap()
			}
			if *kernel {
				host.KernelInfo()
			}
			if *sensorsflag {
				host.Sensors()
			}
			if *loadaverageflag {
				host.Loadaverage()
			}
			if *cpuflag {
				cpu.Cpu()
			}
			if *diskflag {
				disk.Disk()
			}
		}
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}
