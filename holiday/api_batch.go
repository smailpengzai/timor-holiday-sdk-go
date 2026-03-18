package holiday

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// BatchRequest 批量查询请求参数
type BatchRequest struct {
	Dates    []string // 要查询的日期列表
	WithType bool     // 是否返回类型信息
}

// GetBatch 批量查询节假日信息
func (c *Client) GetBatch(req BatchRequest) (*BatchResponse, error) {
	if len(req.Dates) == 0 {
		return nil, fmt.Errorf("至少需要一个日期")
	}
	if len(req.Dates) > 50 {
		return nil, fmt.Errorf("最多支持50个日期")
	}

	// 构建查询参数
	params := make([]string, 0, len(req.Dates))
	for _, date := range req.Dates {
		params = append(params, fmt.Sprintf("d=%s", url.QueryEscape(date)))
	}

	endpoint := fmt.Sprintf("%s/batch?%s", c.BaseURL, strings.Join(params, "&"))

	if req.WithType {
		endpoint += "&type=Y"
	}

	// 发送请求
	body, err := c.GetURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	var result BatchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	if result.Code != 0 {
		return &result, fmt.Errorf("API错误: 代码=%d", result.Code)
	}

	return &result, nil
}
