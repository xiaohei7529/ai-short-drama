package ai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	dashscope "github.com/aliyun/alibabacloud-dashscope-go-sdk/sdk"
)

// DashScopeClient 通义千问客户端
type DashScopeClient struct {
	apiKey    string
	textModel string
	imageModel string
	timeout   time.Duration
	client    *dashscope.Client
}

// NewDashScopeClient 创建通义千问客户端
func NewDashScopeClient(apiKey, textModel, imageModel string, timeoutSec int) *DashScopeClient {
	client := dashscope.NewClient(apiKey)
	return &DashScopeClient{
		apiKey:     apiKey,
		textModel:  textModel,
		imageModel: imageModel,
		timeout:    time.Duration(timeoutSec) * time.Second,
		client:     client,
	}
}

// GenerateText 生成文本
func (c *DashScopeClient) GenerateText(prompt string, systemPrompt string, options ...func(*ChatCompletionRequest)) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	messages := []dashscope.Message{
		{
			Role:    "system",
			Content: systemPrompt,
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	resp, err := c.client.Call(ctx, dashscope.Qwen, c.textModel, messages, nil)
	if err != nil {
		return "", fmt.Errorf("通义千问调用失败：%w", err)
	}

	if resp.Output == nil || resp.Output.Choices == nil || len(resp.Output.Choices) == 0 {
		return "", errors.New("通义千问返回空响应")
	}

	return resp.Output.Choices[0].Message.Content, nil
}

// GenerateImage 生成图像
func (c *DashScopeClient) GenerateImage(prompt string, size string, n int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	// 使用通义万相生成图像
	resp, err := c.client.CallImage(ctx, dashscope.Wanx, c.imageModel, prompt, nil)
	if err != nil {
		return nil, fmt.Errorf("通义万相调用失败：%w", err)
	}

	if resp.Output == nil || resp.Output.Results == nil {
		return nil, errors.New("通义万相返回空响应")
	}

	urls := make([]string, 0, n)
	for i := 0; i < n && i < len(resp.Output.Results); i++ {
		urls = append(urls, resp.Output.Results[i].URL)
	}

	return urls, nil
}

// TestConnection 测试连接
func (c *DashScopeClient) TestConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.client.Call(ctx, dashscope.Qwen, c.textModel, []dashscope.Message{
		{
			Role:    "user",
			Content: "你好",
		},
	}, nil)

	if err != nil {
		return fmt.Errorf("通义千问连接测试失败：%w", err)
	}

	return nil
}

// GetProviderName 返回提供商名称
func (c *DashScopeClient) GetProviderName() string {
	return "dashscope"
}

// GetSupportedModels 返回支持的模型列表
func (c *DashScopeClient) GetSupportedModels() map[string][]string {
	return map[string][]string{
		"text":  {"qwen-turbo", "qwen-plus", "qwen-max", "qwen-max-longcontext"},
		"image": {"wanx-v1", "wanx-v2"},
	}
}
