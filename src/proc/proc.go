package proc

import (
	"fmt"
	"proc-top/colors"

	"github.com/shirou/gopsutil/v3/process"
)

var (
	g = colors.Green
	r = colors.Reset
	p = colors.Purple
)

func Proc() {
	proc, _ := process.Processes()
	fmt.Println(g, proc, r)
}
