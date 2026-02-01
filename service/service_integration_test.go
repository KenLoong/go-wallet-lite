package service

import (
	"os"
	"testing"

	"go-wallet-lite/blockchain"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestMain 用于执行测试前的环境初始化
func TestMain(m *testing.M) {
	// 加载 .env
	_ = godotenv.Load("../.env")

	// 初始化客户端 (部分测试需要连网)
	if os.Getenv("RPC_URL") != "" {
		blockchain.InitClient()
	}

	// 执行测试
	code := m.Run()
	os.Exit(code)
}

// TestGetBalance 测试 ETH 余额查询 (集成测试)
func TestGetBalance(t *testing.T) {
	if os.Getenv("RPC_URL") == "" {
		t.Skip("Skipping GetBalance test: RPC_URL not set")
	}

	// 使用一个已知的测试地址 (比如 Vitalik 的地址)
	address := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"
	balance, err := GetBalance(address)

	// 使用 assert 进行断言
	assert.NoError(t, err)
	assert.NotEmpty(t, balance)

	t.Logf("Balance for %s: %s ETH", address, balance)
}

// TestGetTokenBalance 测试代币余额查询 (集成测试)
func TestGetTokenBalance(t *testing.T) {
	if os.Getenv("RPC_URL") == "" {
		t.Skip("Skipping GetTokenBalance test: RPC_URL not set")
	}

	// 正确的 Sepolia LINK Token 地址
	tokenAddr := "0x779877A7B0D9E8603169DdbD7836e478b4624789"
	userAddr := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"

	balance, err := GetTokenBalance(tokenAddr, userAddr)

	// 严格断言：必须没有错误且余额不为空
	assert.NoError(t, err)
	assert.NotEmpty(t, balance)

	t.Logf("Token Balance: %s", balance)
}
