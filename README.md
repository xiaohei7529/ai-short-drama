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
- 💻 **全平台** - Windows/Linux/Mac一键部署
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

---

## 📊 开发状态

**当前版本**: v1.0.0  
**开发状态**: ✅ 已完成  
**最后更新**: 2026-03-15

### 本周进展 (2026-03-09 至 2026-03-15)

**核心成就**:
- ✅ v1.0.0 正式发布准备完成
- ✅ Release 说明文档编写
- ✅ 打包系统测试完成
- ✅ 文档体系完善 (20+ 份文档)

**Git 统计**:
- 本周提交：28 次
- 代码行数：3500+ 行
- 文档数量：20 份

**下周计划**:
- [ ] GitHub Release v1.0.0 正式发布
- [ ] 用户反馈收集
- [ ] Bug 修复和优化

---

## 📁 项目结构

```
ai-short-drama/
├── api/              # API 层
├── application/      # 应用层
├── cmd/              # 命令行工具
├── configs/          # 配置文件
├── data/             # 数据层
├── domain/           # 领域层
├── infrastructure/   # 基础设施
├── pkg/              # 公共包
├── web/              # Web 前端
├── docs/             # 文档
├── scripts/          # 脚本
└── examples/         # 示例
```

---

## 🛠️ 技术栈

### 后端
- **语言**: Go 1.23+
- **框架**: Gin
- **数据库**: SQLite/MySQL
- **ORM**: GORM

### 前端
- **框架**: Vue 3.4
- **UI 库**: Element Plus 2.5
- **构建工具**: Vite 5.0
- **状态管理**: Pinia 2.1

### AI 服务
- 通义千问 (DashScope)
- 文心一言 (ERNIE)
- 讯飞星火 (iFlytek)
- MiniMax
- Moonshot

---

## 📖 文档导航

| 文档 | 说明 |
|------|------|
| [快速开始](QUICKSTART.md) | 5 分钟快速上手 |
| [API 文档](docs/API.md) | 完整 API 参考 |
| [FAQ](docs/FAQ.md) | 常见问题解答 |
| [数据迁移](docs/DATA_MIGRATION.md) | 数据迁移指南 |
| [运行时修复](docs/RUNTIME_FIX_GUIDE.md) | 运行时问题修复 |

---

## 🤝 贡献指南

欢迎贡献代码、文档或建议！

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交变更 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

---

## 📄 开源协议

本项目采用 MIT 协议开源 - 查看 [LICENSE](LICENSE) 文件了解详情。

---

## 🙏 致谢

感谢以下开源项目：
- [Vue.js](https://vuejs.org/)
- [Element Plus](https://element-plus.org/)
- [DashScope](https://dashscope.aliyun.com/)

---

## 📞 联系方式

- **GitHub**: https://github.com/xiaohei7529/ai-short-drama
- **Issues**: https://github.com/xiaohei7529/ai-short-drama/issues
- **Email**: xiaohei7529@gmail.com

---

*AI Short Drama v1.0.0 - 让 AI 帮你创作精彩短剧！* 🎬

**最后更新**: 2026-03-15
