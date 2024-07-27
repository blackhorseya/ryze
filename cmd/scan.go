package cmd

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	startFlag int
	endFlag   int
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configx.NewConfiguration(viper.GetViper())
		cobra.CheckErr(err)

		tonConfig, ok := config.Networks["ton"]
		if !ok {
			cmd.PrintErr("ton network not found")
			return
		}

		network := "mainnet"
		if tonConfig.Testnet {
			network = "testnet"
		}

		cmd.Printf("Scanning Ton blockchain Network: %s...\n", network)

		var start, end int
		if startFlag == 0 {
			// TODO: 2024/7/27|sean|get start block from blockchain
			// start = tonConfig.StartHeight
		} else {
			start = startFlag
		}

		if endFlag == 0 {
			// TODO: 2024/7/27|sean|get current block from blockchain
			// end = tonConfig.CurrentHeight
		} else {
			end = endFlag
		}

		cmd.Printf("Scanning from block %d to block %d\n", start, end)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().IntVar(&startFlag, "start", 0, "The start block number")
	scanCmd.Flags().IntVar(&endFlag, "end", 0, "The end block number")
}
