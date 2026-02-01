# Go-Wallet-Lite Makefile

# 变量定义 / Variable Definitions
BINARY_NAME=wallet-lite
MAIN_FILE=main.go

.PHONY: all build run clean tidy help

# 默认目标 / Default target
all: tidy build

# 编译项目 / Build the project
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o ./bin/$(BINARY_NAME) $(MAIN_FILE)

# 运行项目 / Run the project
run:
	@echo "Running $(BINARY_NAME)..."
	go run $(MAIN_FILE)

# 清理构建产物 / Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	@if [ -d "bin" ]; then rm -rf bin; fi

# 整理依赖 / Tidy dependencies
tidy:
	@echo "Tidying go modules..."
	go mod tidy

# 运行测试 / Run tests
test:
	@echo "Running tests..."
	go test ./... -v

# 帮助信息 / Help information
help:
	@echo "Usage:"
	@echo "  make build  - Build the binary"
	@echo "  make run    - Run the application"
	@echo "  make clean  - Remove binary and build artifacts"
	@echo "  make tidy   - Tidy go modules"
	@echo "  make test   - Run go tests"
