package cpu

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPUUsage() (float64, error) {
	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return cpuUsage[0], nil
}
