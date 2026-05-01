package gpu

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetGPUUsage() (float64, error) {
	cmd := exec.Command(
		"nvidia-smi",
		"--query-gpu=utilization.gpu",
		"--format=csv,noheader,nounits",
	)
	call, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	gpuUsage := strings.TrimSpace(string(call))

	value, err := strconv.ParseFloat(gpuUsage, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func GetGPUTemperature() (int, error) {
	cmd := exec.Command(
		"nvidia-smi",
		"--query-gpu=temperature.gpu",
		"--format=csv,noheader,nounits",
	)

	call, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	tempStr := strings.TrimSpace(string(call))

	value, err := strconv.Atoi(tempStr)
	if err != nil {
		return 0, err	
	}

	return value, nil
}
