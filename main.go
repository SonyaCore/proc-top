package main

import (
	"fmt"
	cpu "proc-top/src/cpu"
	disk "proc-top/src/disk"
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
	for {
		utils.CallClear()
		fmt.Println(banner())
		memory.Memory()
		cpu.Cpu()
		disk.Disk()
		time.Sleep(1 * time.Second)
	}
}
