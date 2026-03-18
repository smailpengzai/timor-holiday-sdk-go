package holiday

import (
	"encoding/json"
	"fmt"
)

// GetTTS 获取综合放假安排描述
func (c *Client) GetTTS() (*TTSResponse, error) {
	endpoint := c.BaseURL + "/tts"
	return c.getTTSResponse(endpoint)
}

// GetTTSNext 获取下一个节假日描述
func (c *Client) GetTTSNext() (*TTSResponse, error) {
	endpoint := c.BaseURL + "/tts/next"
	return c.getTTSResponse(endpoint)
}

// GetTTSTomorrow 获取明天是否放假描述
func (c *Client) GetTTSTomorrow() (*TTSResponse, error) {
	endpoint := c.BaseURL + "/tts/tomorrow"
	return c.getTTSResponse(endpoint)
}

// getTTSResponse 内部通用的TTS请求方法
func (c *Client) getTTSResponse(endpoint string) (*TTSResponse, error) {
	body, err := c.GetURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	var result TTSResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	if result.Code != 0 {
		return &result, fmt.Errorf("API错误: 代码=%d", result.Code)
	}

	return &result, nil
}
