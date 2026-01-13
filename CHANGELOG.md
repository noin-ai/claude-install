# Changelog

所有重要更改都将记录在此文件中。

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
并遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

## [Unreleased]

### Added
- 初始版本发布
- 跨平台支持（Windows/macOS/Linux）
- 一键安装脚本（PowerShell/Bash）
- 系统环境自动检测
- 网络环境检测（国内/国外自动判断）
- Node.js/npm/Git 依赖检测
- Claude Code CLI 自动安装
- npm 国内镜像自动配置
- API 配置向导（支持 noin.ai、Anthropic 官方、自定义）
- API 连接测试功能

### Commands
- `claude-install setup` - 一键安装配置
- `claude-install check` - 检查当前环境
- `claude-install config` - 配置 API 连接
- `claude-install test` - 测试 API 连接

## [0.1.0] - 2024-01-13

### Added
- 首次发布
