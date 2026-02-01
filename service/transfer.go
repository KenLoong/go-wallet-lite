package service

import (
	"context"
	"fmt"
	"math/big"

	"go-wallet-lite/blockchain"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// TransferRequest 转账请求结构体
// TransferRequest structure for transfers
type TransferRequest struct {
	PrivateKey string `json:"private_key"`
	ToAddress  string `json:"to_address"`
	Amount     string `json:"amount"` // ETH 单位
}

// TransferETH 构造、签名并发送 ETH 转账交易 (EIP-1559)
// TransferETH constructs, signs and sends an ETH transfer (EIP-1559)
func TransferETH(req TransferRequest) (string, error) {
	privateKey, err := crypto.HexToECDSA(req.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 获取 Nonce
	// Get Nonce
	nonce, err := blockchain.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// 将 ETH 转换为 Wei
	// Convert ETH to Wei
	amountETH, ok := new(big.Float).SetString(req.Amount)
	if !ok {
		return "", fmt.Errorf("invalid amount")
	}
	weiValue := new(big.Int)
	amountETH.Mul(amountETH, big.NewFloat(1e18)).Int(weiValue)

	// 获取网络建议费用 (EIP-1559)
	// Suggest gas fees (EIP-1559)
	gasTipCap, err := blockchain.Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas tip cap: %v", err)
	}

	header, err := blockchain.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to get latest block header: %v", err)
	}

	// MaxFeePerGas = BaseFee * 2 + MaxPriorityFeePerGas
	baseFee := header.BaseFee
	gasFeeCap := new(big.Int).Add(
		new(big.Int).Mul(baseFee, big.NewInt(2)),
		gasTipCap,
	)

	gasLimit := uint64(21000) // 标准 ETH 转账
	toAddress := common.HexToAddress(req.ToAddress)

	// 获取 ChainID
	// Get ChainID
	chainID, err := blockchain.Client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %v", err)
	}

	// 构造交易负载
	// Construct dynamic fee transaction
	txData := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     weiValue,
		Data:      nil,
	}

	tx := types.NewTx(txData)

	// 签名交易
	// Sign transaction
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign tx: %v", err)
	}

	// 广播交易
	// Broadcast transaction
	err = blockchain.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}

// GetTransactionStatus 检查交易确认状态
// GetTransactionStatus checks the transaction confirmation status
func GetTransactionStatus(txHash string) (bool, error) {
	hash := common.HexToHash(txHash)
	receipt, err := blockchain.Client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		if err == ethereum.NotFound {
			return false, nil // 交易尚未确认
		}
		return false, err
	}

	return receipt.Status == types.ReceiptStatusSuccessful, nil
}
