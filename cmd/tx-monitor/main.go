// cmd/tx-monitor 交易监控工具
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("⛓️🐹 交易监控工具")
	fmt.Println("监控新交易和区块...\\n")

	// 连接本地节点或 Infura
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/YOUR_KEY")
	if err != nil {
		// 回退到 HTTP
		client, err = ethclient.Dial("https://sepolia.infura.io/v3/YOUR_KEY")
		if err != nil {
			log.Fatalf("连接失败: %v", err)
		}
	}
	defer client.Close()

	ctx := context.Background()

	// 获取起始区块
	startBlock, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatalf("获取区块号失败: %v", err)
	}

	fmt.Printf("开始监控，当前区块: %d\\n", startBlock)
	fmt.Println("按 Ctrl+C 停止\\n")

	// 简单轮询监控
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	lastBlock := startBlock

	for range ticker.C {
		currentBlock, err := client.BlockNumber(ctx)
		if err != nil {
			log.Printf("获取区块号失败: %v", err)
			continue
		}

		if currentBlock > lastBlock {
			for blockNum := lastBlock + 1; blockNum <= currentBlock; blockNum++ {
				processBlock(client, blockNum)
			}
			lastBlock = currentBlock
		}
	}
}

func processBlock(client *ethclient.Client, blockNum uint64) {
	ctx := context.Background()

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockNum)))
	if err != nil {
		log.Printf("获取区块 %d 失败: %v", blockNum, err)
		return
	}

	fmt.Printf("📦 区块 #%d | 时间: %s | 交易: %d\\n",
		blockNum,
		time.Unix(int64(block.Time()), 0).Format("15:04:05"),
		len(block.Transactions()),
	)

	// 显示前 3 笔交易
	for i, tx := range block.Transactions() {
		if i >= 3 {
			break
		}
		displayTransaction(tx)
	}

	if len(block.Transactions()) > 3 {
		fmt.Printf("   ... 还有 %d 笔交易\\n", len(block.Transactions())-3)
	}
	fmt.Println()
}

func displayTransaction(tx *types.Transaction) {
	fmt.Printf("   💸 %s\\n", tx.Hash().Hex()[:20])
	if tx.To() != nil {
		fmt.Printf("      到: %s\\n", tx.To().Hex()[:20])
	}
	fmt.Printf("      金额: %s Wei\\n", tx.Value().String())
}

// 需要导入
import "math/big"
