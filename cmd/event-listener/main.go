// cmd/event-listener 事件监听工具
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ERC20 Transfer 事件 ABI
const erc20TransferABI = `[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`

func main() {
	fmt.Println("⛓️🐹 事件监听工具")
	fmt.Println("监听 ERC20 Transfer 事件...\\n")

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_KEY")
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(erc20TransferABI))
	if err != nil {
		log.Fatalf("解析 ABI 失败: %v", err)
	}

	// Transfer 事件 topic
	transferEvent := parsedABI.Events["Transfer"]
	transferTopic := transferEvent.ID

	fmt.Printf("监听事件 Topic: %s\\n", transferTopic.Hex())
	fmt.Println("按 Ctrl+C 停止\\n")

	// 查询过去的事件
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   nil, // 最新
		Topics: [][]common.Hash{
			{transferTopic},
		},
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		log.Printf("查询日志失败: %v", err)
	} else {
		fmt.Printf("找到 %d 个历史事件\\n", len(logs))
		for _, vLog := range logs {
			displayTransferEvent(parsedABI, vLog)
		}
	}

	// 实时监听（简化版，实际使用订阅）
	fmt.Println("\\n开始实时监控...")
	lastBlock, _ := client.BlockNumber(ctx)

	for {
		time.Sleep(10 * time.Second)

		currentBlock, err := client.BlockNumber(ctx)
		if err != nil {
			continue
		}

		if currentBlock > lastBlock {
			query.FromBlock = big.NewInt(int64(lastBlock + 1))
			query.ToBlock = big.NewInt(int64(currentBlock))

			logs, err := client.FilterLogs(ctx, query)
			if err != nil {
				continue
			}

			for _, vLog := range logs {
				displayTransferEvent(parsedABI, vLog)
			}

			lastBlock = currentBlock
		}
	}
}

func displayTransferEvent(parsedABI abi.ABI, vLog types.Log) {
	fmt.Printf("\\n📤 Transfer 事件\\n")
	fmt.Printf("   区块: %d\\n", vLog.BlockNumber)
	fmt.Printf("   交易: %s\\n", vLog.TxHash.Hex()[:20])
	fmt.Printf("   合约: %s\\n", vLog.Address.Hex())

	// 解析事件参数
	var transferEvent struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}

	err := parsedABI.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
	if err != nil {
		// 尝试从 topics 解析 indexed 参数
		if len(vLog.Topics) >= 3 {
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
		}
	}

	fmt.Printf("   从: %s\\n", transferEvent.From.Hex())
	fmt.Printf("   到: %s\\n", transferEvent.To.Hex())
	if transferEvent.Value != nil {
		fmt.Printf("   金额: %s\\n", transferEvent.Value.String())
	}
}
