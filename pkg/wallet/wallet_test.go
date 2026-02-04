package wallet_test

import (
	"testing"

	"go-eth-learning/pkg/wallet"
)

func TestNewWallet(t *testing.T) {
	w, err := wallet.NewWallet()
	if err != nil {
		t.Fatalf("创建钱包失败: %v", err)
	}

	if w.GetAddressHex() == "" {
		t.Error("地址为空")
	}

	if w.GetPrivateKeyHex() == "" {
		t.Error("私钥为空")
	}

	t.Logf("地址: %s", w.GetAddressHex())
}

func TestFromPrivateKey(t *testing.T) {
	// 先创建一个新钱包获取私钥
	w1, _ := wallet.NewWallet()
	privateKey := w1.GetPrivateKeyHex()

	// 用私钥导入
	w2, err := wallet.FromPrivateKey(privateKey)
	if err != nil {
		t.Fatalf("导入钱包失败: %v", err)
	}

	// 地址应该相同
	if w1.GetAddressHex() != w2.GetAddressHex() {
		t.Error("导入的地址不匹配")
	}
}

func TestFromPrivateKey_Invalid(t *testing.T) {
	_, err := wallet.FromPrivateKey("invalid")
	if err == nil {
		t.Error("应该返回错误")
	}
}
