// cmd/contract 智能合约交互工具
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go-eth-learning/internal/config"
	"go-eth-learning/pkg/ethclient"
)

func main() {
	fmt.Println("⛓️🐹 智能合约交互工具\\n")

	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建客户端
	client, err := ethclient.New(cfg.EthNodeURL)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// 显示网络信息
	fmt.Println("=== 网络信息 ===")
	fmt.Printf("Chain ID: %d\\n", client.ChainID())

	blockNum, err := client.GetBlockNumber(ctx)
	if err != nil {
		log.Printf("获取区块号失败: %v", err)
	} else {
		fmt.Printf("最新区块: %d\\n", blockNum)
	}

	// 查询余额示例
	if len(os.Args) > 1 {
		address := os.Args[1]
		fmt.Printf("\\n=== 查询余额 ===\\n")
		fmt.Printf("地址: %s\\n", address)

		balance, err := client.GetBalance(ctx, address)
		if err != nil {
			log.Printf("查询余额失败: %v", err)
		} else {
			fmt.Printf("余额: %f ETH\\n", balance)
		}
	}
}
