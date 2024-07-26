package get

import (
	"github.com/spf13/cobra"
)

// Cmd GetCmd represents the get command
var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources",
}

func init() {
	Cmd.AddCommand(blocksCmd)
}
