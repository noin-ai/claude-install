package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/noin-ai/claude-install/internal/installer"
	"github.com/noin-ai/claude-install/internal/system"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "一键安装配置 Claude 开发环境",
	Long: `自动执行以下步骤：
1. 检测系统环境
2. 安装缺失的依赖（Node.js, npm 等）
3. 安装 Claude Code CLI
4. 配置 API 连接
5. 验证安装`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🚀 开始配置 Claude 开发环境...")
		fmt.Println()

		// 1. 系统检测
		sysInfo := system.Detect()
		fmt.Printf("📊 系统信息：\n")
		fmt.Printf("   操作系统: %s\n", sysInfo.OS)
		fmt.Printf("   架构: %s\n", sysInfo.Arch)
		fmt.Printf("   网络环境: %s\n", sysInfo.NetworkType)
		fmt.Println()

		// 2. 环境检查
		fmt.Println("🔍 检查依赖环境...")
		deps := system.CheckDependencies()

		missing := []string{}
		for name, installed := range deps {
			status := "✅"
			if !installed {
				status = "❌"
				missing = append(missing, name)
			}
			fmt.Printf("   %s %s\n", status, name)
		}
		fmt.Println()

		// 3. 安装缺失依赖
		if len(missing) > 0 {
			fmt.Printf("📦 需要安装 %d 个依赖...\n", len(missing))
			if err := installer.InstallDependencies(missing, sysInfo); err != nil {
				return fmt.Errorf("安装依赖失败: %w", err)
			}
			fmt.Println("✅ 依赖安装完成")
			fmt.Println()
		}

		// 4. 安装 Claude Code CLI
		fmt.Println("📥 安装 Claude Code CLI...")
		if err := installer.InstallClaudeCLI(sysInfo); err != nil {
			return fmt.Errorf("安装 Claude CLI 失败: %w", err)
		}
		fmt.Println("✅ Claude Code CLI 安装完成")
		fmt.Println()

		// 5. 配置向导
		fmt.Println("⚙️  开始配置...")
		if err := installer.ConfigureAPI(); err != nil {
			return fmt.Errorf("配置失败: %w", err)
		}
		fmt.Println("✅ 配置完成")
		fmt.Println()

		// 6. 测试连接
		fmt.Println("🧪 测试 API 连接...")
		if err := installer.TestConnection(); err != nil {
			return fmt.Errorf("连接测试失败: %w", err)
		}
		fmt.Println()

		fmt.Println("🎉 恭喜！Claude 开发环境配置成功！")
		fmt.Println()
		fmt.Println("下一步：")
		fmt.Println("  1. 在终端输入 'claude' 开始使用")
		fmt.Println("  2. 使用 'claude-install test' 随时测试连接")

		return nil
	},
}
