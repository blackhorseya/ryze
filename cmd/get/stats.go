package get

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get stats",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configx.NewConfiguration(viper.GetViper())
		cobra.CheckErr(err)

		network, ok := config.Networks["ton"]
		if !ok {
			cmd.PrintErr("ton network not found")
			return
		}

		fmt.Println(network)
	},
}
