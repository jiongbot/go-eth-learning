// ERC20 代币操作示例
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-eth-learning/internal/config"
	"go-eth-learning/pkg/contract"
)

func main() {
	fmt.Println("🪙 ERC20 代币操作示例\\n")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	client, err := ethclient.Dial(cfg.EthNodeURL)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// USDT 合约地址 (Ethereum Mainnet)
	usdtAddress := "0xdAC17F958D2ee523a2206206994597C13D831ec7"

	fmt.Println("=== 查询 USDT 代币信息 ===")
	fmt.Printf("合约地址: %s\\n\\n", usdtAddress)

	// 创建合约实例
	token, err := contract.NewERC20Contract(usdtAddress)
	if err != nil {
		log.Fatalf("创建合约实例失败: %v", err)
	}

	// 查询代币信息
	fmt.Printf("代币地址: %s\\n", token.Address.Hex())

	// 查询余额示例
	walletAddress := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
	fmt.Printf("\\n查询地址余额: %s\\n", walletAddress)

	// 使用 ethclient 查询 ETH 余额
	balance, err := client.BalanceAt(ctx, common.HexToAddress(walletAddress), nil)
	if err != nil {
		log.Printf("查询 ETH 余额失败: %v", err)
	} else {
		ethBalance := new(big.Float).Quo(
			new(big.Float).SetInt(balance),
			big.NewFloat(1e18),
		)
		fmt.Printf("ETH 余额: %f\\n", ethBalance)
	}

	fmt.Println("\\n=== 代币操作说明 ===")
	fmt.Println("1. 查询代币余额需要调用合约的 balanceOf 方法")
	fmt.Println("2. 转账需要调用 transfer 方法并签名交易")
	fmt.Println("3. 授权需要调用 approve 方法")
	fmt.Println("4. 实际交互需要完整的合约绑定代码（可用 abigen 生成）")

	fmt.Println("\\n✅ ERC20 示例完成!")
}
