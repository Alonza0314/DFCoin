package cmd

import (
	"github.com/Alonza0314/DFCoin/blockchain"
	"github.com/spf13/cobra"
)

var createChainCmd = &cobra.Command{
	Use:   "createchain",
	Short: "Create a new blockchain",
	Run:   createChainFunc,
}

func createChainFunc(cmd *cobra.Command, args []string) {
	blockchain.NewChain()
}

func init() {
	rootCmd.AddCommand(createChainCmd)
}
