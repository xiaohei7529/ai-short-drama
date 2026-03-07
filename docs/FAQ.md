# AI Short Drama - 常见问题解答

**版本**: v1.0.0  
**最后更新**: 2026-03-07

---

## 🔧 安装部署

### Q1: Windows 系统如何安装？
**A**: 
1. 下载 `ai-short-drama-v1.0.0-windows.zip`
2. 解压到任意目录（建议 D:\AI-Short-Drama）
3. 双击运行 `start.bat`
4. 浏览器访问 http://localhost:5678

### Q2: 提示缺少 Go 环境怎么办？
**A**: 
- 预编译包已包含所有依赖，无需单独安装 Go
- 如需源码编译，访问 https://golang.org/dl/ 下载安装

### Q3: FFmpeg 必须安装吗？
**A**: 
- 基础功能不需要 FFmpeg
- 视频生成/编辑功能需要 FFmpeg
- Windows 用户可下载：https://ffmpeg.org/download.html

### Q4: 端口 5678 被占用怎么办？
**A**: 
编辑 `config.yaml`:
```yaml
server:
  port: 5679  # 修改为其他端口
```

---

## 🤖 AI 配置

### Q5: 支持哪些 AI 提供商？
**A**: 
目前支持 5 家：
- ✅ OpenAI (GPT-4/DALL-E 3)
- ✅ 豆包 (火山引擎)
- ✅ 通义千问 (阿里云)
- ✅ 文心一言 (百度)
- ✅ Gemini (Google)

### Q6: 如何配置 API Key？
**A**: 
编辑 `config.yaml`:
```yaml
ai:
  openai:
    api_key: "sk-your-api-key"
  dashscope:
    api_key: "your-dashscope-key"
```

### Q7: 免费额度是多少？
**A**: 
| 提供商 | 免费额度 | 申请地址 |
|--------|----------|----------|
| OpenAI | $5 (新用户) | platform.openai.com |
| 豆包 | 有免费额度 | www.volcengine.com |
| 通义千问 | 有免费额度 | dashscope.console.aliyun.com |
| 文心一言 | 有免费额度 | cloud.baidu.com |

### Q8: AI 生成失败怎么办？
**A**: 
1. 检查 API Key 是否正确
2. 确认账户余额充足
3. 查看日志 `logs/app.log`
4. 切换其他 AI 提供商

---

## 🎬 使用操作

### Q9: 如何创建第一个短剧？
**A**: 
1. 访问 http://localhost:5678
2. 点击「剧本创作」
3. 输入故事创意（如："霸道总裁爱上我"）
4. 选择集数（1-10 集）
5. 点击「生成剧本」

### Q10: 生成一个短剧需要多久？
**A**: 
| 步骤 | 耗时 |
|------|------|
| 剧本生成 | 30-60 秒 |
| 角色设计 | 1-2 分钟 |
| 分镜生成 | 2-3 分钟 |
| 视频合成 | 3-5 分钟/集 |
| **总计** | **10-15 分钟** |

### Q11: 如何批量生成？
**A**: 
1. 准备剧本列表（JSON 格式）
2. 进入「批量生成」页面
3. 上传 JSON 文件
4. 设置并发数（建议 3）
5. 点击「开始生成」

### Q12: 视频质量如何调整？
**A**: 
编辑 `config.yaml`:
```yaml
ffmpeg:
  crf: 18  # 0-51，越小质量越高（推荐 18-23）
  output_format: "mp4"
  video_codec: "libx264"
```

---

## 🎨 角色与剧本

### Q13: 角色形象不一致怎么办？
**A**: 
1. 使用相同的角色描述词
2. 固定随机种子（seed）
3. 上传参考图
4. 使用角色一致性功能（v2.0+）

### Q14: 有剧本模板吗？
**A**: 
有的！查看 `examples/scripts/README.md`:
- 霸道总裁爱上我（5 集）
- 逆袭人生（8 集）
- 甜宠日常（10 集）

### Q15: 可以自定义剧本吗？
**A**: 
可以！两种方式：
1. 上传自己的剧本（TXT/JSON）
2. 使用 AI 生成后手动编辑

---

## 💾 数据管理

### Q16: 数据存储在哪里？
**A**: 
```
ai-short-drama/
├── data/
│   ├── drama_generator.db  # 数据库
│   └── storage/            # 素材文件
└── logs/
    └── app.log             # 日志文件
```

### Q17: 如何备份数据？
**A**: 
备份 `data/` 目录即可：
```bash
# Windows
xcopy /E /I data D:\Backup\ai-drama-data

# Linux/Mac
cp -r data /backup/ai-drama-data
```

### Q18: 如何迁移到新电脑？
**A**: 
1. 备份 `data/` 目录
2. 新电脑安装程序
3. 恢复 `data/` 目录
4. 编辑 `config.yaml` 配置

---

## 🐛 故障排查

### Q19: 程序无法启动？
**A**: 
1. 检查端口是否被占用
2. 查看日志 `logs/app.log`
3. 确认配置文件格式正确
4. 重启程序

### Q20: 视频生成失败？
**A**: 
1. 检查 FFmpeg 是否安装
2. 确认 AI API Key 有效
3. 查看错误日志
4. 降低并发数

### Q21: 内存不足怎么办？
**A**: 
编辑 `config.yaml`:
```yaml
queue:
  max_concurrent: 1  # 降低并发数
```

### Q22: 生成速度慢？
**A**: 
1. 使用 GPU 加速（如有）
2. 降低视频分辨率
3. 减少并发任务数
4. 升级硬件配置

---

## 🔒 安全隐私

### Q23: 数据会上传到云端吗？
**A**: 
- 所有数据本地存储
- 仅 AI 调用时发送请求到 AI 提供商
- 不会上传用户隐私数据

### Q24: API Key 安全吗？
**A**: 
- API Key 存储在本地配置文件
- 不会上传到 GitHub 或第三方
- 建议不要分享配置文件

---

## 📞 技术支持

### Q25: 如何反馈问题？
**A**: 
1. GitHub Issues: https://github.com/xiaohei7529/ai-short-drama/issues
2. 查看日志 `logs/app.log`
3. 提供详细错误信息

### Q26: 如何贡献代码？
**A**: 
1. Fork 项目
2. 创建分支 (`feature/amazing-feature`)
3. 提交更改
4. 推送并创建 Pull Request

### Q27: 有交流群吗？
**A**: 
- GitHub Discussions: https://github.com/xiaohei7529/ai-short-drama/discussions
- 邮件：xiaohei7529@gmail.com

---

## 💰 商业使用

### Q28: 可以商用吗？
**A**: 
- 可以！本项目采用 MIT 协议开源
- 商用无需授权
- 请保留版权声明

### Q29: 有企业版吗？
**A**: 
- 计划中（v3.0+）
- 企业版功能：
  - 多用户管理
  - 团队协作
  - 私有化部署
  - 技术支持

---

## 📚 更多资源

- **快速开始**: [QUICKSTART.md](QUICKSTART.md)
- **项目规划**: [PROJECT_ROADMAP.md](PROJECT_ROADMAP.md)
- **开发日志**: [DEV_LOG.md](DEV_LOG.md)
- **剧本模板**: [examples/scripts/README.md](examples/scripts/README.md)

---

*最后更新：2026-03-07*  
*版本：v1.0.0*

**🎬 AI Short Drama - 让每个人都能轻松创作 AI 短剧！**
