// Package transaction 提供交易管理功能
package transaction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Manager 交易管理器
type Manager struct {
	client  *ethclient.Client
	chainID *big.Int
}

// NewManager 创建交易管理器
func NewManager(client *ethclient.Client, chainID *big.Int) *Manager {
	return &Manager{
		client:  client,
		chainID: chainID,
	}
}

// BuildTransferTx 构建转账交易
func (m *Manager) BuildTransferTx(
	ctx context.Context,
	from string,
	to string,
	amount *big.Int,
) (*types.Transaction, error) {
	fromAddr := common.HexToAddress(from)
	toAddr := common.HexToAddress(to)

	// 获取 nonce
	nonce, err := m.client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return nil, fmt.Errorf("获取 nonce 失败: %w", err)
	}

	// 获取 Gas 价格
	gasPrice, err := m.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取 Gas 价格失败: %w", err)
	}

	// Gas 限制（ETH 转账固定 21000）
	gasLimit := uint64(21000)

	// 创建交易
	tx := types.NewTransaction(nonce, toAddr, amount, gasLimit, gasPrice, nil)

	return tx, nil
}

// SignAndSend 签名并发送交易
func (m *Manager) SignAndSend(
	ctx context.Context,
	tx *types.Transaction,
	privateKey *ecdsa.PrivateKey,
) (*types.Transaction, error) {
	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(m.chainID), privateKey)
	if err != nil {
		return nil, fmt.Errorf("签名交易失败: %w", err)
	}

	// 发送交易
	if err := m.client.SendTransaction(ctx, signedTx); err != nil {
		return nil, fmt.Errorf("发送交易失败: %w", err)
	}

	return signedTx, nil
}

// Transfer ETH 转账便捷方法
func (m *Manager) Transfer(
	ctx context.Context,
	privateKeyHex string,
	to string,
	amount *big.Int,
) (string, error) {
	// 解析私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("解析私钥失败: %w", err)
	}

	// 获取发送者地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("无法转换公钥")
	}
	from := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// 构建交易
	tx, err := m.BuildTransferTx(ctx, from, to, amount)
	if err != nil {
		return "", err
	}

	// 签名并发送
	signedTx, err := m.SignAndSend(ctx, tx, privateKey)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
