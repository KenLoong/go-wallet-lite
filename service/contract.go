package service

import (
	"fmt"

	"go-wallet-lite/blockchain"
	"go-wallet-lite/contracts"

	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/common"
)

// GetTokenBalance 查询指定地址的 ERC-20 代币余额
// GetTokenBalance fetches the ERC-20 token balance for a given address
func GetTokenBalance(tokenAddress, userAddress string) (string, error) {
	tokenAddr := common.HexToAddress(tokenAddress)
	userAddr := common.HexToAddress(userAddress)

	// 使用 abigen 生成的绑定创建合约实例
	// Create contract instance using abigen generated binding
	instance, err := contracts.NewERC20(tokenAddr, blockchain.Client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %v", err)
	}

	// 调用 BalanceOf 方法 (强类型)
	// Call BalanceOf method (Strongly typed)
	balance, err := instance.BalanceOf(nil, userAddr)
	if err != nil {
		return "", fmt.Errorf("failed to call balanceOf: %v", err)
	}

	// 调用 Decimals 方法
	// Call Decimals method
	decimals, err := instance.Decimals(nil)
	if err != nil {
		// 容错处理
		decimals = 18
	}

	// 使用 shopspring/decimal 处理数值转换 (更优雅)
	// Use shopspring/decimal for numeric conversion (more elegant)
	d := decimal.NewFromBigInt(balance, -int32(decimals))

	return d.StringFixed(int32(decimals)), nil
}
