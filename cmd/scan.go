package cmd

import (
	"context"
	"fmt"

	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
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

		client, err := tonx.NewClient(tonx.Options{Network: network})
		cobra.CheckErr(err)

		api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
		api.SetTrustedBlockFromConfig(client.Config)

		cmd.Printf("Scanning Ton blockchain Network: %s...\n", network)
		master, err := api.GetMasterchainInfo(contextx.Background())
		cobra.CheckErr(err)

		cmd.Println("master proofs chain successfully verified, all data is now safe and trusted!")

		{
			ctx := api.Client().StickyContext(contextx.Background())
			shardLastSeqno := map[string]uint32{}
			firstShards, err2 := api.GetBlockShardsInfo(ctx, master)
			cobra.CheckErr(err2)

			for _, shard := range firstShards {
				shardLastSeqno[getShardID(shard)] = shard.SeqNo
			}

			for {
				cmd.Printf("Scanning %d master block...\n", master.SeqNo)

				currentShards, err3 := api.GetBlockShardsInfo(ctx, master)
				cobra.CheckErr(err3)

				var newShards []*ton.BlockIDExt
				for _, shard := range currentShards {
					notSeen, err4 := getNotSeenShards(ctx, api, shard, shardLastSeqno)
					cobra.CheckErr(err4)

					shardLastSeqno[getShardID(shard)] = shard.SeqNo
					newShards = append(newShards, notSeen...)
				}
				newShards = append(newShards, master)

				var txList []*tlb.Transaction

				for _, shard := range newShards {
					cmd.Printf(
						"scanning block %d of shard %x in workchain %d...\n",
						shard.SeqNo,
						uint64(shard.Shard),
						shard.Workchain,
					)

					var fetchedIDs []ton.TransactionShortInfo
					var after *ton.TransactionID3
					var more = true

					for more {
						fetchedIDs, more, err3 = api.WaitForBlock(master.SeqNo).GetBlockTransactionsV2(ctx, shard, 100, after)
						cobra.CheckErr(err3)

						if more {
							after = fetchedIDs[len(fetchedIDs)-1].ID3()
						}

						for _, id := range fetchedIDs {
							tx, err4 := api.GetTransaction(
								ctx,
								shard,
								address.NewAddress(0, byte(shard.Workchain), id.Account),
								id.LT,
							)
							cobra.CheckErr(err4)

							txList = append(txList, tx)
						}
					}
				}

				cmd.Printf("Found %d transactions\n", len(txList))

				if len(txList) == 0 {
					cmd.Printf("No transactions found in block %d\n", master.SeqNo)
				}

				master, err = api.WaitForBlock(master.SeqNo+1).LookupBlock(ctx, master.Workchain, master.Shard, master.SeqNo+1)
				cobra.CheckErr(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

func getNotSeenShards(
	ctx context.Context,
	api ton.APIClientWrapped,
	shard *ton.BlockIDExt,
	shardLastSeqno map[string]uint32,
) (ret []*ton.BlockIDExt, err error) {
	if no, ok := shardLastSeqno[getShardID(shard)]; ok && no == shard.SeqNo {
		return nil, nil
	}

	b, err := api.GetBlockData(ctx, shard)
	if err != nil {
		return nil, fmt.Errorf("get block data: %w", err)
	}

	parents, err := b.BlockInfo.GetParentBlocks()
	if err != nil {
		return nil, fmt.Errorf("get parent blocks (%d:%x:%d): %w", shard.Workchain, uint64(shard.Shard), shard.Shard, err)
	}

	for _, parent := range parents {
		ext, err2 := getNotSeenShards(ctx, api, parent, shardLastSeqno)
		if err2 != nil {
			return nil, err2
		}
		ret = append(ret, ext...)
	}

	ret = append(ret, shard)
	return ret, nil
}

func getShardID(shard *ton.BlockIDExt) string {
	return fmt.Sprintf("%d|%d", shard.Workchain, shard.Shard)
}
