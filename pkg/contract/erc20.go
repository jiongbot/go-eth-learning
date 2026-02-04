package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// ERC20 标准接口 ABI（简化版）
const ERC20ABI = `[
	{
		"constant": true,
		"inputs": [],
		"name": "name",
		"outputs": [{"name": "", "type": "string"}],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "symbol",
		"outputs": [{"name": "", "type": "string"}],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "decimals",
		"outputs": [{"name": "", "type": "uint8"}],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [{"name": "_owner", "type": "address"}],
		"name": "balanceOf",
		"outputs": [{"name": "balance", "type": "uint256"}],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{"name": "_to", "type": "address"},
			{"name": "_value", "type": "uint256"}
		],
		"name": "transfer",
		"outputs": [{"name": "", "type": "bool"}],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{"indexed": true, "name": "from", "type": "address"},
			{"indexed": true, "name": "to", "type": "address"},
			{"indexed": false, "name": "value", "type": "uint256"}
		],
		"name": "Transfer",
		"type": "event"
	}
]`

// ParseERC20ABI 解析 ERC20 ABI
func ParseERC20ABI() (abi.ABI, error) {
	return abi.JSON(strings.NewReader(ERC20ABI))
}

// ERC20Contract ERC20 合约封装
type ERC20Contract struct {
	Address common.Address
	ABI     abi.ABI
	// Client  *ethclient.Client // 实际使用时添加
}

// NewERC20Contract 创建 ERC20 合约实例
func NewERC20Contract(address string) (*ERC20Contract, error) {
	parsedABI, err := ParseERC20ABI()
	if err != nil {
		return nil, err
	}

	return &ERC20Contract{
		Address: common.HexToAddress(address),
		ABI:     parsedABI,
	}, nil
}
