package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dfcoin",
	Short: "DFCoin is a blockchain implementation",
	Long:  "DFCoin is a blockchain implementation, used for learning blockchain basic concepts and implementation.",
}

// Execute execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
