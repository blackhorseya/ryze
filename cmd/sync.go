package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Syncing Ton blockchain...")
		// TODO: 2024/9/14|sean|sync the blockchain
	},
}
