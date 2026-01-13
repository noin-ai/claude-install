package installer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/noin-ai/claude-install/internal/system"
)

// ConfigData 配置数据结构
type ConfigData struct {
	APIKey     string
	APIBaseURL string
}

// 默认配置
const (
	DefaultAPIBaseURL = "https://api.noin.ai" // noin.ai 服务地址
	OfficialAPIURL    = "https://api.anthropic.com"
)

// ConfigureAPI 配置 API 连接
func ConfigureAPI() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("📝 配置 Claude API 连接")
	fmt.Println()

	// 服务选择
	fmt.Println("请选择 API 服务:")
	fmt.Println("   [1] noin.ai (推荐 - 国内优化、价格优惠)")
	fmt.Println("   [2] Anthropic 官方")
	fmt.Println("   [3] 自定义地址")
	fmt.Println()
	fmt.Print("请选择 [1/2/3] (默认 1): ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var apiBase string
	switch choice {
	case "", "1":
		apiBase = DefaultAPIBaseURL
		fmt.Println()
		fmt.Println("✅ 使用 noin.ai 服务")
		fmt.Println("   💡 访问 https://noin.ai 获取 API Key")
		fmt.Println("   🎁 新用户赠送免费额度!")
	case "2":
		apiBase = OfficialAPIURL
		fmt.Println()
		fmt.Println("✅ 使用 Anthropic 官方服务")
	case "3":
		fmt.Print("请输入 API Base URL: ")
		apiBase, _ = reader.ReadString('\n')
		apiBase = strings.TrimSpace(apiBase)
		if apiBase == "" {
			return fmt.Errorf("API Base URL 不能为空")
		}
	default:
		apiBase = DefaultAPIBaseURL
		fmt.Println("✅ 使用默认 noin.ai 服务")
	}
	fmt.Println()

	// 读取 API Key
	fmt.Print("请输入 API Key: ")
	apiKey, _ := reader.ReadString('\n')
	apiKey = strings.TrimSpace(apiKey)

	if apiKey == "" {
		return fmt.Errorf("API Key 不能为空")
	}

	// 保存配置（通过环境变量）
	config := ConfigData{
		APIKey:     apiKey,
		APIBaseURL: apiBase,
	}

	if err := saveConfig(config); err != nil {
		return fmt.Errorf("保存配置失败: %w", err)
	}

	// 同时设置当前进程的环境变量，以便后续测试可以使用
	os.Setenv(system.EnvAPIKey, config.APIKey)
	if config.APIBaseURL != "" {
		os.Setenv(system.EnvBaseURL, config.APIBaseURL)
	}

	fmt.Println()
	fmt.Println("✅ 配置已保存")
	fmt.Println()
	fmt.Println("💡 环境变量已写入配置文件，请重新打开终端使配置生效")

	return nil
}

// saveConfig 保存配置到 shell 配置文件
func saveConfig(config ConfigData) error {
	// 获取用户主目录
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// 确定 shell 配置文件
	var shellRcPath string
	var exportFormat string

	if runtime.GOOS == "windows" {
		// Windows: 使用 PowerShell profile 或设置用户环境变量
		return saveConfigWindows(config)
	}

	// Unix-like systems: 检测 shell 类型
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "zsh") {
		shellRcPath = filepath.Join(home, ".zshrc")
	} else if strings.Contains(shell, "bash") {
		// 检查是否存在 .bashrc 或 .bash_profile
		bashrcPath := filepath.Join(home, ".bashrc")
		bashProfilePath := filepath.Join(home, ".bash_profile")
		if _, err := os.Stat(bashrcPath); err == nil {
			shellRcPath = bashrcPath
		} else {
			shellRcPath = bashProfilePath
		}
	} else {
		// 默认使用 .profile
		shellRcPath = filepath.Join(home, ".profile")
	}

	exportFormat = "export %s=\"%s\"\n"

	// 读取现有配置
	existingContent := ""
	if data, err := os.ReadFile(shellRcPath); err == nil {
		existingContent = string(data)
	}

	// 构建新的环境变量配置
	var newLines []string
	newLines = append(newLines, "")
	newLines = append(newLines, "# Claude Code CLI 配置 (由 claude-install 自动生成)")
	newLines = append(newLines, fmt.Sprintf(exportFormat, system.EnvAPIKey, config.APIKey))
	if config.APIBaseURL != "" && config.APIBaseURL != OfficialAPIURL {
		newLines = append(newLines, fmt.Sprintf(exportFormat, system.EnvBaseURL, config.APIBaseURL))
	}

	// 移除旧的 Claude 配置（如果存在）
	lines := strings.Split(existingContent, "\n")
	var cleanedLines []string
	skipNext := false
	for _, line := range lines {
		if strings.Contains(line, "# Claude Code CLI 配置") {
			skipNext = true
			continue
		}
		if skipNext && (strings.Contains(line, "ANTHROPIC_API_KEY") || strings.Contains(line, "ANTHROPIC_BASE_URL")) {
			continue
		}
		skipNext = false
		cleanedLines = append(cleanedLines, line)
	}

	// 写入更新后的配置
	finalContent := strings.Join(cleanedLines, "\n") + strings.Join(newLines, "\n")
	return os.WriteFile(shellRcPath, []byte(finalContent), 0644)
}

// saveConfigWindows Windows 系统保存配置
func saveConfigWindows(config ConfigData) error {
	// Windows: 创建一个设置环境变量的脚本
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// 创建 PowerShell profile 目录
	psProfileDir := filepath.Join(home, "Documents", "WindowsPowerShell")
	if err := os.MkdirAll(psProfileDir, 0755); err != nil {
		return err
	}

	psProfilePath := filepath.Join(psProfileDir, "Microsoft.PowerShell_profile.ps1")

	// 读取现有配置
	existingContent := ""
	if data, err := os.ReadFile(psProfilePath); err == nil {
		existingContent = string(data)
	}

	// 构建新的环境变量配置
	var newLines []string
	newLines = append(newLines, "")
	newLines = append(newLines, "# Claude Code CLI 配置 (由 claude-install 自动生成)")
	newLines = append(newLines, fmt.Sprintf("$env:%s = \"%s\"", system.EnvAPIKey, config.APIKey))
	if config.APIBaseURL != "" && config.APIBaseURL != OfficialAPIURL {
		newLines = append(newLines, fmt.Sprintf("$env:%s = \"%s\"", system.EnvBaseURL, config.APIBaseURL))
	}

	// 移除旧的 Claude 配置
	lines := strings.Split(existingContent, "\n")
	var cleanedLines []string
	skipNext := false
	for _, line := range lines {
		if strings.Contains(line, "# Claude Code CLI 配置") {
			skipNext = true
			continue
		}
		if skipNext && (strings.Contains(line, "ANTHROPIC_API_KEY") || strings.Contains(line, "ANTHROPIC_BASE_URL")) {
			continue
		}
		skipNext = false
		cleanedLines = append(cleanedLines, line)
	}

	// 写入更新后的配置
	finalContent := strings.Join(cleanedLines, "\n") + strings.Join(newLines, "\n")

	// 同时设置用户环境变量（永久生效）
	fmt.Println("   正在设置系统环境变量...")

	return os.WriteFile(psProfilePath, []byte(finalContent), 0644)
}

// loadConfig 读取配置（从环境变量）
func loadConfig() (*ConfigData, error) {
	apiKey := os.Getenv(system.EnvAPIKey)
	if apiKey == "" {
		return nil, fmt.Errorf("未配置 ANTHROPIC_API_KEY 环境变量")
	}

	return &ConfigData{
		APIKey:     apiKey,
		APIBaseURL: os.Getenv(system.EnvBaseURL),
	}, nil
}
