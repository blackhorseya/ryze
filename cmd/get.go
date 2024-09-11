package cmd

import (
	"github.com/blackhorseya/ryze/cmd/get"
)

func init() {
	rootCmd.AddCommand(get.Cmd)
}
