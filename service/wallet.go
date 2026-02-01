package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"

	"go-wallet-lite/blockchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// WalletInfo 钱包信息结构体
// WalletInfo structure for wallet information
type WalletInfo struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
}

// CreateWallet 离线生成新的钱包地址与私钥
// CreateWallet generates a new wallet address and private key offline
func CreateWallet() (*WalletInfo, error) {
	// 生成私钥
	// Generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	// 导出私钥字符串
	// Export private key string
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)

	// 获取公钥并转换为地址
	// Get public key and convert it to address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &WalletInfo{
		Address:    address,
		PrivateKey: privateKeyHex,
	}, nil
}

// GetBalance 查询指定地址的 ETH 余额
// GetBalance fetches the ETH balance for a given address
func GetBalance(address string) (string, error) {
	account := common.HexToAddress(address)
	balance, err := blockchain.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "", fmt.Errorf("failed to fetch balance: %v", err)
	}

	// 将 Wei 转换为 ETH (1 ETH = 10^18 Wei)
	// Convert Wei to ETH (1 ETH = 10^18 Wei)
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue.Text('f', 18), nil
}
