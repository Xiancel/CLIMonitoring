package ram

import "github.com/shirou/gopsutil/v3/mem"

func GetRAMUsage() (float64, error) {
	ramUsage, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return ramUsage.UsedPercent, nil

}
