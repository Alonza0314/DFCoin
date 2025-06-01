package cmd

import (
	"fmt"

	"github.com/Alonza0314/DFcoin/blockchain"
	"github.com/Alonza0314/DFcoin/logger"
	"github.com/spf13/cobra"
)

var printChainCmd = &cobra.Command{
	Use:   "printchain",
	Short: "Print all blocks in the blockchain",
	Run:   printChainFunc,
}

func printChainFunc(cmd *cobra.Command, args []string) {
	chain := blockchain.GetChain()
	if chain == nil {
		return
	}
	iter := chain.Iterator()

	for {
		block := iter.Next()
		logger.Log.Info("PRINT CHAIN", fmt.Sprintf("Previous hash: %x", block.PrevBlockHash))
		logger.Log.Info("PRINT CHAIN", fmt.Sprintf("Data: %s", block.Data))
		logger.Log.Info("PRINT CHAIN", fmt.Sprintf("Hash: %x", block.Hash))
		pow := blockchain.NewProofOfWork(block)
		logger.Log.Info("PRINT CHAIN", fmt.Sprintf("PoW: %v", pow.Validate()))
		logger.Log.Info("PRINT CHAIN", "--------------------------------")
		if string(block.Data) == blockchain.GENESIS_BLOCK_DATA {
			break
		}
	}
}

func init() {
	rootCmd.AddCommand(printChainCmd)
}