package otel

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func getMemoryUsage() float64 {
	v, _ := mem.VirtualMemory()
	return v.UsedPercent
}

func getCPUUsage() float64 {
	c, _ := cpu.Percent(time.Second, false)
	if len(c) > 0 {
		return c[0]
	}
	return 0
}
