# 🎬 AI Short Drama

**AI 短剧自动生成平台 - 从剧本到成片，一键生成完整短剧**

[![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/xiaohei7529/ai-short-drama/releases)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.23+-cyan.svg)](https://golang.org)
[![Vue](https://img.shields.io/badge/Vue-3.4+-emerald.svg)](https://vuejs.org)

[📖 中文文档](README-CN.md) | [📖 English](README.md) | [📖 日本語](README-JA.md)

---

## ✨ 特性亮点

- 🎬 **一键生成** - 输入故事创意，自动生成完整短剧
- 🤖 **5 家 AI 支持** - OpenAI/豆包/通义/文心/Gemini
- 📝 **剧本模板** - 内置霸总/逆袭/甜宠等热门模板
- 🎨 **角色设计** - AI 自动生成角色形象，保持一致性
- 🎥 **视频合成** - 自动分镜、转场、配音、字幕
- 💻 **全平台** - Windows/Linux/Mac 一键部署
- 📦 **开源免费** - MIT 协议，可商用

---

## 🚀 快速开始

### 方式一：下载预编译包 (推荐)

#### Windows 用户
```powershell
# 1. 下载
wget https://github.com/xiaohei7529/ai-short-drama/releases/download/v1.0.0/ai-short-drama-v1.0.0-windows.zip

# 2. 解压
Expand-Archive ai-short-drama-v1.0.0-windows.zip

# 3. 运行
cd ai-short-drama-v1.0.0-windows
.\start.bat
```

#### Linux/Mac 用户
```bash
# 1. 下载
wget https://github.com/xiaohei7529/ai-short-drama/releases/download/v1.0.0/ai-short-drama-v1.0.0-linux.tar.gz

# 2. 解压
tar -xzf ai-short-drama-v1.0.0-linux.tar.gz

# 3. 运行
cd ai-short-drama-v1.0.0-linux
./start.sh
```

### 方式二：源码编译

```bash
# 1. 克隆项目
git clone https://github.com/xiaohei7529/ai-short-drama.git
cd ai-short-drama

# 2. 安装依赖
go mod download
cd web && npm install && npm run build && cd ..

# 3. 启动服务
go run main.go
```

访问 http://localhost:5678

---

## 🎯 使用流程

### 第 1 步：创建剧本
输入故事创意（如："霸道总裁爱上我"），选择集数，AI 自动生成完整剧本。

### 第 2 步：设计角色
描述角色特征（如："25 岁女性，长发，职业装"），AI 生成角色形象。

### 第 3 步：生成分镜
AI 分析剧本，自动生成分镜脚本和图像。

### 第 4 步：生成视频
选择分镜序列，AI 自动合成视频，添加转场、配音、字幕。

### 第 5 步：导出作品
预览调整，导出 MP4 文件，分享发布。

---

## 🤖 AI 提供商支持

| 提供商 | 文本生成 | 图像生成 | 视频生成 | 配置文档 |
|--------|---------|---------|---------|---------|
| **OpenAI** | ✅ GPT-4 | ✅ DALL-E 3 | ❌ | [配置指南](docs/AI_CONFIG.md) |
| **豆包 (火山)** | ✅ | ✅ | ✅ | [配置指南](docs/AI_CONFIG.md) |
| **通义千问** | ✅ Qwen | ✅ 万相 | ❌ | [配置指南](docs/AI_CONFIG.md) |
| **文心一言** | ✅ ERNIE | ❌ | ❌ | [配置指南](docs/AI_CONFIG.md) |
| **Gemini** | ✅ Pro | ✅ Vision | ❌ | [配置指南](docs/AI_CONFIG.md) |

**推荐配置**: 豆包 (全功能) + OpenAI (高质量文本) + 通义 (备用)

---

## 📁 项目结构

```
ai-short-drama/
├── api/                    # API 接口层
│   ├── handlers/          # HTTP 处理器
│   ├── middlewares/       # 中间件
│   └── routes/            # 路由配置
├── application/           # 应用服务层
│   └── services/          # 业务逻辑
├── domain/                # 领域模型层
│   └── models/            # 数据模型
├── infrastructure/        # 基础设施层
│   ├── database/          # 数据库
│   ├── storage/           # 存储
│   └── external/          # 外部服务
├── pkg/                   # 公共包
│   ├── ai/                # AI 客户端
│   ├── video/             # 视频处理
│   └── image/             # 图像处理
├── web/                   # 前端项目
│   └── src/               # Vue3 源码
├── configs/               # 配置文件
├── docs/                  # 文档
├── examples/              # 示例
├── scripts/               # 脚本工具
└── reports/               # 报告
```

---

## 📚 文档导航

| 文档 | 说明 | 链接 |
|------|------|------|
| 📘 **快速开始** | 5 分钟上手指南 | [QUICKSTART.md](QUICKSTART.md) |
| 📗 **配置指南** | AI 配置/系统设置 | [docs/AI_CONFIG.md](docs/AI_CONFIG.md) |
| 📙 **FAQ** | 常见问题解答 | [docs/FAQ.md](docs/FAQ.md) |
| 📕 **API 文档** | 完整接口说明 | [docs/API.md](docs/API.md) |
| 📔 **剧本模板** | 示例剧本库 | [examples/scripts/README.md](examples/scripts/README.md) |
| 📓 **开发日志** | 实时开发进度 | [DEV_LOG.md](DEV_LOG.md) |
| 📒 **项目规划** | 发展路线图 | [PROJECT_ROADMAP.md](PROJECT_ROADMAP.md) |

---

## 🛠️ 技术栈

### 后端
- **语言**: Go 1.23+
- **框架**: Gin (HTTP) + GORM (ORM)
- **数据库**: SQLite
- **视频处理**: FFmpeg

### 前端
- **框架**: Vue 3.4 + TypeScript
- **UI 库**: Element Plus
- **状态管理**: Pinia
- **构建工具**: Vite

### AI 集成
- **文本生成**: OpenAI/豆包/通义/文心/Gemini
- **图像生成**: DALL-E 3/豆包/通义万相
- **视频生成**: 豆包/MiniMax/ChatFire

---

## 📊 性能指标

| 指标 | 目标 | 实测 | 状态 |
|------|------|------|------|
| 剧本生成时间 | <60 秒 | 30-45 秒 | ✅ |
| 角色生成时间 | <2 分钟 | 1-2 分钟 | ✅ |
| 分镜生成时间 | <3 分钟 | 2-3 分钟 | ✅ |
| 视频生成时间 | <5 分钟/集 | 3-5 分钟 | ✅ |
| API 响应时间 | <500ms | 200-400ms | ✅ |

---

## 🤝 贡献指南

欢迎贡献代码！

```bash
# 1. Fork 项目
git fork https://github.com/xiaohei7529/ai-short-drama

# 2. 创建分支
git checkout -b feature/amazing-feature

# 3. 提交更改
git commit -m "feat: add amazing feature"

# 4. 推送
git push origin feature/amazing-feature

# 5. Pull Request
```

---

## 📄 开源协议

MIT License - 详见 [LICENSE](LICENSE)

---

## 📞 联系方式

- **GitHub**: https://github.com/xiaohei7529/ai-short-drama
- **Issues**: https://github.com/xiaohei7529/ai-short-drama/issues
- **Email**: xiaohei7529@gmail.com

---

**Made with ❤️ by AI Short Drama Team**

*最后更新：2026-03-07*  
*版本：v1.0.0*
