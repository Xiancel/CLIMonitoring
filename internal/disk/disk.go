package disk

import "github.com/shirou/gopsutil/v3/disk"

func GetDiskSpace() (float64, error) {
	disk, err := disk.Usage("C:\\")
	if err != nil {
		return 0, err
	}

	return disk.UsedPercent, nil
}
