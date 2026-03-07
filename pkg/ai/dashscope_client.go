package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// DashScopeClient 通义千问客户端（HTTP 直连版本）
// 不依赖 alibabacloud-dashscope-go-sdk，使用标准 HTTP 调用
type DashScopeClient struct {
	apiKey     string
	textModel  string
	imageModel string
	timeout    time.Duration
	httpClient *http.Client
}

// DashScopeRequest 通义千问请求
type DashScopeRequest struct {
	Model string          `json:"model"`
	Input DashScopeInput  `json:"input"`
	Parameters DashScopeParameters `json:"parameters,omitempty"`
}

// DashScopeInput 通义千问输入
type DashScopeInput struct {
	Messages []DashScopeMessage `json:"messages"`
	Prompt   string             `json:"prompt,omitempty"`
}

// DashScopeMessage 通义千问消息
type DashScopeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// DashScopeParameters 通义千问参数
type DashScopeParameters struct {
	ResultFormat string `json:"result_format,omitempty"`
}

// DashScopeResponse 通义千问响应
type DashScopeResponse struct {
	Output struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Text   string `json:"text"`
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	} `json:"output"`
	Usage struct {
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
	RequestID string `json:"request_id"`
}

// NewDashScopeClient 创建通义千问客户端
func NewDashScopeClient(apiKey, textModel, imageModel string, timeoutSec int) *DashScopeClient {
	return &DashScopeClient{
		apiKey:     apiKey,
		textModel:  textModel,
		imageModel: imageModel,
		timeout:    time.Duration(timeoutSec) * time.Second,
		httpClient: &http.Client{
			Timeout: time.Duration(timeoutSec) * time.Second,
		},
	}
}

// GenerateText 生成文本（通义千问）
func (c *DashScopeClient) GenerateText(prompt string, systemPrompt string, options ...func(*ChatCompletionRequest)) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	// 构建请求
	url := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
	
	messages := []DashScopeMessage{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: prompt},
	}

	reqBody := DashScopeRequest{
		Model: c.textModel,
		Input: DashScopeInput{Messages: messages},
		Parameters: DashScopeParameters{ResultFormat: "message"},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("序列化请求失败：%w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("创建请求失败：%w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("通义千问调用失败：%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败：%w", err)
	}

	var dashscopeResp DashScopeResponse
	if err := json.Unmarshal(body, &dashscopeResp); err != nil {
		return "", fmt.Errorf("解析响应失败：%w", err)
	}

	// 检查响应
	if len(dashscopeResp.Output.Choices) > 0 {
		return dashscopeResp.Output.Choices[0].Message.Content, nil
	}
	
	if dashscopeResp.Output.Text != "" {
		return dashscopeResp.Output.Text, nil
	}

	return "", errors.New("通义千问返回空响应")
}

// GenerateImage 生成图像（通义万相）
func (c *DashScopeClient) GenerateImage(prompt string, size string, n int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	// 通义万相 API
	url := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-to-image/generation"
	
	reqBody := DashScopeRequest{
		Model: c.imageModel,
		Input: DashScopeInput{Prompt: prompt},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败：%w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败：%w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("通义万相调用失败：%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var dashscopeResp DashScopeResponse
	if err := json.Unmarshal(body, &dashscopeResp); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	urls := make([]string, 0, n)
	for _, img := range dashscopeResp.Output.Images {
		if img.URL != "" {
			urls = append(urls, img.URL)
		}
	}

	if len(urls) == 0 {
		return nil, errors.New("通义万相返回空响应")
	}

	return urls, nil
}

// TestConnection 测试连接
func (c *DashScopeClient) TestConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.GenerateTextWithContext(ctx, "你好", "你是一个助手")
	if err != nil {
		return fmt.Errorf("通义千问连接测试失败：%w", err)
	}

	return nil
}

// GenerateTextWithContext 带上下文的文本生成
func (c *DashScopeClient) GenerateTextWithContext(ctx context.Context, prompt string, systemPrompt string) (string, error) {
	url := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
	
	messages := []DashScopeMessage{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: prompt},
	}

	reqBody := DashScopeRequest{
		Model: c.textModel,
		Input: DashScopeInput{Messages: messages},
	}

	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var dashscopeResp DashScopeResponse
	json.Unmarshal(body, &dashscopeResp)

	if len(dashscopeResp.Output.Choices) > 0 {
		return dashscopeResp.Output.Choices[0].Message.Content, nil
	}
	
	if dashscopeResp.Output.Text != "" {
		return dashscopeResp.Output.Text, nil
	}

	return "", errors.New("通义千问返回空响应")
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
