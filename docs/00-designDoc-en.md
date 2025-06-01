# DFCoin: Self-Designed Cryptocurrency

## System Architecture

| Module | Description |
| - | - |
| Blockchain Core | Block structure, chain linking, tamper-proof verification, genesis block |
| Cryptocurrency Model (GoCoin) | Bitcoin-like implementation with **UTXO** architecture, transaction verification, mempool |
| Consensus Algorithm (Proof of Work) | PoW mining difficulty adjustment, block validity verification |
| Wallet | ECDSA-based public/private key address generation, transaction signature verification |
| Node | Multiple deployable nodes, block synchronization, transaction and block propagation |
| P2P Network (Simple) | Inter-node messaging, automatic chain and transaction sync (using libp2p / TCP) |
| API (REST + CLI) | User transaction sending, blockchain querying, balance checking |
| Block Explorer (Optional) | Web frontend displaying blocks, transactions, mining status (supports React/Vue development) |

## Final Goals

- Blockchain Core
  - Block contents: transaction list, previous block hash, timestamp, difficulty, nonce, Merkle root
  - Automatic difficulty adjustment mechanism
  - Double-spending prevention
- Cryptocurrency Mechanism
  - Complete UTXO model (Unspent Transaction Outputs)
  - Support for transfers, change, multi-input/output transactions
  - Support for mempool (transactions stored temporarily before being packed)
- Wallet and Addresses
  - Using ECDSA key pairs (Go's crypto/ecdsa)
  - Base58 encoded wallet addresses (same format as Bitcoin)
  - Users can generate private keys / public keys / addresses
  - Can sign transactions and verify signatures
- Consensus and Mining
  - Basic Proof of Work mining
  - Dynamic difficulty adjustment (based on average time of past N blocks)
  - Miners can pack transactions and receive block rewards
- Node Network
  - Multi-node synchronization mechanism (TCP / HTTP transmission)
  - Automatic latest chain download by nodes
  - Gossip propagation of transactions and blocks
- RESTful API & CLI
  - /createwallet
  - /getbalance/:address
  - /send (from, to, amount)
  - /mine
  - /chain query current complete blockchain
- Testing and Demonstration
  - Set up three nodes synchronizing data
  - Send multiple transactions and observe how they are packed into blocks
  - System can operate locally without external dependencies

## Key Technologies

| Feature | Corresponding Engineering Skills |
| - | - |
| UTXO Model | Real Bitcoin ledger logic, not account-balance based |
| ECDSA Signatures and Verification | Cryptography application, same as real wallets and transaction verification |
| Proof of Work | Consensus mechanism and mining algorithm fundamentals |
| P2P Communication | Decentralized data synchronization mindset |
| RESTful + CLI Tools | Complete system interface, suitable for project demonstration |
| Multi-node Data Consistency | Key blockchain network capability (expandable to DHT, Raft, libp2p) |

## Phased Design

| Phase | Features and Tasks | Difficulty |
| - | - | - |
| Phase 1 | Blockchain data structure + Mining + CLI | ⭐⭐ |
| Phase 2 | UTXO model + Multi-transaction support + Wallet | ⭐⭐⭐ |
| Phase 3 | API Integration / RESTful Server | ⭐⭐⭐ |
| Phase 4 | Multi-node data sync + Simple P2P communication | ⭐⭐⭐⭐ |
| Phase 5 | Add dynamic difficulty + Block rewards + Complete transaction verification flow | ⭐⭐⭐⭐ |
| Phase 6 | Package as demo project + README + CLI Demo | ⭐⭐ |

## Extended Design

- Support for Merkle Tree, SPV light nodes
- Simulate smart contract execution logic (simplified VM)
- Block reward halving mechanism (mimicking Bitcoin Halving)
- Integrate React to create a simple "GoScan" (block explorer)
- Implement simple DHT (decentralized node discovery)

## External References

- [Design Document-ChatGPT](https://chatgpt.com)
