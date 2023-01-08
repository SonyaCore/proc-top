package memory

import (
	"fmt"
	"proc-top/colors"
	"proc-top/utils"

	"github.com/shirou/gopsutil/v3/mem"
)

var (
	g = colors.Green
	r = colors.Reset
	p = colors.Purple
)

func Memory() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Println(p, "*MEMORY*", r)
	fmt.Printf("Total: %v%v%v\nFree: %v%v%v\nUsedPercent: %v%.2f%%%v\n",
		g, utils.ConvByte(float64(v.Total)), r, g, utils.ConvByte(float64(v.Available)), r, g, v.UsedPercent, r)

}
