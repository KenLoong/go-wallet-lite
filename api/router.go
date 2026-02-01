package api

import (
	"net/http"

	"go-wallet-lite/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化 Gin 路由
// SetupRouter initializes Gin router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		// 钱包创建
		// Wallet creation
		v1.POST("/wallet/create", func(c *gin.Context) {
			wallet, err := service.CreateWallet()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, wallet)
		})

		// ETH 余额查询
		// ETH balance query
		v1.GET("/balance/:address", func(c *gin.Context) {
			address := c.Param("address")
			balance, err := service.GetBalance(address)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"address": address, "balance": balance, "unit": "ETH"})
		})

		// ERC-20 余额查询
		// ERC-20 balance query
		v1.GET("/balance/token/:address", func(c *gin.Context) {
			userAddr := c.Param("address")
			tokenAddr := c.Query("contract") // 从 query 参数获取合约地址
			if tokenAddr == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "contract address is required"})
				return
			}

			balance, err := service.GetTokenBalance(tokenAddr, userAddr)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"address": userAddr, "token_address": tokenAddr, "balance": balance})
		})

		// 转账交易
		// Transfer transaction
		v1.POST("/transfer", func(c *gin.Context) {
			var req service.TransferRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			txHash, err := service.TransferETH(req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"tx_hash": txHash})
		})

		// 交易状态检查
		// Transaction status check
		v1.GET("/tx/:hash", func(c *gin.Context) {
			hash := c.Param("hash")
			status, err := service.GetTransactionStatus(hash)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"tx_hash": hash, "confirmed": status})
		})
	}

	return r
}
