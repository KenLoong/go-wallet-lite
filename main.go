package main

import (
	"log"
	"os"

	"go-wallet-lite/api"
	"go-wallet-lite/blockchain"

	"github.com/joho/godotenv"
)

func main() {
	// 加载配置
	// Load configuration
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// 初始化以太坊客户端
	// Initialize Ethereum client
	blockchain.InitClient()

	// 初始化 API 路由
	// Initialize API router
	r := api.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Go-Wallet-Lite is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
