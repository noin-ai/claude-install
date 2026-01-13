package system

import (
	"os/exec"
	"strings"
)

// CheckDependencies 检查所有依赖的安装状态
func CheckDependencies() map[string]bool {
	deps := map[string]bool{
		"Node.js": isCommandAvailable("node"),
		"npm":     isCommandAvailable("npm"),
		"Git":     isCommandAvailable("git"),
	}
	return deps
}

// isCommandAvailable 检查命令是否可用
func isCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// GetVersion 获取工具版本
func GetVersion(tool string) string {
	var cmd *exec.Cmd

	switch tool {
	case "Node.js":
		cmd = exec.Command("node", "--version")
	case "npm":
		cmd = exec.Command("npm", "--version")
	case "Git":
		cmd = exec.Command("git", "--version")
	default:
		return ""
	}

	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	version := strings.TrimSpace(string(output))
	// 清理输出（git 会输出 "git version x.x.x"）
	if strings.HasPrefix(version, "git version") {
		version = strings.TrimPrefix(version, "git version ")
		version = strings.TrimSpace(version)
	}

	return version
}

// IsClaudeCLIInstalled 检查 Claude CLI 是否已安装
func IsClaudeCLIInstalled() bool {
	return isCommandAvailable("claude")
}

// GetClaudeCLIVersion 获取 Claude CLI 版本
func GetClaudeCLIVersion() string {
	cmd := exec.Command("claude", "--version")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}
