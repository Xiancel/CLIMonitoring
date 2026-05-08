package cmd

import (
	"CLIMonitoring/utils"

	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start monitoring...",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Output()
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
