# AI Short Drama - API 接口文档

**版本**: v1.0.0  
**最后更新**: 2026-03-07  
**基础 URL**: `http://localhost:5678/api/v1`

---

## 📋 接口列表

### 健康检查

#### GET /health
检查服务健康状态

**响应**:
```json
{
  "status": "ok",
  "timestamp": "2026-03-07T12:00:00Z",
  "version": "1.0.0"
}
```

---

### AI 配置

#### GET /ai/configs
获取所有 AI 提供商配置

**响应**:
```json
{
  "code": 0,
  "data": {
    "providers": [
      {
        "name": "openai",
        "type": "text,image",
        "enabled": true
      },
      {
        "name": "doubao",
        "type": "text,image,video",
        "enabled": true
      }
    ]
  }
}
```

#### POST /ai/config/test
测试 AI 提供商连接

**请求**:
```json
{
  "provider": "openai",
  "api_key": "sk-xxx"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "连接成功"
}
```

---

### 剧本生成

#### POST /script/generate
生成短剧剧本

**请求**:
```json
{
  "idea": "霸道总裁爱上我",
  "episodes": 5,
  "genre": "都市言情"
}
```

**响应**:
```json
{
  "code": 0,
  "data": {
    "script_id": 1,
    "status": "generating",
    "title": "霸道总裁爱上我",
    "episodes": 5,
    "estimated_time": 60
  }
}
```

#### GET /script/:id
获取剧本详情

**响应**:
```json
{
  "code": 0,
  "data": {
    "id": 1,
    "title": "霸道总裁爱上我",
    "episodes": [
      {
        "episode": 1,
        "title": "意外相遇",
        "content": "咖啡厅，苏浅浅送外卖..."
      }
    ]
  }
}
```

---

### 角色设计

#### POST /character/generate
生成角色形象

**请求**:
```json
{
  "description": "25 岁女性，长发，职业装",
  "count": 2,
  "style": "realistic"
}
```

**响应**:
```json
{
  "code": 0,
  "data": {
    "task_id": "char_123",
    "images": [
      "http://localhost:5678/static/char_1.png",
      "http://localhost:5678/static/char_2.png"
    ]
  }
}
```

#### POST /character/upload
上传角色图像

**请求**: `multipart/form-data`
- `image`: 图片文件
- `name`: 角色名称

**响应**:
```json
{
  "code": 0,
  "data": {
    "character_id": 1,
    "url": "http://localhost:5678/static/char_1.png"
  }
}
```

---

### 分镜生成

#### POST /storyboard/generate
生成分镜脚本

**请求**:
```json
{
  "script_id": 1,
  "episode": 1
}
```

**响应**:
```json
{
  "code": 0,
  "data": {
    "storyboard_id": 1,
    "scenes": [
      {
        "scene": 1,
        "location": "咖啡厅",
        "description": "苏浅浅送外卖...",
        "image": "http://localhost:5678/static/scene_1.png"
      }
    ]
  }
}
```

#### PUT /storyboard/:id
更新分镜

**请求**:
```json
{
  "scene_id": 1,
  "description": "修改后的描述",
  "image_prompt": "新的图像提示词"
}
```

---

### 视频生成

#### POST /video/generate
生成视频

**请求**:
```json
{
  "storyboard_id": 1,
  "resolution": "1080p",
  "fps": 30,
  "with_audio": true,
  "with_subtitle": true
}
```

**响应**:
```json
{
  "code": 0,
  "data": {
    "task_id": "video_456",
    "status": "processing",
    "estimated_time": 300
  }
}
```

#### GET /video/:id
获取视频状态

**响应**:
```json
{
  "code": 0,
  "data": {
    "id": "video_456",
    "status": "completed",
    "url": "http://localhost:5678/static/video_456.mp4",
    "duration": 180,
    "size": 52428800
  }
}
```

---

### 任务管理

#### GET /tasks
获取任务列表

**查询参数**:
- `status`: pending/processing/completed/failed
- `page`: 页码
- `size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "data": {
    "total": 10,
    "tasks": [
      {
        "id": "task_123",
        "type": "script",
        "status": "completed",
        "created_at": "2026-03-07T12:00:00Z"
      }
    ]
  }
}
```

#### GET /tasks/:id
获取任务详情

**响应**:
```json
{
  "code": 0,
  "data": {
    "id": "task_123",
    "type": "video",
    "status": "completed",
    "progress": 100,
    "result": {
      "url": "http://localhost:5678/static/video.mp4"
    }
  }
}
```

---

### 素材管理

#### GET /assets
获取素材列表

**查询参数**:
- `type`: character/scene/video
- `page`: 页码
- `size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "data": {
    "total": 50,
    "assets": [
      {
        "id": 1,
        "type": "character",
        "url": "http://localhost:5678/static/char_1.png",
        "created_at": "2026-03-07T12:00:00Z"
      }
    ]
  }
}
```

#### DELETE /assets/:id
删除素材

**响应**:
```json
{
  "code": 0,
  "message": "删除成功"
}
```

---

## ❌ 错误码

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 1001 | 参数错误 |
| 1002 | 认证失败 |
| 2001 | AI 服务错误 |
| 2002 | 视频生成失败 |
| 3001 | 文件上传失败 |
| 4001 | 任务不存在 |
| 5000 | 服务器内部错误 |

---

## 🔐 认证

当前版本无需认证，后续版本将支持 API Key 认证。

---

## 📊 限流

- 默认：60 请求/分钟
- 视频生成：3 并发
- 图像生成：5 并发

---

*最后更新：2026-03-07*  
*版本：v1.0.0*

**🎬 AI Short Drama - 让每个人都能轻松创作 AI 短剧！**
