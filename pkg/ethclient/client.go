// Package ethclient 提供以太坊客户端封装
package ethclient

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client 封装以太坊客户端
type Client struct {
	client  *ethclient.Client
	chainID *big.Int
}

// New 创建新的客户端
func New(nodeURL string) (*Client, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, fmt.Errorf("连接以太坊节点失败: %w", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("获取 Chain ID 失败: %w", err)
	}

	return &Client{
		client:  client,
		chainID: chainID,
	}, nil
}

// Close 关闭客户端连接
func (c *Client) Close() {
	c.client.Close()
}

// ChainID 返回当前链 ID
func (c *Client) ChainID() *big.Int {
	return c.chainID
}

// GetBalance 查询地址余额（ETH）
func (c *Client) GetBalance(ctx context.Context, address string) (*big.Float, error) {
	addr := common.HexToAddress(address)
	balance, err := c.client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return nil, fmt.Errorf("查询余额失败: %w", err)
	}

	// 转换为 ETH
	ethBalance := new(big.Float).Quo(
		new(big.Float).SetInt(balance),
		big.NewFloat(1e18),
	)

	return ethBalance, nil
}

// GetBlockNumber 获取最新区块号
func (c *Client) GetBlockNumber(ctx context.Context) (uint64, error) {
	header, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("获取区块头失败: %w", err)
	}
	return header.Number.Uint64(), nil
}

// GetTransaction 查询交易详情
func (c *Client) GetTransaction(ctx context.Context, txHash string) (*types.Transaction, bool, error) {
	hash := common.HexToHash(txHash)
	return c.client.TransactionByHash(ctx, hash)
}

// WaitMined 等待交易被确认
func (c *Client) WaitMined(ctx context.Context, txHash string) (*types.Receipt, error) {
	hash := common.HexToHash(txHash)
	receipt, err := c.client.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("获取交易收据失败: %w", err)
	}
	return receipt, nil
}

// SendRawTransaction 发送原始交易
func (c *Client) SendRawTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.client.SendTransaction(ctx, tx)
}

// PendingNonce 获取待处理 nonce
func (c *Client) PendingNonce(ctx context.Context, address string) (uint64, error) {
	addr := common.HexToAddress(address)
	return c.client.PendingNonceAt(ctx, addr)
}

// SuggestGasPrice 获取建议 Gas 价格
func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.client.SuggestGasPrice(ctx)
}
