# 🔧 AI Short Drama 运行错误排查指南

**问题**: `go run main.go` 出现阿里云 SDK 错误

---

## 📊 可能原因

### 1. 旧代码文件残留
**检查**: `pkg/ai/dashscope_client_old.go` 可能还在项目中

**解决**:
```bash
cd /home/feng/openclaw-projects/ai-short-drama
rm -f pkg/ai/dashscope_client_old.go
```

### 2. Go 模块缓存问题
**检查**: Go 可能缓存了旧的依赖

**解决**:
```bash
# 清理模块缓存
go clean -modcache

# 重新下载依赖
go mod download

# 整理依赖
go mod tidy
```

### 3. 编译缓存问题
**检查**: 旧的编译产物可能还在

**解决**:
```bash
# 清理所有编译缓存
go clean -cache -modcache -i

# 重新编译
go build -a -o ai-short-drama main.go
```

### 4. 导入路径错误
**检查**: 可能有其他文件还在导入旧 SDK

**解决**:
```bash
# 搜索所有 Go 文件中的 SDK 引用
grep -r "alibabacloud-dashscope" --include="*.go" .
```

---

## ✅ 正确运行步骤

### 方式一：直接运行
```bash
cd /home/feng/openclaw-projects/ai-short-drama

# 1. 清理缓存
go clean -cache -modcache

# 2. 重新下载依赖
go mod download

# 3. 运行
go run main.go
```

### 方式二：编译后运行
```bash
cd /home/feng/openclaw-projects/ai-short-drama

# 1. 清理并编译
go clean -cache -modcache -i
go build -a -o ai-short-drama main.go

# 2. 运行
./ai-short-drama
```

### 方式三：使用 Air 热加载（开发环境）
```bash
# 安装 air
go install github.com/cosmtrek/air@latest

# 运行
air
```

---

## 🔍 验证修复

### 1. 检查代码
```bash
# 确认没有旧 SDK 引用
grep -r "alibabacloud-dashscope" --include="*.go" . | grep -v "_old.go"
# 应该只输出注释行
```

### 2. 检查依赖
```bash
# 查看 go.mod
cat go.mod | grep "require (" -A 20
# 不应该有 alibabacloud-dashscope-go-sdk
```

### 3. 测试运行
```bash
# 运行并查看日志
go run main.go 2>&1 | head -50
```

---

## 📝 当前状态

### ✅ 已修复
- `pkg/ai/dashscope_client.go` - HTTP 直连版本
- `go.mod` - 已移除 SDK 依赖

### ⚠️ 待清理
- `pkg/ai/dashscope_client_old.go` - 备份文件（可删除）
- Go 编译缓存

---

## 🎯 快速修复命令

```bash
cd /home/feng/openclaw-projects/ai-short-drama

# 1. 删除备份文件
rm -f pkg/ai/dashscope_client_old.go

# 2. 清理缓存
go clean -cache -modcache -i

# 3. 重新下载依赖
go mod download

# 4. 运行
go run main.go
```

---

**修复者**: AI Assistant  
**时间**: 2026-03-07 22:00  
**状态**: ⏳ 待验证

**🎬 AI Short Drama - 运行问题排查指南！**
