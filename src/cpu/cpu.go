package cpu

import (
	"fmt"
	"log"
	"proc-top/colors"

	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	g = colors.Green
	r = colors.Reset
	p = colors.Purple
)

func Cpu() {
	cpuinfo, err := cpu.Info()

	if err != nil {
		log.Fatal(err)
	}

	cpudict := cpuinfo[0]

	fmt.Println(p, "*CPU INFO*", r)
	fmt.Printf("Model : %v%v%v\nCores : %v%v%v\n", g, cpudict.ModelName, r, g, cpudict.Cores, r)
}
