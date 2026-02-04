package main

import (
	"context"
	"fmt"
	"log"

	"go-eth-learning/internal/config"
	"go-eth-learning/pkg/ethclient"
	"go-eth-learning/pkg/transaction"
	"go-eth-learning/pkg/utils"
)

func main() {
	fmt.Println("⛓️🐹 Go-Ethereum 基础示例\\n")

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

	// 1. 查询区块信息
	fmt.Println("=== 网络信息 ===")
	fmt.Printf("Chain ID: %d\\n", client.ChainID())

	blockNum, err := client.GetBlockNumber(ctx)
	if err != nil {
		log.Printf("获取区块号失败: %v", err)
	} else {
		fmt.Printf("最新区块: %d\\n", blockNum)
	}

	// 2. 查询余额示例
	fmt.Println("\\n=== 查询余额 ===")
	address := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	fmt.Printf("查询地址: %s\\n", address)

	balance, err := client.GetBalance(ctx, address)
	if err != nil {
		log.Printf("查询余额失败: %v", err)
	} else {
		fmt.Printf("余额: %f ETH\\n", balance)
	}

	// 3. 交易管理器示例
	fmt.Println("\\n=== 交易管理器 ===")
	txManager := transaction.NewManager(client, client.ChainID())
	_ = txManager

	// 4. 工具函数示例
	fmt.Println("\\n=== 工具函数 ===")
	wei := utils.EtherToWei(1.5)
	fmt.Printf("1.5 ETH = %s Wei\\n", wei.String())

	ether := utils.WeiToEther(wei)
	fmt.Printf("转回 ETH: %f\\n", ether)

	// 验证地址
	valid := utils.IsValidAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	fmt.Printf("地址有效: %v\\n", valid)

	fmt.Println("\\n✅ 示例完成!")
}
