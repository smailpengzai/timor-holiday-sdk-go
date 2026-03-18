package holiday

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client 节假日API客户端
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient 创建新的客户端
func NewClient() *Client {
	return &Client{
		BaseURL: "https://timor.tech/api/holiday",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// WithBaseURL 自定义API地址
func (c *Client) WithBaseURL(url string) *Client {
	c.BaseURL = url
	return c
}

// WithHTTPClient 自定义HTTP客户端
func (c *Client) WithHTTPClient(client *http.Client) *Client {
	c.HTTPClient = client
	return c
}

// get 请求
func (c *Client) GetURL(endpoint string) ([]byte, error) {
	// 创建请求
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return nil, err
	}

	// 设置请求头，模拟浏览器访问
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return nil, err
	}

	return body, nil
}

// post 请求
func (c *Client) PostURL(endpoint string, data interface{}) ([]byte, error) {
	// 创建请求
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("JSON序列化失败: %w", err)
	}

	// 2. 创建请求
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return nil, err
	}

	// 设置请求头，模拟浏览器访问
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return nil, err
	}

	return body, nil
}
