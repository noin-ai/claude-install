# claude-install Makefile
# 跨平台构建脚本

BINARY_NAME=claude-install
VERSION?=0.1.0
BUILD_DIR=dist
GO=/usr/local/go/bin/go

# 构建信息
LDFLAGS=-ldflags "-s -w -X main.Version=$(VERSION)"

.PHONY: all clean build-all build-darwin build-linux build-windows

all: build-all

clean:
	rm -rf $(BUILD_DIR)

# 创建输出目录
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# 构建所有平台
build-all: $(BUILD_DIR) build-darwin build-linux build-windows
	@echo "✅ 所有平台构建完成！"
	@ls -la $(BUILD_DIR)/

# macOS (Intel + Apple Silicon)
build-darwin: $(BUILD_DIR)
	@echo "🍎 构建 macOS amd64..."
	GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	@echo "🍎 构建 macOS arm64..."
	GOOS=darwin GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .

# Linux (amd64 + arm64)
build-linux: $(BUILD_DIR)
	@echo "🐧 构建 Linux amd64..."
	GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	@echo "🐧 构建 Linux arm64..."
	GOOS=linux GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 .

# Windows (amd64 + arm64)
build-windows: $(BUILD_DIR)
	@echo "🪟 构建 Windows amd64..."
	GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	@echo "🪟 构建 Windows arm64..."
	GOOS=windows GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe .

# 本地开发构建
build:
	$(GO) build -o $(BINARY_NAME) .

# 运行测试
test:
	$(GO) test ./...

# 安装到本地
install: build
	mv $(BINARY_NAME) /usr/local/bin/
