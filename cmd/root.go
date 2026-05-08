package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "climon",
	Short: "CLI monitoring system",
	Run:   welcomeMess,
}

func welcomeMess(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to CLI Monitoring! Use --help for commands")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
