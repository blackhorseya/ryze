package cmd

import (
	blockGRPC "github.com/blackhorseya/ryze/adapter/block/grpc"
	platformGRPC "github.com/blackhorseya/ryze/adapter/platform/grpc"
	"github.com/blackhorseya/ryze/pkg/cmdx"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the service",
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.AddCommand(cmdx.NewServiceCmd("platform-grpc", "start platform grpc service", platformGRPC.New))
	startCmd.AddCommand(cmdx.NewServiceCmd("block-grpc", "start block grpc service", blockGRPC.New))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
