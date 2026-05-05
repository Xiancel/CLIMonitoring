package utils

import (
	"CLIMonitoring/internal/bar"
	"CLIMonitoring/internal/cpu"
	"CLIMonitoring/internal/disk"
	"CLIMonitoring/internal/gpu"
	"CLIMonitoring/internal/ram"
	"CLIMonitoring/internal/stats"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func Output() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	s := stats.NewStats()

	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	for {
		select {
		case <-ctx.Done():
			fmt.Print("\033[2J\033[H")

			fmt.Println("+-------------------------------+")
			fmt.Println("| Stop monitoring...            | ")
			fmt.Println("| Thank for using CLIMonitoring | ")
			fmt.Printf("| Uptime: %s                    |\n", s.Uptime().Truncate(time.Second))
			fmt.Printf("| Max CPU: %.2f%%               |\n", s.MaxCPU)
			fmt.Printf("| Min CPU: %.2f%%                |\n", s.MinCPU)
			fmt.Println("+-------------------------------+")

			time.Sleep(500 * time.Millisecond)
			fmt.Print("\033[?25h")
			return
		default:
			fmt.Print("\033[H")

			usgcpu, err := cpu.GetCPUUsage()
			if err != nil {
				fmt.Println(err)
				return
			}
			s.Update(usgcpu)
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

			fmt.Println("+------------------------------------------------+")
			fmt.Printf("| CPU   %s             |\n", bar.ProgressBar(usgcpu))
			fmt.Printf("| RAM   %s            |\n", bar.ProgressBar(usgmem))
			fmt.Printf("| DISK  %s            |\n", bar.ProgressBar(usgdisk))
			fmt.Printf("| GPU   %s || %d°C     |\n", bar.ProgressBar(usggpu), tempgpu)
			fmt.Println("+------------------------------------------------+")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
