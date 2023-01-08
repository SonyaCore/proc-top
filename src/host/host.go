package host

import (
	"fmt"
	"log"
	"proc-top/colors"
	"strconv"

	"github.com/shirou/gopsutil/v3/host"
)

var (
	g = colors.Green
	r = colors.Reset
	p = colors.Purple
)

func KernelInfo() {
	kversion, _ := host.KernelVersion()
	karch, _ := host.KernelArch()
	utime, _ := host.Uptime()
	platfrom, _, _, _ := host.PlatformInformation()
	var time string

	if utime >= 120 {
		utime = utime / 60
		time = strconv.FormatUint(utime, 10) + " min"
	} else {
		time = strconv.FormatUint(utime, 10)
	}
	fmt.Println(p, "KERNEL", r)
	fmt.Printf("%v\n%v\n%v\n%v\n", g+platfrom+r, g+kversion+r, g+karch+r, g+time+r)
}

func Sensors() {
	stat, err := host.SensorsTemperatures()
	if err != nil {
		log.Fatal(err)
	}
	statdict := stat[0]

	fmt.Println(p, "SENSORS", r)
	fmt.Printf("%v", statdict.String())
	fmt.Printf("%v\n", statdict.SensorKey)
}
