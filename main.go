package main

import (
	"github.com/Alonza0314/DFCoin/cmd"
	"github.com/Alonza0314/DFCoin/logger"
)

func main() {
	logger.Log = logger.NewLogger()
	cmd.Execute()
}
