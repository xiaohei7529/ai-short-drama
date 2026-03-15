#!/bin/bash
# 每日开发报告生成脚本
# 执行时间：每日 21:00

WORKSPACE="/home/feng/openclaw-projects/ai-short-drama"
REPORTS_DIR="$WORKSPACE/reports"
DATE=$(date +%Y-%m-%d)
REPORT_FILE="$REPORTS_DIR/daily-$DATE.md"

mkdir -p "$REPORTS_DIR"

# 获取今日 Git 提交
COMMITS=$(cd "$WORKSPACE" && git log --since="00:00" --until="23:59" --oneline 2>/dev/null | wc -l)
COMMIT_LIST=$(cd "$WORKSPACE" && git log --since="00:00" --until="23:59" --oneline 2>/dev/null)

# 获取文件变更
FILES_CHANGED=$(cd "$WORKSPACE" && git diff --shortstat HEAD~1 2>/dev/null || echo "无变更")

# 生成报告
cat > "$REPORT_FILE" << REPORT
# 📅 AI Short Drama - 每日开发报告

**日期**: $DATE  
**项目**: ai-short-drama (Web 应用)  
**版本**: v1.0.0

---

## 📊 今日概览

- **Git 提交**: $COMMITS 次
- **文件变更**: $FILES_CHANGED
- **开发时长**: - 小时
- **完成功能**: - 个

---

## ✅ 完成内容

$COMMIT_LIST

---

## 📝 文档更新

- [ ] 开发日志
- [ ] API 文档
- [ ] 用户文档

---

## 🐛 问题解决

| 问题 | 状态 | 解决方案 |
|------|------|----------|
| - | - | - |

---

## 📈 明日计划

- [ ] -
- [ ] -
- [ ] -

---

*报告生成时间：$(date "+%Y-%m-%d %H:%M")*
REPORT

echo "✅ 日报生成完成：$REPORT_FILE"
