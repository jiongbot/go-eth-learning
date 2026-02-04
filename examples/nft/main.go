// NFT 操作示例
package main

import (
	"fmt"
	"log"

	"go-eth-learning/internal/config"
)

func main() {
	fmt.Println("🖼️ NFT 操作示例\\n")

	_, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	fmt.Println("=== NFT 标准 (ERC721) ===")
	fmt.Println()
	fmt.Println("主要方法:")
	fmt.Println("  - balanceOf(address): 查询持有数量")
	fmt.Println("  - ownerOf(tokenId): 查询 NFT 所有者")
	fmt.Println("  - transferFrom(from, to, tokenId): 转移 NFT")
	fmt.Println("  - approve(to, tokenId): 授权转移")
	fmt.Println("  - setApprovalForAll(operator, approved): 批量授权")
	fmt.Println()
	fmt.Println("元数据:")
	fmt.Println("  - tokenURI(tokenId): 获取 NFT 元数据链接")
	fmt.Println("  - name(): 集合名称")
	fmt.Println("  - symbol(): 集合符号")
	fmt.Println()
	fmt.Println("=== NFT 标准 (ERC1155) ===")
	fmt.Println()
	fmt.Println("多代币标准，一个合约支持多种 NFT")
	fmt.Println("  - balanceOf(address, tokenId): 查询特定代币余额")
	fmt.Println("  - balanceOfBatch: 批量查询")
	fmt.Println("  - safeTransferFrom: 安全转移")
	fmt.Println()
	fmt.Println("=== 流行 NFT 合约 ===")
	fmt.Println("  - CryptoPunks: 0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB")
	fmt.Println("  - BAYC: 0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D")
	fmt.Println("  - Azuki: 0xED5AF388653567Af2F388E6224dC7C4b3241C544")
	fmt.Println()
	fmt.Println("✅ NFT 示例完成!")
	fmt.Println("提示: 使用 abigen 生成完整合约绑定后可实现完整交互")
}
