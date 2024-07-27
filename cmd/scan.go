package cmd

import (
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("scan called")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
