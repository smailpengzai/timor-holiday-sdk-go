package holiday

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// GetYear 获取指定年份或月份的所有节假日信息
//
// 参数说明：
//
//	date: 指定年份或年月份。格式可以是：
//	      - "2026"       (查询2026年全年)
//	      - "2026/"      (查询2026年全年，推荐格式)
//	      - "2026-04"    (查询2026年4月)
//	      - "" 或 "today" (查询当前年份)
//	withType: 是否返回每个节假日的日期类型信息
//	withWeek: 节假日是否包含周末（周六周日也算作节假日）
//
// 注意事项（来自文档）：
// 1. 查询整年时，年份后加斜杠 "/" 是必须的（如 "2026/"）
// 2. 目前只配置了最多比当前时间往后一年的节假日
// 3. 如果 withWeek=true，则节日和周末冲突时，以节日为准
func (c *Client) GetYear(date string, withType, withWeek bool) (*YearResponse, error) {
	// 处理默认值
	if date == "" || date == "today" {
		date = "" // 使用接口默认行为，返回当前年份
	}

	// 构建URL路径
	endpoint := fmt.Sprintf("%s/year/", c.BaseURL)
	if date != "" {
		// 对路径部分进行编码
		endpoint += url.PathEscape(date)
	}

	// 构建查询参数
	queryParams := []string{}
	if withType {
		queryParams = append(queryParams, "type=Y")
	}
	if withWeek {
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
	var result YearResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w, 响应: %s", err, string(body))
	}

	// 检查API返回码
	if result.Code != 0 {
		return &result, fmt.Errorf("API错误: 代码=%d", result.Code)
	}

	return &result, nil
}

// YearRequest 是 GetYear 的另一种参数形式（可选，用于更清晰的API）
type YearRequest struct {
	Date     string // 年份或年月，如 "2026", "2026/", "2026-04"
	WithType bool   // 是否返回类型信息
	WithWeek bool   // 是否包含周末
}

// GetYearWithRequest 使用结构体参数的版本
func (c *Client) GetYearWithRequest(req YearRequest) (*YearResponse, error) {
	return c.GetYear(req.Date, req.WithType, req.WithWeek)
}
