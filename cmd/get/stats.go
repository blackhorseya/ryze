package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get stats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stats")
	},
}
