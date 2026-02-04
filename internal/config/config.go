// Package config 提供配置管理
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config 应用配置
type Config struct {
	// 以太坊节点配置
	EthNodeURL string
	ChainID    int64

	// 钱包配置
	PrivateKey string

	// 合约地址
	ContractAddresses map[string]string
}

// Load 从环境变量加载配置
func Load() (*Config, error) {
	// 加载 .env 文件（如果存在）
	_ = godotenv.Load()

	return &Config{
		EthNodeURL: getEnv("ETH_NODE_URL", "https://sepolia.infura.io/v3/YOUR_KEY"),
		ChainID:    11155111, // Sepolia
		PrivateKey: getEnv("PRIVATE_KEY", ""),
		ContractAddresses: map[string]string{
			"USDT": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
