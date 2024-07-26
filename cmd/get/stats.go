package get

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get stats",
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

		api, err := tonx.NewAPIClient(tonx.Options{Network: network})
		cobra.CheckErr(err)

		info, err := api.GetMasterchainInfo(contextx.Background())
		cobra.CheckErr(err)

		// display info
		cmd.Println("Masterchain Info:")
		cmd.Printf("  Height: %d\n", info.SeqNo)
	},
}
