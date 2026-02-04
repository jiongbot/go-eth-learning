// cmd/wallet 钱包管理工具
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-eth-learning/pkg/wallet"
)

func main() {
	fmt.Println("⛓️🐹 以太坊钱包管理工具\\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		showMenu()
		choice := getInput(reader, "请选择操作: ")

		switch choice {
		case "1":
			createWallet()
		case "2":
			importWallet(reader)
		case "3":
			fmt.Println("再见!")
			return
		default:
			fmt.Println("无效选择，请重试")
		}

		fmt.Println()
	}
}

func showMenu() {
	fmt.Println("=== 钱包管理 ===")
	fmt.Println("1. 创建新钱包")
	fmt.Println("2. 导入钱包")
	fmt.Println("3. 退出")
	fmt.Println()
}

func getInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\\n')
	return strings.TrimSpace(input)
}

func createWallet() {
	fmt.Println("\\n创建新钱包...")

	w, err := wallet.NewWallet()
	if err != nil {
		fmt.Printf("创建失败: %v\\n", err)
		return
	}

	fmt.Println("✅ 钱包创建成功!")
	fmt.Printf("   地址: %s\\n", w.GetAddressHex())
	fmt.Printf("   私钥: %s\\n", w.GetPrivateKeyHex())
	fmt.Println("   ⚠️  请安全保存私钥，不要泄露!")
}

func importWallet(reader *bufio.Reader) {
	fmt.Println("\\n导入钱包...")

	privateKey := getInput(reader, "请输入私钥 (hex, 0x 可选): ")
	privateKey = strings.TrimPrefix(privateKey, "0x")

	w, err := wallet.FromPrivateKey(privateKey)
	if err != nil {
		fmt.Printf("导入失败: %v\\n", err)
		return
	}

	fmt.Println("✅ 钱包导入成功!")
	fmt.Printf("   地址: %s\\n", w.GetAddressHex())
}
