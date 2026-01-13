package installer

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/noin-ai/claude-install/internal/system"
)

// InstallDependencies 安装缺失的依赖
func InstallDependencies(missing []string, sysInfo *system.SystemInfo) error {
	for _, dep := range missing {
		fmt.Printf("   📥 正在安装 %s...\n", dep)

		var err error
		switch dep {
		case "Node.js":
			err = installNodeJS(sysInfo)
		case "npm":
			// npm 通常随 Node.js 一起安装
			err = installNodeJS(sysInfo)
		case "Git":
			err = installGit(sysInfo)
		default:
			return fmt.Errorf("未知依赖: %s", dep)
		}

		if err != nil {
			return fmt.Errorf("安装 %s 失败: %w", dep, err)
		}

		fmt.Printf("   ✅ %s 安装完成\n", dep)
	}

	return nil
}

// installNodeJS 安装 Node.js
func installNodeJS(sysInfo *system.SystemInfo) error {
	fmt.Println()
	fmt.Println("⚠️  检测到系统缺少 Node.js")
	fmt.Println()

	switch runtime.GOOS {
	case "darwin":
		return installNodeJSMacOS(sysInfo)
	case "linux":
		return installNodeJSLinux(sysInfo)
	case "windows":
		return installNodeJSWindows(sysInfo)
	default:
		return fmt.Errorf("不支持的操作系统: %s", runtime.GOOS)
	}
}

// installNodeJSMacOS macOS 上安装 Node.js
func installNodeJSMacOS(sysInfo *system.SystemInfo) error {
	// 检查是否有 brew
	if _, err := exec.LookPath("brew"); err == nil {
		fmt.Println("   使用 Homebrew 安装 Node.js...")
		cmd := exec.Command("brew", "install", "node")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("Homebrew 安装失败: %w\n%s", err, string(output))
		}
		return nil
	}

	// 提供手动安装指导
	return fmt.Errorf(`请手动安装 Node.js:
方式 1: 使用 Homebrew (推荐)
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  brew install node

方式 2: 从官网下载
  访问: https://nodejs.org/
  下载并安装 macOS 安装包

安装完成后，重新运行 'claude-install setup'`)
}

// installNodeJSLinux Linux 上安装 Node.js
func installNodeJSLinux(sysInfo *system.SystemInfo) error {
	fmt.Println("   正在使用包管理器安装 Node.js...")

	// 尝试不同的包管理器
	managers := []struct {
		name    string
		check   string
		install []string
	}{
		{"apt", "apt-get", []string{"sudo", "apt-get", "update", "&&", "sudo", "apt-get", "install", "-y", "nodejs", "npm"}},
		{"yum", "yum", []string{"sudo", "yum", "install", "-y", "nodejs", "npm"}},
		{"dnf", "dnf", []string{"sudo", "dnf", "install", "-y", "nodejs", "npm"}},
	}

	for _, mgr := range managers {
		if _, err := exec.LookPath(mgr.check); err == nil {
			fmt.Printf("   使用 %s 安装...\n", mgr.name)
			cmd := exec.Command(mgr.install[0], mgr.install[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("安装失败: %w\n%s", err, string(output))
			}
			return nil
		}
	}

	return fmt.Errorf(`请手动安装 Node.js:
Ubuntu/Debian:
  sudo apt-get update
  sudo apt-get install -y nodejs npm

CentOS/RHEL:
  sudo yum install -y nodejs npm

或访问: https://nodejs.org/`)
}

// installNodeJSWindows Windows 上安装 Node.js
func installNodeJSWindows(sysInfo *system.SystemInfo) error {
	return fmt.Errorf(`请手动安装 Node.js:
1. 访问: https://nodejs.org/
2. 下载 Windows 安装包 (.msi)
3. 运行安装程序
4. 重启终端后，重新运行 'claude-install setup'

国内用户可以使用淘宝镜像:
https://npmmirror.com/mirrors/node/`)
}

// installGit 安装 Git
func installGit(sysInfo *system.SystemInfo) error {
	fmt.Println()
	fmt.Println("⚠️  检测到系统缺少 Git")
	fmt.Println()

	switch runtime.GOOS {
	case "darwin":
		fmt.Println("   macOS 系统会在首次使用时自动安装 Git")
		fmt.Println("   请在终端输入 'git' 并按提示安装")
		return nil
	case "linux":
		fmt.Println("   请使用包管理器安装 Git:")
		fmt.Println("   Ubuntu/Debian: sudo apt-get install git")
		fmt.Println("   CentOS/RHEL: sudo yum install git")
		return fmt.Errorf("请手动安装 Git 后重新运行")
	case "windows":
		return fmt.Errorf(`请手动安装 Git:
访问: https://git-scm.com/download/win
下载并安装 Git for Windows`)
	}

	return fmt.Errorf("不支持的操作系统")
}

// InstallClaudeCLI 安装 Claude Code CLI
func InstallClaudeCLI(sysInfo *system.SystemInfo) error {
	// 如果是国内网络，先配置 npm 镜像
	if sysInfo.IsChina() {
		fmt.Println("   检测到国内网络环境，配置 npm 镜像加速...")
		if err := configureNPMMirror(); err != nil {
			fmt.Printf("   ⚠️  配置镜像失败（将继续使用默认源）: %v\n", err)
		} else {
			fmt.Println("   ✅ npm 镜像配置完成")
		}
		fmt.Println()
	}

	// 安装 Claude Code CLI
	fmt.Println("   正在通过 npm 安装 Claude Code CLI...")
	fmt.Println("   这可能需要几分钟，请耐心等待...")
	fmt.Println()

	cmd := exec.Command("npm", "install", "-g", "@anthropic-ai/claude-code")
	output, err := cmd.CombinedOutput()

	if err != nil {
		// 显示详细错误信息
		outputStr := string(output)
		if strings.Contains(outputStr, "EACCES") || strings.Contains(outputStr, "permission denied") {
			return fmt.Errorf(`权限不足，请尝试：
macOS/Linux:
  sudo npm install -g @anthropic-ai/claude-code

Windows:
  以管理员身份运行终端，然后重新执行 'claude-install setup'

或者配置 npm 全局目录（推荐）:
  mkdir ~/.npm-global
  npm config set prefix '~/.npm-global'
  然后将 ~/.npm-global/bin 添加到 PATH`)
		}

		return fmt.Errorf("安装失败: %w\n输出: %s", err, outputStr)
	}

	return nil
}

// configureNPMMirror 配置 npm 国内镜像
func configureNPMMirror() error {
	// 配置淘宝镜像
	cmd := exec.Command("npm", "config", "set", "registry", "https://registry.npmmirror.com")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
