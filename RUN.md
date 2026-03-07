# 🚀 AI Short Drama - 运行指南

**版本**: v1.0.0  
**最后更新**: 2026-03-07

---

## 📋 环境要求

### 必需
- **Go**: 1.23+ (可选，仅开发需要)
- **Node.js**: 18+ (前端开发)
- **FFmpeg**: 4.0+ (视频处理)

### 可选
- **Air**: 热加载工具 (开发环境)

---

## 🔧 快速开始

### 方式一：直接运行（推荐）

```bash
cd /home/feng/openclaw-projects/ai-short-drama

# 1. 清理缓存
go clean -cache -modcache

# 2. 重新下载依赖
go mod download

# 3. 运行
go run main.go
```

**访问**: http://localhost:5678

---

### 方式二：编译后运行

```bash
cd /home/feng/openclaw-projects/ai-short-drama

# 1. 清理并编译
go clean -cache -modcache -i
go build -a -o ai-short-drama main.go

# 2. 运行
./ai-short-drama
```

---

### 方式三：使用 Air 热加载（开发）

```bash
# 安装 air
go install github.com/cosmtrek/air@latest

# 运行
air
```

---

## ⚠️ 常见问题

### 1. 阿里云 SDK 错误

**错误信息**:
```
cannot find module providing package github.com/aliyun/alibabacloud-dashscope-go-sdk/sdk
```

**解决方案**:
```bash
# 1. 确认已修复（当前版本已移除 SDK 依赖）
cat go.mod | grep alibabacloud

# 2. 清理缓存
go clean -cache -modcache -i

# 3. 重新下载依赖
go mod download

# 4. 运行
go run main.go
```

**原因**: 旧版本使用了阿里云 SDK，v1.0.0 已改为 HTTP 直连方式

---

### 2. 端口被占用

**错误信息**:
```
bind: address already in use
```

**解决方案**:
```bash
# 1. 查找占用端口的进程
lsof -i :5678

# 2. 杀死进程
kill -9 <PID>

# 3. 或者修改配置文件
编辑 configs/config.yaml
server:
  port: 5679  # 改为其他端口
```

---

### 3. 数据库连接失败

**错误信息**:
```
failed to connect to database
```

**解决方案**:
```bash
# 1. 检查数据库服务
# SQLite: 确认文件路径正确
# MySQL: 确认服务运行

# 2. 修改配置
编辑 configs/config.yaml
database:
  type: "sqlite"
  path: "./data/drama_generator.db"
```

---

### 4. 前端资源 404

**错误信息**: 访问页面显示空白或 404

**解决方案**:
```bash
# 1. 构建前端
cd web
npm install
npm run build
cd ..

# 2. 确认 web/dist 目录存在
ls -la web/dist/

# 3. 重启服务
go run main.go
```

---

## 📊 验证运行

### 1. 检查日志
```bash
# 运行后应该看到:
INFO Starting Drama Generator API Server...
INFO Database connected successfully
INFO Database tables migrated successfully
```

### 2. 测试 API
```bash
# 健康检查
curl http://localhost:5678/api/v1/health

# 应该返回:
{"status":"ok"}
```

### 3. 访问前端
```
http://localhost:5678
```

---

## 🔗 相关文档

- [快速开始](QUICKSTART.md)
- [配置指南](docs/AI_CONFIG.md)
- [API 文档](docs/API.md)
- [FAQ](docs/FAQ.md)
- [运行问题排查](docs/RUNTIME_FIX_GUIDE.md)

---

**Made with ❤️ by AI Short Drama Team**
