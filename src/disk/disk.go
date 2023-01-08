package disk

import (
	"fmt"
	"log"
	"proc-top/colors"
	"proc-top/utils"

	"github.com/shirou/gopsutil/v3/disk"
)

var (
	g = colors.Green
	r = colors.Reset
	p = colors.Purple
)

func Disk() {
	diskinfo, err := disk.Usage("/")

	if err != nil {
		log.Fatal(err)
	}
	fstype := diskinfo.Fstype
	total := float64(diskinfo.Total)
	used := float64(diskinfo.Used)
	free := float64(diskinfo.Free)

	fmt.Println(p, "*DISK INFO*", r)
	fmt.Println("Type :", g, fstype, r)
	fmt.Println("Total :", g, utils.ConvByte(total), r)
	fmt.Println("Used :", g, utils.ConvByte(used), r)
	fmt.Println("Free :", g, utils.ConvByte(free), r)
}
