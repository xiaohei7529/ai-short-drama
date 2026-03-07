# AI Short Drama - Linux/Mac 打包构建脚本
# 用于在 Linux/macOS 系统上打包可执行文件
# 使用方法：bash build.sh

#!/bin/bash

set -e

# 配置
PROJECT_NAME="ai-short-drama"
VERSION="1.0.0"
BUILD_DIR="build"
OUTPUT_DIR="$BUILD_DIR/$PROJECT_NAME-v$VERSION"

echo "🎬 AI Short Drama - 打包构建"
echo "====================================="

# 检查 Go 环境
echo -e "\n📋 检查环境..."
if ! command -v go &> /dev/null; then
    echo "❌ 错误：未检测到 Go 环境，请先安装 Go 1.23+"
    echo "💡 下载地址：https://golang.org/dl/"
    exit 1
fi
echo "✓ Go 环境：$(go version)"

# 检查 Node.js 环境
if ! command -v node &> /dev/null; then
    echo "❌ 错误：未检测到 Node.js 环境，请先安装 Node.js 18+"
    echo "💡 下载地址：https://nodejs.org/"
    exit 1
fi
echo "✓ Node.js 环境：$(node --version)"

# 检查 FFmpeg
if command -v ffmpeg &> /dev/null; then
    echo "✓ FFmpeg: $(ffmpeg -version | head -1)"
else
    echo "⚠️  警告：未检测到 FFmpeg，视频功能可能不可用"
    echo "💡 安装：brew install ffmpeg (macOS) 或 sudo apt install ffmpeg (Linux)"
fi

# 清理旧构建
echo -e "\n🧹 清理旧构建..."
rm -rf "$BUILD_DIR"
mkdir -p "$OUTPUT_DIR"
echo "✓ 创建构建目录：$OUTPUT_DIR"

# 构建前端
echo -e "\n🎨 构建前端..."
cd web
npm install
npm run build
cd ..
echo "✓ 前端构建完成"

# 构建后端（当前平台）
echo -e "\n⚙️  构建后端..."
go build -o "$OUTPUT_DIR/$PROJECT_NAME" -ldflags="-s -w" main.go
echo "✓ 后端构建完成：$PROJECT_NAME"

# 复制配置文件
echo -e "\n📄 复制配置文件..."
cp configs/config.example.yaml "$OUTPUT_DIR/config.yaml"
echo "✓ 配置文件已复制"

# 复制前端构建产物
echo -e "\n📦 复制前端资源..."
cp -r web/dist "$OUTPUT_DIR/web"
echo "✓ 前端资源已复制"

# 创建启动脚本
echo -e "\n🚀 创建启动脚本..."
cat > "$OUTPUT_DIR/start.sh" << 'EOF'
#!/bin/bash
echo "========================================"
echo "  AI Short Drama - 启动脚本"
echo "========================================"
echo ""
echo "正在启动服务..."
echo ""
echo "API 地址：http://localhost:5678"
echo "前端地址：http://localhost:5678"
echo ""
echo "按 Ctrl+C 停止服务"
echo ""
./ai-short-drama
EOF
chmod +x "$OUTPUT_DIR/start.sh"
echo "✓ 启动脚本已创建：start.sh"

# 创建 README
echo -e "\n📖 创建使用说明..."
cat > "$OUTPUT_DIR/README.md" << EOF
# AI Short Drama - $(uname -s) 版

## 版本信息
- 版本号：v$VERSION
- 构建日期：$(date "+%Y-%m-%d %H:%M:%S")
- 平台：$(uname -s) $(uname -m)

## 快速开始

### 1. 配置
编辑 \`config.yaml\` 文件，配置 AI 服务提供商 API Key

### 2. 启动
\`\`\`bash
./start.sh
\`\`\`

### 3. 访问
打开浏览器访问：http://localhost:5678

## 功能特性
✅ AI 剧本生成 | ✅ 角色设计 | ✅ 分镜生成
✅ 视频合成 | ✅ 素材管理 | ✅ 任务跟踪

## 技术支持
GitHub: https://github.com/xiaohei7529/ai-short-drama
EOF
echo "✓ 使用说明已创建"

# 创建 data 目录
mkdir -p "$OUTPUT_DIR/data"
touch "$OUTPUT_DIR/data/.gitkeep"

# 打包
echo -e "\n📦 打包..."
cd "$BUILD_DIR"
tar -czf "$PROJECT_NAME-v$VERSION-$(uname -s | tr '[:upper:]' '[:lower:]').tar.gz" "$PROJECT_NAME-v$VERSION"
cd ..
echo "✓ 打包完成：$BUILD_DIR/$PROJECT_NAME-v$VERSION-$(uname -s | tr '[:upper:]' '[:lower:]').tar.gz"

echo -e "\n====================================="
echo "🎉 构建完成！"
echo "====================================="
echo "📦 安装包：$BUILD_DIR/$PROJECT_NAME-v$VERSION-$(uname -s | tr '[:upper:]' '[:lower:]').tar.gz"
echo "📁 解压目录：$OUTPUT_DIR"
echo "🚀 启动方式：./start.sh"
echo "====================================="
