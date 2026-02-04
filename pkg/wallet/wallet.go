// Package wallet 提供以太坊钱包功能
package wallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Wallet 以太坊钱包
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    common.Address
}

// NewWallet 创建新钱包
func NewWallet() (*Wallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("生成私钥失败: %w", err)
	}

	return walletFromPrivateKey(privateKey), nil
}

// FromPrivateKey 从私钥字符串加载钱包
func FromPrivateKey(hexKey string) (*Wallet, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, fmt.Errorf("解析私钥失败: %w", err)
	}

	return walletFromPrivateKey(privateKey), nil
}

// walletFromPrivateKey 从私钥创建钱包
func walletFromPrivateKey(privateKey *ecdsa.PrivateKey) *Wallet {
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKeyECDSA,
		Address:    address,
	}
}

// GetPrivateKeyHex 获取十六进制私钥
func (w *Wallet) GetPrivateKeyHex() string {
	return fmt.Sprintf("%x", crypto.FromECDSA(w.PrivateKey))
}

// GetAddressHex 获取地址字符串
func (w *Wallet) GetAddressHex() string {
	return w.Address.Hex()
}
