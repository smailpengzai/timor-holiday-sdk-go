package holiday

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// NextRequest 下一个节假日查询参数
type NextRequest struct {
	Date     string // 查询起点日期，空表示今天
	WithType bool   // 是否返回类型信息
	WithWeek bool   // 是否包含周末
}

// GetNextHoliday 获取下一个节假日
func (c *Client) GetNextHoliday(req NextRequest) (*NextResponse, error) {
	date := req.Date
	if date == "" {
		date = "today"
	}

	endpoint := fmt.Sprintf("%s/next/%s", c.BaseURL, url.PathEscape(date))

	// 添加查询参数
	queryParams := []string{}
	if req.WithType {
		queryParams = append(queryParams, "type=Y")
	}
	if req.WithWeek {
		queryParams = append(queryParams, "week=Y")
	}

	if len(queryParams) > 0 {
		endpoint += "?" + strings.Join(queryParams, "&")
	}

	// 发送请求
	resp, err := c.HTTPClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 处理响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	var result NextResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w, 响应: %s", err, string(body))
	}

	if result.Code != 0 {
		return &result, fmt.Errorf("API错误: 代码=%d", result.Code)
	}

	return &result, nil
}
