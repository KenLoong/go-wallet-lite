这是一份为你量身定制的 `README.md`。它不仅是一个项目说明，更是一份**学习笔记和面试复盘指南**。

你可以直接在你的 Go 项目根目录下创建这个文件。

---

# Go-Wallet-Lite 🚀

`Go-Wallet-Lite` 是一个基于 Golang 实现的轻量级以太坊后端服务 Demo。

该项目旨在通过代码实现，深入理解 **MetaMask** 和 **Coinbase Wallet** 等竞品钱包的底层上链逻辑，包括地址生成、资产查询、交易签名、Nonce 管理以及 EIP-1559 协议。

## 核心特性

- **账户管理**：基于 ECDSA 算法离线生成钱包地址与私钥。
- **资产查询**：实时查询以太坊（ETH）及标准 ERC-20 代币余额。
- **交易流程**：支持构造、离线签名并广播以太坊交易（兼容 EIP-1559）。
- **状态追踪**：通过交易哈希（TxHash）异步追踪链上确认状态。
- **竞品参考**：代码实现逻辑参考了 MetaMask (Infura 模式) 与 Coinbase Wallet (Smart Wallet 趋势)。

## 技术栈

- **语言**: Go 1.21+
- **以太坊库**: `go-ethereum` (Geth) —— 行业标准库
- **Web 框架**: `Gin` —— 高性能 API 开发
- **外部接入**: Alchemy RPC (Sepolia Testnet)
- **配置管理**: `godotenv`

## 目录结构

```text
.
├── api/
│   └── router.go       // Gin 路由与接口入口
├── service/
│   ├── wallet.go       // 账户生成与余额逻辑
│   ├── transfer.go     // 构造交易、签名与发送逻辑
│   └── contract.go     // ERC-20 合约交互 (ABI 封装)
├── blockchain/
│   └── client.go       // RPC 客户端初始化
├── abi/
│   └── erc20.json      // 标准 ERC-20 智能合约 ABI
├── .env                // 环境变量（API Key, 私钥等）
├── main.go             // 程序入口
└── README.md
```

## 快速开始

### 1. 获取 RPC 节点
前往 [Alchemy](https://www.alchemy.com/) 注册免费账号，并创建一个 **Sepolia** 测试网项目，获取你的 `HTTPS URL`。

### 2. 环境配置
创建 `.env` 文件：
```env
RPC_URL=https://eth-sepolia.g.alchemy.com/v2/你的API_KEY
PORT=8080
```

### 3. 运行项目
```bash
go mod tidy
go run main.go
```

## API 接口清单

| 接口 | 方法 | 功能 |
| :--- | :--- | :--- |
| `/v1/wallet/create` | `POST` | 离线生成新钱包（地址与私钥） |
| `/v1/balance/:address` | `GET` | 查询 ETH 余额 |
| `/v1/balance/token/:address` | `GET` | 查询指定 ERC-20 代币余额 |
| `/v1/transfer` | `POST` | 构造、签名并发送交易 |
| `/v1/tx/:hash` | `GET` | 检查交易确认状态 |

## 关键学习点 (针对 HR 回复内容)

在开发本项目时，需重点掌握以下 Web3 后端核心概念：

1.  **Nonce 管理**：为什么用户的交易必须是连续递增的？如何处理 Pending 交易导致的 Nonce 阻塞？
2.  **Gas 策略 (EIP-1559)**：理解 `BaseFee`, `MaxPriorityFee` (小费) 和 `MaxFee` 的关系。如何模仿 MetaMask 给用户提供“快、中、慢”三种费率选择？
3.  **ABI 交互**：后端如何通过合约地址和 ABI 与链上智能合约（如 USDT）进行读写交互？
4.  **离线签名**：为什么私钥永远不应该离开后端服务器（或硬件模块）？如何通过 `crypto/ecdsa` 进行 Secp256k1 签名？
5.  **竞品差异化**：
    *   **MetaMask**: 经典的 EOA (Externally Owned Account) 模式。
    *   **Coinbase Wallet**: 正在向 **Account Abstraction (账户抽象)** 演进，支持 Passkey 和无助记词体验。本项目实现的 `service/transfer.go` 是理解这些高级特性的基础。

## 风险提示
本项目仅供学习使用。**请勿在生产环境中存储明文私钥。** 在真实项目中，建议使用 AWS KMS、HashiCorp Vault 或专用的 HSM 模块进行密钥管理。

---

### 下一步计划
- [ ] 集成 **Safe (Gnosis Safe)** 多签合约接口。
- [ ] 增加数据库支持，记录用户历史交易记录（数据清洗/Indexer 雏形）。
- [ ] 支持 Webhook 监听链上到账通知。

---

### 💡 如何使用这个 README？
1. 在你的 Go 项目文件夹里创建 `README.md`。
2. 按照里面的目录结构创建文件夹。
3. 你的下一步就是填充 `blockchain/client.go` 里的连接代码。
