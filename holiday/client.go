package holiday

import (
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
