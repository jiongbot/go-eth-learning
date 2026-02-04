# Go-Ethereum 开发指南

## 环境准备

### 1. 安装依赖

```bash
# 安装 go-ethereum
go get github.com/ethereum/go-ethereum

# 安装其他依赖
go mod tidy
```

### 2. 配置环境变量

创建 `.env` 文件：

```env
# 以太坊节点 RPC 地址
ETH_NODE_URL=https://sepolia.infura.io/v3/YOUR_INFURA_KEY

# 私钥（用于发送交易，可选）
PRIVATE_KEY=your_private_key_here
```

### 3. 获取 Infura API Key

1. 访问 https://infura.io
2. 注册账号
3. 创建新项目
4. 复制 API Key

## 快速开始

### 查询余额

```bash
go run examples/basic/main.go
```

### 创建钱包

```bash
go run cmd/wallet/main.go
```

### 监控交易

```bash
go run cmd/tx-monitor/main.go
```

## 核心概念

### 连接以太坊节点

```go
client, err := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_KEY")
if err != nil {
    log.Fatal(err)
}
defer client.Close()
```

### 查询余额

```go
address := common.HexToAddress("0x...")
balance, err := client.BalanceAt(context.Background(), address, nil)
```

### 发送交易

```go
// 1. 创建交易
tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

// 2. 签名交易
signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

// 3. 发送交易
err = client.SendTransaction(context.Background(), signedTx)
```

## 项目结构说明

```
go-eth-learning/
├── cmd/              # 可执行命令
│   ├── wallet/       # 钱包管理
│   ├── contract/     # 合约交互
│   ├── tx-monitor/   # 交易监控
│   └── event-listener/ # 事件监听
├── pkg/              # 公共库
│   ├── ethclient/    # 客户端封装
│   ├── wallet/       # 钱包功能
│   ├── transaction/  # 交易管理
│   ├── contract/     # 合约 ABI
│   └── utils/        # 工具函数
├── internal/         # 私有代码
│   └── config/       # 配置管理
└── examples/         # 示例代码
```

## 常用操作

### 1. 创建账户

```go
privateKey, _ := crypto.GenerateKey()
address := crypto.PubkeyToAddress(privateKey.PublicKey)
```

### 2. 加载私钥

```go
privateKey, err := crypto.HexToECDSA(privateKeyHex)
```

### 3. 查询区块

```go
block, err := client.BlockByNumber(context.Background(), big.NewInt(12345))
```

### 4. 监听事件

```go
query := ethereum.FilterQuery{
    Addresses: []common.Address{contractAddress},
    Topics:    [][]common.Hash{{eventTopic}},
}

logs, err := client.FilterLogs(context.Background(), query)
```

## 测试网水龙头

- Sepolia: https://sepoliafaucet.com
- Goerli: https://goerlifaucet.com

## 安全提示

⚠️ **永远不要：**
- 在代码中硬编码私钥
- 将私钥提交到 Git
- 在公共频道分享私钥

✅ **最佳实践：**
- 使用环境变量存储敏感信息
- 使用硬件钱包管理大额资金
- 测试网开发，主网部署前审计

## 参考资源

- [go-ethereum 文档](https://geth.ethereum.org/docs/)
- [Ethereum Go API](https://pkg.go.dev/github.com/ethereum/go-ethereum)
- [Web3.js 对比](https://ethereum.org/en/developers/docs/apis/javascript/)

---
*Happy Hacking with Go + Ethereum!*
