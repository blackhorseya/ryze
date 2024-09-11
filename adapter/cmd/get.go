package cmd

import (
	"github.com/blackhorseya/ryze/adapter/cmd/get"
)

func init() {
	rootCmd.AddCommand(get.Cmd)
}
