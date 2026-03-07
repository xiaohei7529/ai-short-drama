# AI Short Drama - Windows 打包构建脚本
# 用于在 Windows 系统上打包可执行文件
# 使用方法：powershell -ExecutionPolicy Bypass -File build-windows.ps1

Write-Host "🎬 AI Short Drama - Windows 打包构建" -ForegroundColor Cyan
Write-Host "=====================================" -ForegroundColor Cyan

# 配置
$PROJECT_NAME = "ai-short-drama"
$VERSION = "1.0.0"
$BUILD_DIR = "build-windows"
$OUTPUT_DIR = "$BUILD_DIR/$PROJECT_NAME-v$VERSION-windows"

# 检查 Go 环境
Write-Host "`n📋 检查环境..." -ForegroundColor Yellow
$goVersion = go version 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 错误：未检测到 Go 环境，请先安装 Go 1.23+" -ForegroundColor Red
    Write-Host "💡 下载地址：https://golang.org/dl/" -ForegroundColor Yellow
    exit 1
}
Write-Host "✓ Go 环境：$goVersion" -ForegroundColor Green

# 检查 Node.js 环境
$nodeVersion = node --version 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 错误：未检测到 Node.js 环境，请先安装 Node.js 18+" -ForegroundColor Red
    Write-Host "💡 下载地址：https://nodejs.org/" -ForegroundColor Yellow
    exit 1
}
Write-Host "✓ Node.js 环境：$nodeVersion" -ForegroundColor Green

# 检查 FFmpeg
$ffmpegVersion = ffmpeg -version 2>&1 | Select-Object -First 1
if ($LASTEXITCODE -ne 0) {
    Write-Host "⚠️  警告：未检测到 FFmpeg，视频功能可能不可用" -ForegroundColor Yellow
    Write-Host "💡 下载地址：https://ffmpeg.org/download.html" -ForegroundColor Yellow
} else {
    Write-Host "✓ FFmpeg: $ffmpegVersion" -ForegroundColor Green
}

# 清理旧构建
Write-Host "`n🧹 清理旧构建..." -ForegroundColor Yellow
if (Test-Path $BUILD_DIR) {
    Remove-Item -Recurse -Force $BUILD_DIR
    Write-Host "✓ 已清理 $BUILD_DIR" -ForegroundColor Green
}

# 创建输出目录
New-Item -ItemType Directory -Path $OUTPUT_DIR -Force | Out-Null
Write-Host "✓ 创建构建目录：$OUTPUT_DIR" -ForegroundColor Green

# 构建前端
Write-Host "`n🎨 构建前端..." -ForegroundColor Yellow
Set-Location web
npm install 2>&1 | Out-Null
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 前端依赖安装失败" -ForegroundColor Red
    exit 1
}

npm run build 2>&1 | Out-Null
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 前端构建失败" -ForegroundColor Red
    exit 1
}
Write-Host "✓ 前端构建完成" -ForegroundColor Green
Set-Location ..

# 构建后端
Write-Host "`n⚙️  构建后端..." -ForegroundColor Yellow
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o "$OUTPUT_DIR/$PROJECT_NAME.exe" -ldflags="-s -w" main.go 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 后端构建失败" -ForegroundColor Red
    exit 1
}
Write-Host "✓ 后端构建完成：$PROJECT_NAME.exe" -ForegroundColor Green

# 复制配置文件
Write-Host "`n📄 复制配置文件..." -ForegroundColor Yellow
Copy-Item -Path "configs/config.example.yaml" -Destination "$OUTPUT_DIR/config.yaml" -Force
Write-Host "✓ 配置文件已复制" -ForegroundColor Green

# 复制前端构建产物
Write-Host "`n📦 复制前端资源..." -ForegroundColor Yellow
if (Test-Path "web/dist") {
    Copy-Item -Path "web/dist" -Destination "$OUTPUT_DIR/web" -Recurse -Force
    Write-Host "✓ 前端资源已复制" -ForegroundColor Green
}

# 创建启动脚本
Write-Host "`n🚀 创建启动脚本..." -ForegroundColor Yellow
$startScript = @"
@echo off
chcp 65001 >nul
echo ========================================
echo   AI Short Drama - Windows 启动脚本
echo ========================================
echo.
echo 正在启动服务...
echo.
echo API 地址：http://localhost:5678
echo 前端地址：http://localhost:5678
echo.
echo 按 Ctrl+C 停止服务
echo.
%~dp0ai-short-drama.exe
"@
$startScript | Out-File -FilePath "$OUTPUT_DIR/start.bat" -Encoding UTF8
Write-Host "✓ 启动脚本已创建：start.bat" -ForegroundColor Green

# 创建 README
Write-Host "`n📖 创建使用说明..." -ForegroundColor Yellow
$readme = @"
# AI Short Drama - Windows 版

## 版本信息
- 版本号：v$VERSION
- 构建日期：$(Get-Date -Format "yyyy-MM-dd HH:mm:ss")
- 平台：Windows x64

## 环境要求

### 必需
- Windows 10/11 x64
- 内存：4GB+ (推荐 8GB+)
- 磁盘空间：1GB+

### 可选（用于视频处理）
- FFmpeg 4.0+ (已包含基础功能，完整功能建议单独安装)

## 快速开始

### 1. 配置
编辑 \`config.yaml\` 文件，配置 AI 服务提供商 API Key：

\`\`\`yaml
ai:
  openai:
    api_key: "your-openai-api-key"
  doubao:
    api_key: "your-doubao-api-key"
\`\`\`

### 2. 启动
双击运行 \`start.bat\` 或在命令行执行：
\`\`\`
ai-short-drama.exe
\`\`\`

### 3. 访问
打开浏览器访问：http://localhost:5678

## 功能特性

✅ AI 剧本生成 - 一句话生成完整短剧剧本
✅ 角色设计 - AI 自动生成角色形象
✅ 分镜生成 - 自动生成分镜脚本和图像
✅ 视频合成 - 图像转视频，自动剪辑
✅ 素材管理 - 统一的素材库管理
✅ 任务跟踪 - 实时查看生成进度

## 目录结构

\`\`\`
ai-short-drama/
├── ai-short-drama.exe    # 主程序
├── config.yaml           # 配置文件
├── start.bat            # 启动脚本
├── data/                # 数据目录（自动生成）
│   ├── drama_generator.db  # SQLite 数据库
│   └── storage/         # 素材存储
└── web/                 # 前端资源
\`\`\`

## 常见问题

### Q: 程序无法启动？
A: 确保端口 5678 未被占用，检查防火墙设置。

### Q: 视频生成失败？
A: 检查是否安装 FFmpeg，并添加到系统环境变量。

### Q: API 调用失败？
A: 检查 config.yaml 中的 API Key 是否正确，确保网络畅通。

## 技术支持

- GitHub: https://github.com/xiaohei7529/ai-short-drama
- Issues: https://github.com/xiaohei7529/ai-short-drama/issues

## 更新日志

### v1.0.0 (2026-03-07)
- 首次发布
- 基于 huobao-drama 优化
- 增加 Windows 一键打包
- 优化中文界面

---
Made with ❤️ by AI Short Drama Team
"@
$readme | Out-File -FilePath "$OUTPUT_DIR/README.md" -Encoding UTF8
Write-Host "✓ 使用说明已创建：README.md" -ForegroundColor Green

# 创建 .gitignore
Write-Host "`n📝 创建 .gitignore..." -ForegroundColor Yellow
$gitignore = @"
# 数据文件
data/*.db
data/storage/*
!data/.gitkeep

# 日志
logs/*.log

# 临时文件
*.tmp
*.temp

# 系统文件
.DS_Store
Thumbs.db
"@
$gitignore | Out-File -FilePath "$OUTPUT_DIR/.gitignore" -Encoding UTF8
Write-Host "✓ .gitignore 已创建" -ForegroundColor Green

# 创建 data 目录
New-Item -ItemType Directory -Path "$OUTPUT_DIR/data" -Force | Out-Null
New-Item -ItemType File -Path "$OUTPUT_DIR/data/.gitkeep" -Force | Out-Null

# 打包 ZIP
Write-Host "`n📦 打包 ZIP..." -ForegroundColor Yellow
$zipFile = "$BUILD_DIR/$PROJECT_NAME-v$VERSION-windows.zip"
Compress-Archive -Path "$OUTPUT_DIR/*" -DestinationPath $zipFile -Force
Write-Host "✓ 打包完成：$zipFile" -ForegroundColor Green

# 显示构建信息
Write-Host "`n=====================================" -ForegroundColor Cyan
Write-Host "🎉 构建完成！" -ForegroundColor Green
Write-Host "=====================================" -ForegroundColor Cyan
Write-Host "📦 安装包：$zipFile" -ForegroundColor Cyan
Write-Host "📁 解压目录：$OUTPUT_DIR" -ForegroundColor Cyan
Write-Host "🚀 启动方式：运行 start.bat 或 ai-short-drama.exe" -ForegroundColor Cyan
Write-Host "=====================================" -ForegroundColor Cyan

# 返回原目录
Set-Location ..
