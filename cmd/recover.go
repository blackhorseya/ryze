package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(recoverCmd)
}

var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Recover the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Recovering blockchain...")
		// TODO: 2024/9/14|sean|recover the blockchain
	},
}
