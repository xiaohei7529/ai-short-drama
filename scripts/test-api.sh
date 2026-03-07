#!/bin/bash
# AI Short Drama - API 接口测试脚本
# 用于测试核心 API 功能

BASE_URL="http://localhost:5678/api/v1"

echo "🧪 AI Short Drama - API 接口测试"
echo "====================================="
echo ""

# 测试健康检查
echo "📋 测试 1/6: 健康检查..."
curl -s "$BASE_URL/health" | jq . && echo "✅ 通过" || echo "❌ 失败"
echo ""

# 测试 AI 配置
echo "🤖 测试 2/6: AI 配置获取..."
curl -s "$BASE_URL/ai/configs" | jq . && echo "✅ 通过" || echo "❌ 失败"
echo ""

# 测试剧本生成
echo "📝 测试 3/6: 剧本生成..."
curl -s -X POST "$BASE_URL/script/generate" \
  -H "Content-Type: application/json" \
  -d '{"idea":"霸道总裁爱上我","episodes":3}' | jq . && echo "✅ 通过" || echo "❌ 失败"
echo ""

# 测试角色生成
echo "🎨 测试 4/6: 角色生成..."
curl -s -X POST "$BASE_URL/character/generate" \
  -H "Content-Type: application/json" \
  -d '{"description":"25 岁女性，长发，职业装","count":2}' | jq . && echo "✅ 通过" || echo "❌ 失败"
echo ""

# 测试分镜生成
echo "🎬 测试 5/6: 分镜生成..."
curl -s -X POST "$BASE_URL/storyboard/generate" \
  -H "Content-Type: application/json" \
  -d '{"script_id":1}' | jq . && echo "✅ 通过" || echo "❌ 失败"
echo ""

# 测试任务列表
echo "📊 测试 6/6: 任务列表..."
curl -s "$BASE_URL/tasks" | jq . && echo "✅ 通过" || echo "❌ 失败"
echo ""

echo "====================================="
echo "✅ 测试完成！"
