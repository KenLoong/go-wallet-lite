package service

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// TestCreateWallet 测试离线钱包生成
func TestCreateWallet(t *testing.T) {
	wallet, err := CreateWallet()
	assert.NoError(t, err)

	// 1. 检查地址是否为空
	assert.NotEmpty(t, wallet.Address)

	// 2. 检查地址格式 (以 0x 开头，且长度正确)
	assert.True(t, common.IsHexAddress(wallet.Address), "Invalid address format: %s", wallet.Address)

	// 3. 检查私钥是否为空
	assert.NotEmpty(t, wallet.PrivateKey)

	// 4. 检查私钥格式 (通常是 0x 开头的 66 位 16 进制字符串)
	assert.Len(t, wallet.PrivateKey, 66, "Unexpected private key length")

	t.Logf("Generated Wallet Address: %s", wallet.Address)
}
