package cpu

import (
	"fmt"
	"proc-top/colors"

	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	g = colors.Green
	r = colors.Reset
	p = colors.Purple
)

func Cpu() {
	cpuinfo, _ := cpu.Info()
	cpudict := cpuinfo[0]

	fmt.Println(p, "*CPU INFO*", r)
	fmt.Printf("Model : %v%v%v\nCores : %v%v%v\n", g, cpudict.ModelName, r, g, cpudict.Cores, r)
}
