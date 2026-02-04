// Package service 提供业务逻辑层
package service

import (
	"context"
	"fmt"
	"math/big"

	"go-eth-learning/pkg/ethclient"
	"go-eth-learning/pkg/transaction"
	"go-eth-learning/pkg/wallet"
)

// AccountService 账户服务
type AccountService struct {
	client *ethclient.Client
}

// NewAccountService 创建账户服务
func NewAccountService(client *ethclient.Client) *AccountService {
	return &AccountService{client: client}
}

// GetBalance 获取账户余额
func (s *AccountService) GetBalance(ctx context.Context, address string) (*big.Float, error) {
	return s.client.GetBalance(ctx, address)
}

// CreateWallet 创建新钱包
func (s *AccountService) CreateWallet() (*wallet.Wallet, error) {
	return wallet.NewWallet()
}

// ImportWallet 从私钥导入钱包
func (s *AccountService) ImportWallet(privateKey string) (*wallet.Wallet, error) {
	return wallet.FromPrivateKey(privateKey)
}

// TransactionService 交易服务
type TransactionService struct {
	client *ethclient.Client
	txMgr  *transaction.Manager
}

// NewTransactionService 创建交易服务
func NewTransactionService(client *ethclient.Client) *TransactionService {
	return &TransactionService{
		client: client,
		txMgr:  transaction.NewManager(client, client.ChainID()),
	}
}

// SendETH 发送 ETH
func (s *TransactionService) SendETH(ctx context.Context, privateKey, to string, amount *big.Int) (string, error) {
	txHash, err := s.txMgr.Transfer(ctx, privateKey, to, amount)
	if err != nil {
		return "", fmt.Errorf("发送 ETH 失败: %w", err)
	}
	return txHash, nil
}

// GetTransactionStatus 获取交易状态
func (s *TransactionService) GetTransactionStatus(ctx context.Context, txHash string) (bool, error) {
	receipt, err := s.client.WaitMined(ctx, txHash)
	if err != nil {
		return false, err
	}
	return receipt.Status == 1, nil
}

// BlockService 区块服务
type BlockService struct {
	client *ethclient.Client
}

// NewBlockService 创建区块服务
func NewBlockService(client *ethclient.Client) *BlockService {
	return &BlockService{client: client}
}

// GetLatestBlock 获取最新区块号
func (s *BlockService) GetLatestBlock(ctx context.Context) (uint64, error) {
	return s.client.GetBlockNumber(ctx)
}
