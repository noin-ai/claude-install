package cmd

import (
	"github.com/spf13/cobra"
	"github.com/noin-ai/claude-install/internal/installer"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "测试 API 连接",
	Long:  `验证当前配置的 API 连接是否正常工作。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return installer.TestConnection()
	},
}
