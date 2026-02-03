// Go-Ethereum 基础操作示例
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	godotenv.Load()

	fmt.Println("⛓️🐹 Go-Ethereum 基础操作演示\\n")

	// 1. 连接到以太坊节点
	client := connectToNode()
	defer client.Close()

	// 2. 查询区块信息
	queryBlockInfo(client)

	// 3. 查询账户余额
	queryBalance(client)

	// 4. 创建新账户
	createAccount()

	// 5. 发送交易（需要私钥和资金）
	// sendTransaction(client)

	// 6. 查询交易
	queryTransaction(client)
}

// ==================== 连接到节点 ====================
func connectToNode() *ethclient.Client {
	fmt.Println("=== 连接到以太坊节点 ===")

	// 使用 Infura 或 Alchemy 的 RPC 端点
	// 或者本地节点: http://localhost:8545
	rpcURL := os.Getenv("ETH_RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://sepolia.infura.io/v3/YOUR_INFURA_KEY"
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}

	// 检查连接
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("获取网络 ID 失败: %v", err)
	}

	fmt.Printf("✅ 已连接到以太坊网络\\n")
	fmt.Printf("   Chain ID: %d\\n", chainID)

	return client
}

// ==================== 查询区块信息 ====================
func queryBlockInfo(client *ethclient.Client) {
	fmt.Println("\\n=== 查询区块信息 ===")

	// 获取最新区块号
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("获取区块头失败: %v", err)
	}

	fmt.Printf("最新区块号: %d\\n", header.Number.Uint64())
	fmt.Printf("区块哈希: %s\\n", header.Hash().Hex())
	fmt.Printf("时间戳: %d\\n", header.Time)
	fmt.Printf("Gas 限制: %d\\n", header.GasLimit)
	fmt.Printf("Gas 使用: %d\\n", header.GasUsed)

	// 获取特定区块
	blockNumber := big.NewInt(56789)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		fmt.Printf("获取区块 %d 失败: %v\\n", blockNumber, err)
	} else {
		fmt.Printf("\\n区块 %d 信息:\\n", blockNumber)
		fmt.Printf("  交易数: %d\\n", len(block.Transactions()))
		fmt.Printf("  难度: %d\\n", block.Difficulty())
	}
}

// ==================== 查询余额 ====================
func queryBalance(client *ethclient.Client) {
	fmt.Println("\\n=== 查询账户余额 ===")

	// 示例地址（Vitalik 的地址）
	address := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")

	// 查询余额
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("查询余额失败: %v", err)
	}

	// 转换为 ETH
	ethBalance := new(big.Float).Quo(
		new(big.Float).SetInt(balance),
		big.NewFloat(1e18),
	)

	fmt.Printf("地址: %s\\n", address.Hex())
	fmt.Printf("余额: %f ETH\\n", ethBalance)

	// 查询特定区块的余额
	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), address, blockNumber)
	if err != nil {
		log.Fatalf("查询历史余额失败: %v", err)
	}

	ethBalanceAt := new(big.Float).Quo(
		new(big.Float).SetInt(balanceAt),
		big.NewFloat(1e18),
	)
	fmt.Printf("区块 %d 余额: %f ETH\\n", blockNumber, ethBalanceAt)

	// 查询待处理余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		log.Fatalf("查询待处理余额失败: %v", err)
	}

	pendingEthBalance := new(big.Float).Quo(
		new(big.Float).SetInt(pendingBalance),
		big.NewFloat(1e18),
	)
	fmt.Printf("待处理余额: %f ETH\\n", pendingEthBalance)
}

// ==================== 创建账户 ====================
func createAccount() {
	fmt.Println("\\n=== 创建新账户 ===")

	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("生成私钥失败: %v", err)
	}

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法转换公钥")
	}

	// 获取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 私钥转换为字符串
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)

	fmt.Printf("✅ 新账户创建成功\\n")
	fmt.Printf("   地址: %s\\n", address.Hex())
	fmt.Printf("   私钥: %s\\n", privateKeyHex)
	fmt.Println("   ⚠️  请安全保存私钥，不要泄露！")
}

// ==================== 发送交易 ====================
func sendTransaction(client *ethclient.Client) {
	fmt.Println("\\n=== 发送交易 ===")

	// 从环境变量读取私钥
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		fmt.Println("未设置 PRIVATE_KEY 环境变量，跳过交易发送")
		return
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("加载私钥失败: %v", err)
	}

	// 获取发送者地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法转换公钥")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("获取 nonce 失败: %v", err)
	}

	// 获取 Gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("获取 Gas 价格失败: %v", err)
	}

	// 目标地址和金额
	toAddress := common.HexToAddress("0xRecipientAddress")
	value := big.NewInt(1000000000000000000) // 1 ETH

	// Gas 限制
	gasLimit := uint64(21000)

	// 获取 Chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("获取 Chain ID 失败: %v", err)
	}

	// 创建交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("签名交易失败: %v", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("发送交易失败: %v", err)
	}

	fmt.Printf("✅ 交易已发送\\n")
	fmt.Printf("   交易哈希: %s\\n", signedTx.Hash().Hex())
	fmt.Printf("   查看: https://sepolia.etherscan.io/tx/%s\\n", signedTx.Hash().Hex())
}

// ==================== 查询交易 ====================
func queryTransaction(client *ethclient.Client) {
	fmt.Println("\\n=== 查询交易 ===")

	// 示例交易哈希
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7d9a2b8c5d4e3f2a1b0c9d8e7f6a5b4c3d2e1f0a1b")

	// 查询交易
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Printf("查询交易失败: %v\\n", err)
		return
	}

	fmt.Printf("交易哈希: %s\\n", tx.Hash().Hex())
	fmt.Printf("待处理: %v\\n", isPending)
	fmt.Printf("发送者: %s\\n", getSender(tx))
	fmt.Printf("接收者: %s\\n", tx.To().Hex())
	fmt.Printf("金额: %s Wei\\n", tx.Value().String())
	fmt.Printf("Gas 限制: %d\\n", tx.Gas())
	fmt.Printf("Gas 价格: %s\\n", tx.GasPrice().String())
	fmt.Printf("Nonce: %d\\n", tx.Nonce())

	// 查询收据（确认状态）
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		fmt.Printf("查询收据失败: %v\\n", err)
		return
	}

	fmt.Printf("\\n交易收据:\\n")
	fmt.Printf("  状态: %d (1=成功, 0=失败)\\n", receipt.Status)
	fmt.Printf("  Gas 使用: %d\\n", receipt.GasUsed)
	fmt.Printf("  区块号: %d\\n", receipt.BlockNumber.Uint64())
	fmt.Printf("  区块哈希: %s\\n", receipt.BlockHash.Hex())
}

// 辅助函数：获取交易发送者
func getSender(tx *types.Transaction) string {
	// 简化处理，实际需要解析签名
	return "需要从签名恢复"
}
