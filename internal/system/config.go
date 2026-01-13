package system

import (
	"os"
	"path/filepath"
)

// Claude Code CLI 使用的环境变量
const (
	EnvAPIKey  = "ANTHROPIC_API_KEY"
	EnvBaseURL = "ANTHROPIC_BASE_URL"
)

// IsConfigured 检查是否已配置（检查环境变量）
func IsConfigured() bool {
	apiKey := os.Getenv(EnvAPIKey)
	return apiKey != ""
}

// GetConfigPath 获取 Claude Code CLI 配置目录
// 官方路径: ~/.claude/settings.json (所有平台)
func GetConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".claude", "settings.json")
}

// GetClaudeDir 获取 Claude 配置目录
func GetClaudeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".claude")
}

// GetConfigValue 读取配置值（从环境变量）
func GetConfigValue(key string) string {
	switch key {
	case "apiBase", "apiBaseUrl":
		return os.Getenv(EnvBaseURL)
	case "apiKey":
		apiKey := os.Getenv(EnvAPIKey)
		if apiKey != "" {
			// 隐藏 API Key，只显示前几位
			if len(apiKey) > 8 {
				return apiKey[:8] + "..."
			}
			return "***"
		}
		return ""
	default:
		return ""
	}
}
