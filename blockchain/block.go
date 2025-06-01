package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/Alonza0314/DFcoin/logger"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock creates a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates a new genesis block
func NewGenesisBlock() *Block {
	return NewBlock(GENESIS_BLOCK_DATA, []byte{})
}

// BlockSerialize serializes the block
func BlockSerialize(b *Block) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	if err := encoder.Encode(b); err != nil {
		logger.Log.Error("BLOCK", fmt.Sprintf("Serialize error: %v", err))
	}

	return result.Bytes()
}

// BlockDeserialize deserializes the block
func BlockDeserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	if err := decoder.Decode(&block); err != nil {
		logger.Log.Error("BLOCK", fmt.Sprintf("Deserialize error: %v", err))
	}

	return &block
}
