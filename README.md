# claude-install

一键配置 Claude Code CLI 开发环境的跨平台工具。

## 特性

- 🚀 一键安装配置，零门槛上手
- 🔍 智能检测系统环境和缺失依赖
- 🌐 国内网络自动加速（npm 镜像等）
- ✅ 自动验证 API 连接
- 🖥️ 支持 Windows/macOS/Linux

## 快速开始

### 安装

**从源码构建（当前推荐）：**

```bash
# 克隆并构建
git clone https://github.com/noin-ai/claude-install.git
cd claude-install
go build -o claude-install .

# 移动到 PATH（可选）
sudo mv claude-install /usr/local/bin/
```

**一键安装（Release 发布后可用）：**

```bash
# macOS/Linux
curl -fsSL https://raw.githubusercontent.com/noin-ai/claude-install/main/scripts/install.sh | bash

# Windows PowerShell
irm https://raw.githubusercontent.com/noin-ai/claude-install/main/scripts/install.ps1 | iex
```

### 使用

```bash
# 一键配置 Claude 环境
claude-install setup

# 检查当前环境状态
claude-install check

# 配置或更新 API 连接
claude-install config

# 测试 API 连接
claude-install test
```

## 工作流程

1. **系统检测** - 自动识别操作系统和架构
2. **依赖检查** - 检测 Node.js、npm、Git 等必要工具
3. **智能安装** - 自动安装缺失的依赖（国内网络自动加速）
4. **Claude CLI** - 安装 Claude Code CLI
5. **配置向导** - 交互式配置 API Key
6. **连接测试** - 验证配置是否正常

## 开发

```bash
# 克隆项目
git clone https://github.com/noin-ai/claude-install.git
cd claude-install

# 安装依赖
go mod tidy

# 运行
go run main.go setup

# 构建
go build -o claude-install main.go
```

## 项目结构

```
claude-install/
├── main.go                     # 入口文件
├── cmd/                        # 命令行命令
│   ├── root.go                 # 根命令
│   ├── setup.go                # 一键安装
│   ├── check.go                # 环境检查
│   ├── config.go               # 配置管理
│   └── test.go                 # 连接测试
├── internal/
│   ├── system/                 # 系统检测
│   │   ├── detect.go           # 系统信息检测
│   │   ├── dependency.go       # 依赖检查
│   │   └── config.go           # 配置读取
│   └── installer/              # 安装器
│       ├── dependency.go       # 依赖安装
│       ├── config.go           # 配置管理
│       └── test.go             # 连接测试
└── README.md
```

## 许可证

MIT

## 贡献

欢迎提交 Issue 和 Pull Request！
