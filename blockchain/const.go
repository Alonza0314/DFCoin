package blockchain

import "math"

const (
	TARGET_BITS = 4 // leading zeros in the hash
	MAX_NONCE   = math.MaxInt64

	GENESIS_BLOCK_DATA = "Genesis Block"

	DB_PATH = ".db/blockchain.db"

	BLOCK_BUCKET    = "blocks"
	LAST_BLOCK_HASH = "last_block_hash"

	NOT_FOUND_ERROR = "not found"
)
