package blockchain

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// Client 全局以太坊客户端实例
// Global Ethereum client instance
var Client *ethclient.Client

// InitClient 初始化以太坊 RPC 客户端
// InitClient initializes the Ethereum RPC client
func InitClient() {
	// 加载 .env 文件
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatal("RPC_URL is not set in .env or environment")
	}

	// 连接到以太坊节点
	// Connect to the Ethereum node
	// 本质上是创建一个client，然后通过client去调用以太坊的接口
	Client, err = ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	log.Printf("Successfully connected to Ethereum RPC: %s", rpcURL)
}
