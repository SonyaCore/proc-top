package host

import (
	"fmt"
	"log"
	"proc-top/utils/colors"
	"strconv"

	"github.com/shirou/gopsutil/load"
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
	fmt.Println(p, "*KERNEL*", r)
	fmt.Printf("%v\n%v\n%v\n%v\n", "Platform : "+g+platfrom+r, "Version : "+g+kversion+r, "Arch : "+g+karch+r, "Uptime : "+g+time+r)
}

func Sensors() {
	stat, err := host.SensorsTemperatures()
	if err != nil {
		log.Fatal(err)
	}
	statdict := stat[0]
	model := statdict.SensorKey
	temp := statdict.Temperature
	high := statdict.High
	critical := statdict.Critical

	fmt.Println(p, "*SENSORS*", r)
	fmt.Println("Model :", g, model, r)
	fmt.Println("Temp :", g, temp, r)
	fmt.Println("High :", g, high, r)
	fmt.Println("Critical :", g, critical, r)
}

func Loadaverage() {

	load, err := load.Avg()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p, "*LOAD AVERAGE*", r)
	fmt.Printf("1 min ave: %v%.2f%v\n", g, load.Load1, r)
	fmt.Printf("5 min ave: %v%.2f%v\n", g, load.Load5, r)
	fmt.Printf("15 min ave: %v%.2f%v\n", g, load.Load15, r)

}
