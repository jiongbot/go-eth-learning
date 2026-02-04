package utils_test

import (
	"math/big"
	"testing"

	"go-eth-learning/pkg/utils"
)

func TestWeiToEther(t *testing.T) {
	tests := []struct {
		wei      int64
		expected float64
	}{
		{1000000000000000000, 1.0},      // 1 ETH
		{500000000000000000, 0.5},       // 0.5 ETH
		{1000000000000000, 0.001},       // 0.001 ETH
		{0, 0},                          // 0 ETH
	}

	for _, tt := range tests {
		wei := big.NewInt(tt.wei)
		result := utils.WeiToEther(wei)

		// 转换为 float64 比较
		resultFloat, _ := result.Float64()
		if resultFloat != tt.expected {
			t.Errorf("WeiToEther(%d) = %f, want %f", tt.wei, resultFloat, tt.expected)
		}
	}
}

func TestEtherToWei(t *testing.T) {
	tests := []struct {
		ether    float64
		expected int64
	}{
		{1.0, 1000000000000000000},
		{0.5, 500000000000000000},
		{0.001, 1000000000000000},
		{0, 0},
	}

	for _, tt := range tests {
		result := utils.EtherToWei(tt.ether)
		if result.Int64() != tt.expected {
			t.Errorf("EtherToWei(%f) = %d, want %d", tt.ether, result.Int64(), tt.expected)
		}
	}
}

func TestIsValidAddress(t *testing.T) {
	tests := []struct {
		address string
		valid   bool
	}{
		{"0xdAC17F958D2ee523a2206206994597C13D831ec7", true},
		{"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", true},
		{"0x1234567890abcdef1234567890abcdef12345678", true},
		{"0xINVALID", false},
		{"invalid", false},
		{"0x", false},
		{"", false},
		{"0xGGGG", false}, // 无效字符
	}

	for _, tt := range tests {
		result := utils.IsValidAddress(tt.address)
		if result != tt.valid {
			t.Errorf("IsValidAddress(%s) = %v, want %v", tt.address, result, tt.valid)
		}
	}
}
