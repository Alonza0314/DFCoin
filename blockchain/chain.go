package blockchain

import (
	"fmt"
	"strings"

	"github.com/Alonza0314/DFcoin/bboltdb"
	"github.com/Alonza0314/DFcoin/logger"
)

type Chain struct {
	tip []byte // last block hash
}

// NewChain creates a new blockchain or loads the existing one
func NewChain() *Chain {
	// try to load last block hash from database
	lastHash, err := bboltdb.DbLoad(DB_PATH, []byte(BLOCK_BUCKET), []byte(LAST_BLOCK_HASH))
	if err != nil && !strings.Contains(err.Error(), NOT_FOUND_ERROR) {
		logger.Log.Error("CHAIN", fmt.Sprintf("NewChain error: %v", err))
		return nil
	}

	if lastHash != nil {
		logger.Log.Error("CHAIN", "Blockchain existed in database")
		return nil
	}

	logger.Log.Info("CHAIN", "No blockchain in database, creating genesis block")

	// no blockchain in database, create genesis block
	genesis := NewGenesisBlock()

	err = bboltdb.DbSave(DB_PATH, []byte(BLOCK_BUCKET), genesis.Hash, BlockSerialize(genesis))
	if err != nil {
		logger.Log.Error("CHAIN", fmt.Sprintf("NewChain error: %v", err))
		return nil
	}

	err = bboltdb.DbSave(DB_PATH, []byte(BLOCK_BUCKET), []byte(LAST_BLOCK_HASH), genesis.Hash)
	if err != nil {
		logger.Log.Error("CHAIN", fmt.Sprintf("NewChain error: %v", err))
		return nil
	}

	logger.Log.Info("CHAIN", "Blockchain created")

	return &Chain{tip: genesis.Hash}
}

// GetChain gets the blockchain from the database
func GetChain() *Chain {
	lastHash, err := bboltdb.DbLoad(DB_PATH, []byte(BLOCK_BUCKET), []byte(LAST_BLOCK_HASH))
	if err != nil {
		logger.Log.Error("CHAIN", fmt.Sprintf("GetChain error: %v", err))
		return nil
	}

	return &Chain{tip: lastHash}
}

// AddBlock adds a new block to the blockchain
func (c *Chain) AddBlock(data string) {
	// get last block hash
	lastHash, err := bboltdb.DbLoad(DB_PATH, []byte(BLOCK_BUCKET), []byte(LAST_BLOCK_HASH))
	if err != nil {
		logger.Log.Error("CHAIN", fmt.Sprintf("AddBlock error: %v", err))
		return
	}

	// create new block
	newBlock := NewBlock(data, lastHash)

	// save new block
	err = bboltdb.DbSave(DB_PATH, []byte(BLOCK_BUCKET), newBlock.Hash, BlockSerialize(newBlock))
	if err != nil {
		logger.Log.Error("CHAIN", fmt.Sprintf("AddBlock error: %v", err))
		return
	}

	// update last block hash
	err = bboltdb.DbUpdate(DB_PATH, []byte(BLOCK_BUCKET), []byte(LAST_BLOCK_HASH), newBlock.Hash)
	if err != nil {
		logger.Log.Error("CHAIN", fmt.Sprintf("AddBlock error: %v", err))
		return
	}

	c.tip = newBlock.Hash
}

// add a iterator to iterate through the blockchain
type ChainIterator struct {
	currentHash []byte
}

func (c *Chain) Iterator() *ChainIterator {
	return &ChainIterator{c.tip}
}

func (i *ChainIterator) Next() *Block {
	var block *Block

	// get block from database
	blockData, err := bboltdb.DbLoad(DB_PATH, []byte(BLOCK_BUCKET), i.currentHash)
	if err != nil {
		logger.Log.Error("CHAIN ITERATOR", fmt.Sprintf("Next error: %v", err))
		return nil
	}
	if blockData == nil {
		return nil
	}

	block = BlockDeserialize(blockData)

	// update iterator, point to previous block
	i.currentHash = block.PrevBlockHash

	return block
}
