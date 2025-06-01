package blockchain_test

import (
	"errors"
	"testing"

	"github.com/Alonza0314/DFcoin/blockchain"
)

const (
	Serialize   = "serialize"
	Deserialize = "deserialize"
)

var testBlockSerializeCases = []struct {
	name     string
	block    *blockchain.Block
	testType string
}{
	{
		name: "test block serialize",
		block: &blockchain.Block{
			Timestamp:     1620000000,
			Data:          []byte("test"),
			PrevBlockHash: []byte("prev"),
			Hash:          []byte("hash"),
			Nonce:         0,
		},
		testType: Serialize,
	},
	{
		name: "test block deserialize",
		block: &blockchain.Block{
			Timestamp:     1620000000,
			Data:          []byte("test"),
			PrevBlockHash: []byte("prev"),
			Hash:          []byte("hash"),
			Nonce:         0,
		},
		testType: Deserialize,
	},
}

func TestBlockSerialize(t *testing.T) {
	var serializedBytes []byte

	for _, testCase := range testBlockSerializeCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case Serialize:
				serializedBytes = blockchain.BlockSerialize(testCase.block)
				if len(serializedBytes) == 0 {
					t.Errorf("Serialize error: %v", errors.New("serialize error"))
				}
			case Deserialize:
				deserializedBlock := blockchain.BlockDeserialize(serializedBytes)
				if deserializedBlock == nil {
					t.Errorf("Deserialize error: %v", errors.New("deserialize error"))
				}
			}
		})
	}
}
