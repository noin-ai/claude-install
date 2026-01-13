package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "claude-install",
	Short: "一键配置 Claude Code CLI 开发环境",
	Long: `claude-install - Claude 环境配置工具

自动检测系统环境，安装必要依赖，配置 API 连接。
专为新手设计，零配置难度。

使用方式：
  claude-install setup    一键安装配置
  claude-install check    检查当前环境
  claude-install config   配置 API 连接
  claude-install test     测试 API 连接`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("欢迎使用 claude-install!")
		fmt.Println("使用 'claude-install setup' 开始一键配置")
		fmt.Println("使用 'claude-install --help' 查看所有命令")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(testCmd)
}
