# DFCoin: 自設計加密貨幣

## 設計階段

### 第一階段：核心區塊鏈與 CLI 初步

| 模組            | 功能                                                        |
| ------------- | --------------------------------------------------------- |
| Block         | 區塊結構（索引、時間戳、前一區塊雜湊、交易、Nonce）                              |
| Blockchain    | 區塊鏈結構與儲存（記憶體 or BoltDB）                                   |
| Proof-of-Work | 基於簡單難度的挖礦邏輯                                               |
| CLI           | Cobra 製作 `createblockchain`, `addblock`, `printchain` 等命令 |

### 第二階段：交易與 UTXO 模型

| 模組          | 功能                                      |
| ----------- | --------------------------------------- |
| Transaction | 定義交易結構：輸入 (`TXInput`) / 輸出 (`TXOutput`) |
| UTXO        | 可用輸出管理邏輯（查找、更新、索引）                      |
| Coinbase TX | 區塊獎勵交易支援                                |
| CLI         | `send`, `getbalance` 命令加入               |

### 第三階段：錢包與金鑰管理

| 模組      | 功能                                  |
| ------- | ----------------------------------- |
| Wallet  | 金鑰對生成（使用 ECDSA 或 secp256k1）         |
| Address | 錢包地址格式與 Base58Check 編碼              |
| Wallets | 多錢包管理與保存                            |
| CLI     | `createwallet`, `listaddresses` 等命令 |

### 第四階段：P2P 網路與節點同步

| 模組               | 功能                                                  |
| ---------------- | --------------------------------------------------- |
| P2P              | 基於 TCP 實作節點之間的資料傳播                                  |
| Message Protocol | 定義 `version`, `inv`, `getdata`, `block`, `tx` 等訊息協議 |
| Node             | 節點同步區塊鏈、接收交易                                        |
| CLI              | `startnode` 加入網路節點命令                                |

### 第五階段：API 與前端擴展介面

| 模組         | 功能                                        |
| ---------- | ----------------------------------------- |
| REST API   | 提供 `/balance`, `/send`, `/mine` 等 JSON 介面 |
| Web Wallet | 簡易 Web UI for 錢包與交易操作                     |
| Swagger    | 文件說明與商業導向 API 文檔                          |

### 第六階段：測試網、部署與商業優化

| 項目                  | 功能                        |
| ------------------- | ------------------------- |
| Testnet             | 測試鏈：獨立節點、測試幣              |
| Docker / Kubernetes | 輕量部署架構支援                  |
| Security            | 證書加密、金鑰保管策略（如 keystore）   |
| 商業介接                | 支援支付系統、費用系統、Token Wrapper |

## 專案架構

```bash
DFcoin/
├── blockchain/
│   ├── block.go
│   ├── blockchain.go
│   ├── pow.go
├── transaction/
│   ├── transaction.go
│   ├── utxo.go
├── wallet/
│   ├── wallet.go
│   ├── wallets.go
├── network/
│   ├── node.go
│   ├── server.go
├── cli/
│   └── cli.go
├── db/
│   └── bolt.go
├── main.go

```
