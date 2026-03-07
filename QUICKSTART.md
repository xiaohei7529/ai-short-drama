# 🎬 AI Short Drama - 快速开始指南

**版本**: v1.0.0  
**最后更新**: 2026-03-07

---

## 🚀 5 分钟快速开始

### 方式一：下载预编译包 (推荐)

#### Windows 用户
1. 下载 [ai-short-drama-v1.0.0-windows.zip](https://github.com/xiaohei7529/ai-short-drama/releases)
2. 解压到任意目录
3. 双击运行 `start.bat`
4. 浏览器访问 http://localhost:5678

#### Linux/Mac 用户
```bash
# 下载
wget https://github.com/xiaohei7529/ai-short-drama/releases/download/v1.0.0/ai-short-drama-v1.0.0-linux.tar.gz

# 解压
tar -xzf ai-short-drama-v1.0.0-linux.tar.gz

# 进入目录
cd ai-short-drama-v1.0.0-linux

# 运行
./start.sh
```

---

### 方式二：源码编译

#### 环境要求
- **Go** 1.23+ ([下载](https://golang.org/dl/))
- **Node.js** 18+ ([下载](https://nodejs.org/))
- **FFmpeg** 4.0+ (可选，用于视频处理)
  - macOS: `brew install ffmpeg`
  - Ubuntu: `sudo apt install ffmpeg`
  - Windows: [下载](https://ffmpeg.org/download.html)

#### 编译步骤
```bash
# 1. 克隆项目
git clone https://github.com/xiaohei7529/ai-short-drama.git
cd ai-short-drama

# 2. 安装 Go 依赖
go mod download

# 3. 安装前端依赖
cd web
npm install
cd ..

# 4. 构建前端
cd web
npm run build
cd ..

# 5. 启动服务
go run main.go
```

访问 http://localhost:5678

---

## ⚙️ 配置指南

### 1. 复制配置文件
```bash
cp configs/config.example.yaml configs/config.yaml
```

### 2. 编辑配置文件
```yaml
ai:
  # 选择你的 AI 提供商
  default_text_provider: "openai"  # 或 doubao, dashscope, wenxin
  
  # OpenAI 配置
  openai:
    api_key: "sk-your-openai-api-key"
    text_model: "gpt-4o"
    image_model: "dall-e-3"
  
  # 豆包配置
  doubao:
    api_key: "your-doubao-api-key"
    api_secret: "your-doubao-api-secret"
  
  # 通义千问配置
  dashscope:
    api_key: "your-dashscope-api-key"
```

### 3. 获取 API Key

| 提供商 | 申请地址 | 免费额度 |
|--------|----------|----------|
| OpenAI | https://platform.openai.com | $5 (新用户) |
| 豆包 | https://www.volcengine.com/ | 有免费额度 |
| 通义千问 | https://dashscope.console.aliyun.com/ | 有免费额度 |
| 文心一言 | https://cloud.baidu.com/product/wenxinworkshop | 有免费额度 |

---

## 🎯 使用流程

### 第一步：创建剧本
1. 点击「剧本创作」
2. 输入故事创意（如："霸道总裁爱上我"）
3. 选择剧集数量（1-10 集）
4. 点击「生成剧本」

### 第二步：设计角色
1. 进入「角色设计」
2. 描述角色特征（如："25 岁女性，长发，职业装"）
3. 选择生成数量
4. 点击「生成角色」

### 第三步：生成分镜
1. 进入「分镜生成」
2. 选择已创建的剧本
3. 自动分析场景和镜头
4. 生成分镜图像

### 第四步：生成视频
1. 进入「视频生成」
2. 选择分镜序列
3. 选择转场效果
4. 点击「生成视频」

### 第五步：导出作品
1. 预览生成的短剧
2. 调整不满意的部分
3. 点击「导出」
4. 下载 MP4 文件

---

## 📖 详细文档

- [项目规划](PROJECT_ROADMAP.md) - 发展路线和功能计划
- [开发日志](DEV_LOG.md) - 实时开发进度追踪
- [API 文档](docs/API.md) - 完整 API 接口说明
- [部署指南](docs/DEPLOYMENT.md) - 生产环境部署

---

## ❓ 常见问题

### Q: 视频生成失败？
A: 
1. 检查是否安装 FFmpeg
2. 确认 API Key 有效
3. 查看日志文件 `logs/app.log`

### Q: 角色形象不一致？
A: 
1. 使用相同的角色描述
2. 固定随机种子 (seed)
3. 使用角色参考图

### Q: 生成速度慢？
A: 
1. 视频生成需要 2-5 分钟/集
2. 可调整并发数 `queue.max_concurrent`
3. 使用本地 GPU 加速

### Q: 如何批量生成？
A: 
1. 准备剧本列表 (JSON 格式)
2. 使用批量上传功能
3. 设置并发任务数

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

## 📞 联系方式

- **GitHub**: https://github.com/xiaohei7529/ai-short-drama
- **Issues**: https://github.com/xiaohei7529/ai-short-drama/issues
- **Email**: xiaohei7529@gmail.com

---

## 📄 开源协议

MIT License - 详见 [LICENSE](LICENSE)

---

**Made with ❤️ by AI Short Drama Team**
