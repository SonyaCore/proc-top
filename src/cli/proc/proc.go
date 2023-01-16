package proc

import (
	"fmt"
	"log"
	"proc-top/utils/colors"

	"github.com/shirou/gopsutil/v3/process"
)

var (
	g = colors.Green
	r = colors.Reset
)

func Proc() {
	proc, err := process.Processes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g, proc, r)
}
