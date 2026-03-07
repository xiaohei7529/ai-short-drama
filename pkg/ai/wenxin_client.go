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

// WenXinClient 文心一言客户端
type WenXinClient struct {
	apiKey    string
	secretKey string
	textModel string
	timeout   time.Duration
	httpClient *http.Client
}

// WenXinTokenResponse 获取 access_token 响应
type WenXinTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Error       string `json:"error,omitempty"`
}

// WenXinRequest 文心一言请求
type WenXinRequest struct {
	Messages []WenXinMessage `json:"messages"`
}

// WenXinMessage 文心一言消息
type WenXinMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// WenXinResponse 文心一言响应
type WenXinResponse struct {
	Result string `json:"result"`
	ErrorCode int `json:"error_code,omitempty"`
	ErrorMsg string `json:"error_msg,omitempty"`
}

// NewWenXinClient 创建文心一言客户端
func NewWenXinClient(apiKey, secretKey, textModel string, timeoutSec int) *WenXinClient {
	return &WenXinClient{
		apiKey:    apiKey,
		secretKey: secretKey,
		textModel: textModel,
		timeout:   time.Duration(timeoutSec) * time.Second,
		httpClient: &http.Client{
			Timeout: time.Duration(timeoutSec) * time.Second,
		},
	}
}

// GetAccessToken 获取 access_token
func (c *WenXinClient) GetAccessToken() (string, error) {
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s",
		c.apiKey, c.secretKey)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("获取 access_token 失败：%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败：%w", err)
	}

	var tokenResp WenXinTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("解析响应失败：%w", err)
	}

	if tokenResp.Error != "" {
		return "", errors.New(tokenResp.Error)
	}

	return tokenResp.AccessToken, nil
}

// GenerateText 生成文本
func (c *WenXinClient) GenerateText(prompt string, systemPrompt string, options ...func(*ChatCompletionRequest)) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	// 获取 access_token
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}

	// 构建请求
	url := fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/%s?access_token=%s",
		c.textModel, accessToken)

	messages := []WenXinMessage{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: prompt},
	}

	reqBody := WenXinRequest{Messages: messages}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("序列化请求失败：%w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("创建请求失败：%w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("文心一言调用失败：%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败：%w", err)
	}

	var wenxinResp WenXinResponse
	if err := json.Unmarshal(body, &wenxinResp); err != nil {
		return "", fmt.Errorf("解析响应失败：%w", err)
	}

	if wenxinResp.ErrorCode != 0 {
		return "", fmt.Errorf("文心一言错误 [%d]: %s", wenxinResp.ErrorCode, wenxinResp.ErrorMsg)
	}

	return wenxinResp.Result, nil
}

// GenerateImage 文心一言暂不支持图像生成，返回错误
func (c *WenXinClient) GenerateImage(prompt string, size string, n int) ([]string, error) {
	return nil, errors.New("文心一言暂不支持图像生成，请使用通义万相或 DALL-E 3")
}

// TestConnection 测试连接
func (c *WenXinClient) TestConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.GenerateTextWithContext(ctx, "你好", "你是一个助手")
	if err != nil {
		return fmt.Errorf("文心一言连接测试失败：%w", err)
	}

	return nil
}

// GenerateTextWithContext 带上下文的文本生成
func (c *WenXinClient) GenerateTextWithContext(ctx context.Context, prompt string, systemPrompt string) (string, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/%s?access_token=%s",
		c.textModel, accessToken)

	messages := []WenXinMessage{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: prompt},
	}

	reqBody := WenXinRequest{Messages: messages}
	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var wenxinResp WenXinResponse
	json.Unmarshal(body, &wenxinResp)

	if wenxinResp.ErrorCode != 0 {
		return "", fmt.Errorf("文心一言错误 [%d]: %s", wenxinResp.ErrorCode, wenxinResp.ErrorMsg)
	}

	return wenxinResp.Result, nil
}

// GetProviderName 返回提供商名称
func (c *WenXinClient) GetProviderName() string {
	return "wenxin"
}

// GetSupportedModels 返回支持的模型列表
func (c *WenXinClient) GetSupportedModels() map[string][]string {
	return map[string][]string{
		"text": {"ernie-bot", "ernie-bot-turbo", "ernie-bot-4", "ernie-bot-8k"},
	}
}
