#!/bin/bash
# Claude Code CLI 一键安装脚本 (Linux/macOS)
# 使用方式: curl -fsSL https://noin.ai/install.sh | bash

set -e

# 配置
BINARY_NAME="claude-install"
DOWNLOAD_BASE="https://github.com/noin-ai/claude-install/releases/latest/download"
INSTALL_DIR="$HOME/.claude-install"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
NC='\033[0m' # No Color

print_step() {
    echo -e "\n${CYAN}$1${NC}"
}

print_success() {
    echo -e "${GREEN}$1${NC}"
}

print_error() {
    echo -e "${RED}$1${NC}"
}

# 检测操作系统
detect_os() {
    case "$(uname -s)" in
        Darwin*) echo "darwin" ;;
        Linux*)  echo "linux" ;;
        *)       echo "unknown" ;;
    esac
}

# 检测架构
detect_arch() {
    case "$(uname -m)" in
        x86_64)  echo "amd64" ;;
        amd64)   echo "amd64" ;;
        arm64)   echo "arm64" ;;
        aarch64) echo "arm64" ;;
        *)       echo "amd64" ;;
    esac
}

# 添加到 PATH
add_to_path() {
    local shell_rc=""

    if [ -n "$ZSH_VERSION" ] || [ -f "$HOME/.zshrc" ]; then
        shell_rc="$HOME/.zshrc"
    elif [ -n "$BASH_VERSION" ] || [ -f "$HOME/.bashrc" ]; then
        shell_rc="$HOME/.bashrc"
    elif [ -f "$HOME/.profile" ]; then
        shell_rc="$HOME/.profile"
    fi

    if [ -n "$shell_rc" ]; then
        if ! grep -q "$INSTALL_DIR" "$shell_rc" 2>/dev/null; then
            echo "" >> "$shell_rc"
            echo "# Claude Install" >> "$shell_rc"
            echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$shell_rc"
            print_success "   已添加到 $shell_rc"
        fi
    fi

    export PATH="$PATH:$INSTALL_DIR"
}

# 显示 Banner
echo -e "${MAGENTA}"
cat << 'EOF'

   ██████╗██╗      █████╗ ██╗   ██╗██████╗ ███████╗
  ██╔════╝██║     ██╔══██╗██║   ██║██╔══██╗██╔════╝
  ██║     ██║     ███████║██║   ██║██║  ██║█████╗
  ██║     ██║     ██╔══██║██║   ██║██║  ██║██╔══╝
  ╚██████╗███████╗██║  ██║╚██████╔╝██████╔╝███████╗
   ╚═════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝

  Claude Code CLI 一键安装工具
  By noin.ai

EOF
echo -e "${NC}"

# 检测系统
print_step "1. 检测系统环境..."
OS=$(detect_os)
ARCH=$(detect_arch)

if [ "$OS" = "unknown" ]; then
    print_error "   不支持的操作系统"
    exit 1
fi

if [ "$OS" = "darwin" ]; then
    print_success "   macOS $ARCH"
else
    print_success "   Linux $ARCH"
fi

# 创建安装目录
print_step "2. 创建安装目录..."
mkdir -p "$INSTALL_DIR"
print_success "   $INSTALL_DIR"

# 下载二进制文件
print_step "3. 下载安装程序..."
FILE_NAME="${BINARY_NAME}-${OS}-${ARCH}"
DOWNLOAD_URL="${DOWNLOAD_BASE}/${FILE_NAME}"
TARGET_PATH="${INSTALL_DIR}/${BINARY_NAME}"

echo "   正在下载: $DOWNLOAD_URL"

if command -v curl &> /dev/null; then
    if curl -fsSL "$DOWNLOAD_URL" -o "$TARGET_PATH"; then
        print_success "   下载完成!"
    else
        print_error "   下载失败"
        echo ""
        echo "请手动下载并安装:"
        echo "   $DOWNLOAD_URL"
        exit 1
    fi
elif command -v wget &> /dev/null; then
    if wget -q "$DOWNLOAD_URL" -O "$TARGET_PATH"; then
        print_success "   下载完成!"
    else
        print_error "   下载失败"
        exit 1
    fi
else
    print_error "   需要 curl 或 wget"
    exit 1
fi

# 设置执行权限
chmod +x "$TARGET_PATH"

# 配置环境变量
print_step "4. 配置环境变量..."
add_to_path

# 运行安装向导
print_step "5. 运行安装向导..."
echo ""

"$TARGET_PATH" setup

echo -e "${GREEN}"
cat << 'EOF'

============================================
  安装完成!

  重新打开终端后，可以使用:
    claude-install check  - 检查环境状态
    claude-install config - 更新配置
    claude-install test   - 测试连接

  需要帮助? 访问: https://noin.ai/help
============================================

EOF
echo -e "${NC}"
