# Go-Ethereum 10天学习计划 📅

> 从 Go 基础到以太坊开发，每天2-3小时

---

## 前置要求

- ✅ Go 基础语法（goroutine、channel、interface）
- ✅ 了解区块链基本概念
- ❌ 不需要 Solidity 基础

---

## Day 1: 环境搭建与项目结构

**目标**: 跑通第一个以太坊连接

### 上午 (1h)
- [ ] 安装依赖: `go mod tidy`
- [ ] 注册 Infura 获取 API Key: https://infura.io
- [ ] 创建 `.env` 文件

```bash
# .env
ETH_NODE_URL=https://sepolia.infura.io/v3/YOUR_KEY
```

### 下午 (1h)
- [ ] 阅读: `pkg/ethclient/client.go` (只读连接部分)
- [ ] 运行: `go run examples/basic/main.go`
- [ ] 理解: Chain ID、区块号的含义

**关键代码**: [pkg/ethclient/client.go#L28-L45](./pkg/ethclient/client.go)

---

## Day 2: 账户与钱包

**目标**: 创建和管理以太坊账户

### 上午 (1h)
- [ ] 阅读: `pkg/wallet/wallet.go`
- [ ] 理解: 私钥、公钥、地址的关系
- [ ] 运行: `go run cmd/wallet/main.go`

### 下午 (1.5h)
- [ ] 创建 3 个新钱包
- [ ] 保存地址和私钥到安全地方
- [ ] 阅读: `pkg/utils/utils.go` 的地址验证

**关键代码**:
- 创建钱包: [pkg/wallet/wallet.go#L20-L28](./pkg/wallet/wallet.go)
- 地址验证: [pkg/utils/utils.go#L28-L42](./pkg/utils/utils.go)

---

## Day 3: 查询操作

**目标**: 查询余额和区块信息

### 上午 (1.5h)
- [ ] 阅读: `pkg/ethclient/client.go` GetBalance 方法
- [ ] 查询任意地址 ETH 余额
- [ ] 理解: Wei vs Ether 单位转换

### 下午 (1.5h)
- [ ] 查询最新区块号
- [ ] 查询特定区块信息
- [ ] 阅读: `examples/basic/main.go`

**关键代码**:
- 余额查询: [pkg/ethclient/client.go#L47-L62](./pkg/ethclient/client.go)
- 单位转换: [pkg/utils/utils.go#L11-L22](./pkg/utils/utils.go)

**练习**: 写一个程序查询 5 个地址的余额

---

## Day 4: 交易基础

**目标**: 理解交易的构成

### 上午 (1.5h)
- [ ] 阅读: `pkg/transaction/manager.go`
- [ ] 理解: nonce、gasPrice、gasLimit
- [ ] 理解: 交易签名流程

### 下午 (1.5h)
- [ ] 从 Sepolia 水龙头获取测试币
- [ ] 阅读交易构建代码，不发送

**关键代码**:
- 交易构建: [pkg/transaction/manager.go#L26-L56](./pkg/transaction/manager.go)
- 签名发送: [pkg/transaction/manager.go#L58-L86](./pkg/transaction/manager.go)

**水龙头**: https://sepoliafaucet.com

---

## Day 5: 发送第一笔交易

**目标**: 成功发送 ETH 转账

### 上午 (2h)
- [ ] 准备: 两个钱包，一个有测试币
- [ ] 修改: `examples/basic/main.go` 添加转账代码
- [ ] 发送: 0.001 ETH 到另一个地址

### 下午 (1h)
- [ ] 在 Etherscan 查看交易
- [ ] 理解: 交易状态、确认数
- [ ] 阅读: `pkg/ethclient/client.go` WaitMined

**关键代码**: [pkg/transaction/manager.go#L88-L108](./pkg/transaction/manager.go)

**验证**: Sepolia Etherscan 查看交易状态

---

## Day 6: 智能合约基础

**目标**: 理解合约 ABI 和调用

### 上午 (1.5h)
- [ ] 阅读: `contracts/SimpleStorage.sol`
- [ ] 理解: 合约部署和调用的区别
- [ ] 阅读: `pkg/contract/erc20.go`

### 下午 (1.5h)
- [ ] 学习 ABI 是什么
- [ ] 阅读 ERC20 标准接口
- [ ] 理解 `Transfer` 事件

**关键代码**:
- ERC20 ABI: [pkg/contract/erc20.go](./pkg/contract/erc20.go)
- Solidity 合约: [contracts/SimpleStorage.sol](./contracts/SimpleStorage.sol)

---

## Day 7: ERC20 代币交互

**目标**: 查询代币余额

### 上午 (2h)
- [ ] 阅读: `examples/token/main.go`
- [ ] 运行代币示例
- [ ] 查询 USDT 合约信息

### 下午 (1h)
- [ ] 理解 `balanceOf` 调用
- [ ] 尝试查询其他代币
- [ ] 阅读: `internal/service/service.go`

**关键代码**: [examples/token/main.go](./examples/token/main.go)

---

## Day 8: 事件监听

**目标**: 监听区块链事件

### 上午 (1.5h)
- [ ] 阅读: `cmd/event-listener/main.go`
- [ ] 理解: Event Topic、Filter
- [ ] 理解: 日志结构

### 下午 (1.5h)
- [ ] 运行事件监听器
- [ ] 等待并观察 Transfer 事件
- [ ] 修改代码监听其他事件

**关键代码**: [cmd/event-listener/main.go](./cmd/event-listener/main.go)

---

## Day 9: 高级功能

**目标**: 掌握监控和批量操作

### 上午 (1.5h)
- [ ] 阅读: `cmd/tx-monitor/main.go`
- [ ] 运行交易监控
- [ ] 理解区块监听逻辑

### 下午 (1.5h)
- [ ] 阅读: `internal/service/service.go` 完整服务层
- [ ] 理解业务逻辑分层
- [ ] 尝试扩展服务功能

**关键代码**:
- 交易监控: [cmd/tx-monitor/main.go](./cmd/tx-monitor/main.go)
- 服务层: [internal/service/service.go](./internal/service/service.go)

---

## Day 10: 综合实战

**目标**: 完成一个完整功能

### 全天 (3h)
选择以下一个项目完成：

**选项 A**: 钱包监控工具
- 监控指定地址的余额变化
- 余额变动时打印通知
- 参考: `cmd/tx-monitor/`

**选项 B**: 批量查询工具
- 从文件读取地址列表
- 批量查询余额并输出 CSV
- 参考: `examples/basic/`

**选项 C**: 简单转账工具
- 交互式转账程序
- 输入地址和金额，确认后发送
- 参考: `cmd/wallet/` + `pkg/transaction/`

---

## 学习路径图

```
Day 1-2: 基础连接 + 账户
    ↓
Day 3-5: 查询 + 交易
    ↓
Day 6-7: 合约基础
    ↓
Day 8-9: 事件 + 监控
    ↓
Day 10: 综合实战
```

---

## 每日学习流程

```
1. 阅读指定代码文件 (30min)
2. 运行示例程序 (30min)
3. 修改代码实验 (1h)
4. 记录问题和收获 (笔记)
```

---

## 关键文件索引

| 功能 | 文件路径 |
|------|----------|
| 连接节点 | `pkg/ethclient/client.go` |
| 钱包管理 | `pkg/wallet/wallet.go` |
| 交易管理 | `pkg/transaction/manager.go` |
| 工具函数 | `pkg/utils/utils.go` |
| 配置管理 | `internal/config/config.go` |
| 业务服务 | `internal/service/service.go` |
| 合约 ABI | `pkg/contract/erc20.go` |
| 钱包 CLI | `cmd/wallet/main.go` |
| 事件监听 | `cmd/event-listener/main.go` |
| 交易监控 | `cmd/tx-monitor/main.go` |
| 基础示例 | `examples/basic/main.go` |
| 代币示例 | `examples/token/main.go` |
| NFT 示例 | `examples/nft/main.go` |
| Solidity | `contracts/*.sol` |

---

## 遇到问题？

1. **代码看不懂**: 先看注释，再看函数签名
2. **运行报错**: 检查 `.env` 配置和网络连接
3. **概念不理解**: 查阅 `docs/guide.md`
4. **需要更多示例**: 查看 `tests/` 目录的测试代码

---

## 下一步

完成本计划后，可以学习：
- 使用 `abigen` 生成完整合约绑定
- 部署自己的合约
- Layer2 (Polygon, Arbitrum) 开发
- DeFi 协议交互

---

*Start small, build consistently.* 🚀
