package cmd

import (
	"github.com/spf13/cobra"
	"github.com/noin-ai/claude-install/internal/installer"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置或更新 API 连接信息",
	Long:  `交互式配置 API Base URL 和 API Key。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return installer.ConfigureAPI()
	},
}
