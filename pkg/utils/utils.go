// Package utils 提供工具函数
package utils

import (
	"math/big"
)

// WeiToEther 将 Wei 转换为 Ether
func WeiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(
		new(big.Float).SetInt(wei),
		big.NewFloat(1e18),
	)
}

// EtherToWei 将 Ether 转换为 Wei
func EtherToWei(ether float64) *big.Int {
	eth := big.NewFloat(ether)
	wei := new(big.Float).Mul(eth, big.NewFloat(1e18))
	result, _ := wei.Int(nil)
	return result
}

// IsValidAddress 验证以太坊地址格式
func IsValidAddress(address string) bool {
	if len(address) != 42 {
		return false
	}
	if address[:2] != "0x" {
		return false
	}
	for i := 2; i < 42; i++ {
		c := address[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}
