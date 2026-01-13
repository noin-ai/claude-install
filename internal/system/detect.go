package system

import (
	"net/http"
	"runtime"
	"time"
)

// SystemInfo 系统信息
type SystemInfo struct {
	OS          string // windows, darwin, linux
	Arch        string // amd64, arm64
	NetworkType string // domestic, international
}

// Detect 检测系统信息
func Detect() *SystemInfo {
	info := &SystemInfo{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}

	// 检测网络环境（判断是否需要加速）
	info.NetworkType = detectNetworkType()

	return info
}

// detectNetworkType 检测网络类型
func detectNetworkType() string {
	// 通过测试访问 GitHub 和国内镜像源的速度判断网络环境
	// 如果国内镜像源更快或 GitHub 无法访问，则判定为国内网络

	githubLatency := testURL("https://github.com")
	chinaLatency := testURL("https://registry.npmmirror.com")

	// 如果 GitHub 访问失败或延迟超过 3 秒，判定为需要国内加速
	if githubLatency < 0 || githubLatency > 3000 {
		return "domestic"
	}

	// 如果国内镜像明显更快（快 2 倍以上），判定为国内网络
	if chinaLatency > 0 && githubLatency > chinaLatency*2 {
		return "domestic"
	}

	return "international"
}

// testURL 测试 URL 访问延迟（毫秒），失败返回 -1
func testURL(url string) int64 {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	start := time.Now()
	resp, err := client.Head(url)
	if err != nil {
		return -1
	}
	defer resp.Body.Close()

	latency := time.Since(start).Milliseconds()
	return latency
}

// IsChina 是否为国内网络环境
func (s *SystemInfo) IsChina() bool {
	return s.NetworkType == "domestic"
}

// IsWindows 是否为 Windows 系统
func (s *SystemInfo) IsWindows() bool {
	return s.OS == "windows"
}

// IsMacOS 是否为 macOS 系统
func (s *SystemInfo) IsMacOS() bool {
	return s.OS == "darwin"
}

// IsLinux 是否为 Linux 系统
func (s *SystemInfo) IsLinux() bool {
	return s.OS == "linux"
}
