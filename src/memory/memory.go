package memory

import (
	"fmt"
	"log"
	"proc-top/colors"
	"proc-top/utils"

	"github.com/shirou/gopsutil/v3/mem"
)

var (
	g      = colors.Green
	r      = colors.Reset
	p      = colors.Purple
	v, err = mem.VirtualMemory()
)

func Memory() {

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p, "*MEMORY*", r)
	fmt.Printf("Total: %v%v%v\nFree: %v%v%v\nUsedPercent: %v%.2f%%%v\n",
		g, utils.ConvByte(float64(v.Total)), r, g, utils.ConvByte(float64(v.Available)), r, g, v.UsedPercent, r)

}

func Swap() {

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p, "*SWAP*", r)
	fmt.Printf("Cached : %v%v%v \nFree : %v%v%v \nTotal : %v%v%v \n", g, v.SwapCached, r, g, v.SwapFree, r, g, v.SwapTotal, r)
}
