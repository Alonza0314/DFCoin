package cmd

import (
	"fmt"
	"os"

	"github.com/Alonza0314/DFcoin/blockchain"
	"github.com/Alonza0314/DFcoin/logger"
	"github.com/spf13/cobra"
)

var addBlockCmd = &cobra.Command{
	Use:   "addblock",
	Short: "Add a new block",
	Run:   addBlockFunc,
}

func addBlockFunc(cmd *cobra.Command, args []string) {
	data, err := cmd.Flags().GetString("data")
	if err != nil {
		logger.Log.Error("ADD BLOCK", fmt.Sprintf("Error getting data: %v", err))
		os.Exit(1)
	}

	chain := blockchain.GetChain()
	if chain == nil {
		return
	}
	chain.AddBlock(data)
	logger.Log.Info("ADD BLOCK", "Block added successfully!")
}

func init() {
	rootCmd.AddCommand(addBlockCmd)
	addBlockCmd.Flags().StringP("data", "d", "", "Block data")
	if err := addBlockCmd.MarkFlagRequired("data"); err != nil {
		logger.Log.Error("ADD BLOCK", fmt.Sprintf("Error marking data as required: %v", err))
		os.Exit(1)
	}
}
