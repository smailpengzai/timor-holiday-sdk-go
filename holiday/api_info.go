package holiday

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetInfo 查询指定日期的节假日信息
// date格式: "2026-04-05", "today", 或空字符串表示今天
func (c *Client) GetInfo(date string) (*InfoResponse, error) {
	if date == "" {
		date = "today"
	}

	// 构建URL
	endpoint := fmt.Sprintf("%s/info/%s", c.BaseURL, url.PathEscape(date))

	// 发送请求
	resp, err := c.HTTPClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			return nil, ErrRateLimitExceeded
		}
		return nil, fmt.Errorf("HTTP错误: %s", resp.Status)
	}

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 解析JSON
	var result InfoResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w, 响应: %s", err, string(body))
	}

	// 检查API返回码
	if result.Code != 0 {
		return &result, fmt.Errorf("API错误: 代码=%d", result.Code)
	}

	return &result, nil
}
