# 贡献指南

感谢你对 claude-install 的关注！欢迎提交 Issue 和 Pull Request。

## 开发环境

### 前置要求

- Go 1.21 或更高版本
- Git

### 本地开发

```bash
# 克隆项目
git clone https://github.com/noin-ai/claude-install.git
cd claude-install

# 安装依赖
go mod tidy

# 本地构建
go build -o claude-install .

# 运行测试
go test ./...

# 本地测试
./claude-install check
./claude-install --help
```

### 跨平台构建

```bash
# macOS
GOOS=darwin GOARCH=amd64 go build -o dist/claude-install-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -o dist/claude-install-darwin-arm64 .

# Linux
GOOS=linux GOARCH=amd64 go build -o dist/claude-install-linux-amd64 .
GOOS=linux GOARCH=arm64 go build -o dist/claude-install-linux-arm64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o dist/claude-install-windows-amd64.exe .
GOOS=windows GOARCH=arm64 go build -o dist/claude-install-windows-arm64.exe .
```

## 提交规范

### Commit 信息格式

```
<type>: <description>

[optional body]
```

**类型 (type):**
- `feat`: 新功能
- `fix`: 修复 bug
- `docs`: 文档更新
- `style`: 代码格式（不影响功能）
- `refactor`: 重构
- `test`: 测试相关
- `chore`: 构建/工具相关

**示例:**
```
feat: 添加 Windows 自动安装 Node.js 功能
fix: 修复国内网络检测误判问题
docs: 更新 README 安装说明
```

## Pull Request 流程

1. Fork 本仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'feat: 添加某功能'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

### PR 检查清单

- [ ] 代码通过 `go build` 编译
- [ ] 代码通过 `go test ./...` 测试
- [ ] 更新了相关文档（如需要）
- [ ] Commit 信息符合规范

## 问题反馈

### 报告 Bug

请在 Issue 中包含：
- 操作系统和版本
- Go 版本
- 复现步骤
- 期望行为 vs 实际行为
- 错误信息（如有）

### 功能建议

请描述：
- 你想要的功能
- 为什么需要这个功能
- 可能的实现方式（可选）

## 代码规范

- 使用 `gofmt` 格式化代码
- 遵循 Go 标准命名规范
- 添加必要的注释（特别是公开函数）
- 错误信息使用中文，面向普通用户

## 联系我们

- 官网: [https://noin.ai](https://noin.ai)
- Issue: [GitHub Issues](https://github.com/noin-ai/claude-install/issues)
