package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/noin-ai/claude-install/internal/system"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "检查当前系统环境",
	Long:  `检查系统信息和依赖安装状态，不进行任何安装操作。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔍 正在检查系统环境...")
		fmt.Println()

		// 系统信息
		sysInfo := system.Detect()
		fmt.Println("📊 系统信息：")
		fmt.Printf("   操作系统: %s\n", sysInfo.OS)
		fmt.Printf("   架构: %s\n", sysInfo.Arch)
		fmt.Printf("   网络环境: %s\n", sysInfo.NetworkType)
		fmt.Println()

		// 依赖检查
		fmt.Println("📦 依赖状态：")
		deps := system.CheckDependencies()

		allInstalled := true
		for name, installed := range deps {
			status := "✅ 已安装"
			if !installed {
				status = "❌ 未安装"
				allInstalled = false
			}

			version := system.GetVersion(name)
			if version != "" {
				fmt.Printf("   %s - %s (%s)\n", name, status, version)
			} else {
				fmt.Printf("   %s - %s\n", name, status)
			}
		}
		fmt.Println()

		// Claude CLI 检查
		fmt.Println("🤖 Claude 工具：")
		claudeInstalled := system.IsClaudeCLIInstalled()
		if claudeInstalled {
			version := system.GetClaudeCLIVersion()
			fmt.Printf("   ✅ Claude Code CLI - 已安装 (%s)\n", version)
		} else {
			fmt.Println("   ❌ Claude Code CLI - 未安装")
			allInstalled = false
		}
		fmt.Println()

		// 配置检查
		fmt.Println("⚙️  配置状态：")
		configExists := system.IsConfigured()
		if configExists {
			fmt.Println("   ✅ API 配置文件存在")

			// 检查配置有效性（不验证 key 真实性）
			if apiBase := system.GetConfigValue("apiBase"); apiBase != "" {
				fmt.Printf("   ✅ API Base URL: %s\n", apiBase)
			}
			if system.GetConfigValue("apiKey") != "" {
				fmt.Println("   ✅ API Key: 已配置")
			}
		} else {
			fmt.Println("   ❌ API 配置文件不存在")
			allInstalled = false
		}
		fmt.Println()

		// 总结
		if allInstalled {
			fmt.Println("✅ 环境完整，可以正常使用 Claude!")
		} else {
			fmt.Println("⚠️  环境不完整，请运行 'claude-install setup' 进行配置")
		}

		return nil
	},
}
