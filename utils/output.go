package utils

import (
	"CLIMonitoring/internal/bar"
	"CLIMonitoring/internal/cpu"
	"CLIMonitoring/internal/disk"
	"CLIMonitoring/internal/gpu"
	"CLIMonitoring/internal/ram"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func Output() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nStopped monitoring.")
			return
		default:
			usgcpu, err := cpu.GetCPUUsage()
			if err != nil {
				fmt.Println(err)
				return
			}
			usgmem, err := ram.GetRAMUsage()
			if err != nil {
				fmt.Println(err)
				return
			}
			usgdisk, err := disk.GetDiskSpace()
			if err != nil {
				fmt.Println(err)
				return
			}
			usggpu, err := gpu.GetGPUUsage()
			if err != nil {
				fmt.Println(err)
				return
			}
			tempgpu, err := gpu.GetGPUTemperature()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf(
				"\rCPU: %s || RAM: %s || DISK: %s || GPU: %s || %d°C  ",
				bar.ProgressBar(usgcpu),
				bar.ProgressBar(usgmem),
				bar.ProgressBar(usgdisk),
				bar.ProgressBar(usggpu),
				tempgpu,
			)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
