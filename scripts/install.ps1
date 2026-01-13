# Claude Code CLI 一键安装脚本 (Windows PowerShell)
# 使用方式: irm https://noin.ai/install.ps1 | iex

$ErrorActionPreference = "Stop"

# 配置
$BINARY_NAME = "claude-install"
$DOWNLOAD_BASE = "https://github.com/noin-ai/claude-install/releases/latest/download"
$INSTALL_DIR = "$env:USERPROFILE\.claude-install"

function Write-Step {
    param([string]$Message)
    Write-Host "`n$Message" -ForegroundColor Cyan
}

function Write-Success {
    param([string]$Message)
    Write-Host $Message -ForegroundColor Green
}

function Write-Error {
    param([string]$Message)
    Write-Host $Message -ForegroundColor Red
}

function Get-Architecture {
    $arch = [System.Environment]::GetEnvironmentVariable("PROCESSOR_ARCHITECTURE")
    switch ($arch) {
        "AMD64" { return "amd64" }
        "ARM64" { return "arm64" }
        default { return "amd64" }
    }
}

function Add-ToPath {
    param([string]$Path)

    $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($currentPath -notlike "*$Path*") {
        [Environment]::SetEnvironmentVariable("Path", "$currentPath;$Path", "User")
        $env:Path = "$env:Path;$Path"
        Write-Success "   已添加到 PATH 环境变量"
    }
}

# 开始安装
Write-Host @"

   ██████╗██╗      █████╗ ██╗   ██╗██████╗ ███████╗
  ██╔════╝██║     ██╔══██╗██║   ██║██╔══██╗██╔════╝
  ██║     ██║     ███████║██║   ██║██║  ██║█████╗
  ██║     ██║     ██╔══██║██║   ██║██║  ██║██╔══╝
  ╚██████╗███████╗██║  ██║╚██████╔╝██████╔╝███████╗
   ╚═════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝

  Claude Code CLI 一键安装工具
  By noin.ai

"@ -ForegroundColor Magenta

Write-Step "1. 检测系统架构..."
$arch = Get-Architecture
Write-Success "   Windows $arch"

Write-Step "2. 创建安装目录..."
if (!(Test-Path $INSTALL_DIR)) {
    New-Item -ItemType Directory -Path $INSTALL_DIR -Force | Out-Null
}
Write-Success "   $INSTALL_DIR"

Write-Step "3. 下载安装程序..."
$fileName = "$BINARY_NAME-windows-$arch.exe"
$downloadUrl = "$DOWNLOAD_BASE/$fileName"
$targetPath = "$INSTALL_DIR\$BINARY_NAME.exe"

try {
    # 尝试从 GitHub 下载
    Write-Host "   正在下载: $downloadUrl"
    Invoke-WebRequest -Uri $downloadUrl -OutFile $targetPath -UseBasicParsing
    Write-Success "   下载完成!"
} catch {
    Write-Error "   下载失败: $_"
    Write-Host "`n请手动下载并安装:" -ForegroundColor Yellow
    Write-Host "   $downloadUrl" -ForegroundColor White
    exit 1
}

Write-Step "4. 配置环境变量..."
Add-ToPath $INSTALL_DIR

Write-Step "5. 运行安装向导..."
Write-Host ""

# 运行安装程序
& $targetPath setup

Write-Host @"

============================================
  安装完成!

  重新打开 PowerShell 后，可以使用:
    claude-install check  - 检查环境状态
    claude-install config - 更新配置
    claude-install test   - 测试连接

  需要帮助? 访问: https://noin.ai/help
============================================

"@ -ForegroundColor Green
