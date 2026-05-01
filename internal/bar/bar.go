package bar

import "fmt"

func ProgressBar(percent float64) string {
	totalBars := 10
	filledBars := percent * float64(totalBars) / 100

	bar := ""

	for i := 0; i < int(filledBars); i++ {
		bar += "█"
	}

	for i := filledBars; i < float64(totalBars); i++ {
		bar += "-"
	}

	return fmt.Sprintf("[%s] %.2f%%", bar, percent)
}
