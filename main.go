package main

import (
	"github.com/Alonza0314/DFcoin/cmd"
	"github.com/Alonza0314/DFcoin/logger"
)

func main() {
	logger.Log = logger.NewLogger()
	cmd.Execute()
}
