package ethclient_test

import (
	"context"
	"testing"
	"time"

	"go-eth-learning/pkg/ethclient"
)

func TestNew(t *testing.T) {
	// 使用公共节点测试
	client, err := ethclient.New("https://cloudflare-eth.com")
	if err != nil {
		t.Skipf("跳过测试，无法连接节点: %v", err)
	}
	defer client.Close()

	if client.ChainID().Int64() == 0 {
		t.Error("Chain ID 不应该为 0")
	}
}

func TestGetBalance(t *testing.T) {
	client, err := ethclient.New("https://cloudflare-eth.com")
	if err != nil {
		t.Skipf("跳过测试，无法连接节点: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 使用已知地址测试
	address := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	balance, err := client.GetBalance(ctx, address)
	if err != nil {
		t.Errorf("查询余额失败: %v", err)
	}

	if balance == nil {
		t.Error("余额不应该为 nil")
	}

	t.Logf("余额: %f ETH", balance)
}

func TestGetBlockNumber(t *testing.T) {
	client, err := ethclient.New("https://cloudflare-eth.com")
	if err != nil {
		t.Skipf("跳过测试，无法连接节点: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blockNum, err := client.GetBlockNumber(ctx)
	if err != nil {
		t.Errorf("获取区块号失败: %v", err)
	}

	if blockNum == 0 {
		t.Error("区块号不应该为 0")
	}

	t.Logf("最新区块: %d", blockNum)
}
