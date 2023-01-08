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
	"time"
)

var (
	version  = "0.2.0"
	build    = "Custom"
	codename = "ProcTop , CLI monitor tool."
)

func Version() string {
	return version
}

// VersionStatement returns a list of strings representing the full version info.
func VersionStatement() []string {
	return []string{
		"ProcTop ", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")",
	}
}

func banner() string {
	return "PROC TOP"
}

func main() {
	// flags declaration using flag package
	kernel := flag.Bool("kernel", true, "Show kernel info & uptime")
	memoryflag := flag.Bool("memory", false, "Show memory usage")
	sensorsflag := flag.Bool("sensors", false, "Show sensors")
	cpuflag := flag.Bool("cpu", true, "Show Cpu info")
	diskflag := flag.Bool("disk", false, "Show disk usage")
	full := flag.Bool("full", false, "Show all information")

	flag.Parse()
	for {
		utils.CallClear()
		fmt.Println(banner())
		if *full {
			host.KernelInfo()
			memory.Memory()
			cpu.Cpu()
			host.Sensors()
			disk.Disk()
		} else {
			if *memoryflag {
				memory.Memory()
			}
			// proc.Proc()
			if *kernel {
				host.KernelInfo()
			}
			if *sensorsflag {
				host.Sensors()
			}
			if *cpuflag {
				cpu.Cpu()
			}
			if *diskflag {
				disk.Disk()
			}
		}
		time.Sleep(1 * time.Second)
	}
}
