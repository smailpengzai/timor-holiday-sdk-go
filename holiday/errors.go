package holiday

import "errors"

var (
	// ErrRateLimitExceeded 超出API调用限制
	ErrRateLimitExceeded = errors.New("API调用频率超限，请24小时后再试")

	// ErrInvalidDate 日期格式错误
	ErrInvalidDate = errors.New("日期格式错误，应为YYYY-MM-DD或YYYY-M-D格式")

	// ErrAPIUnavailable API服务不可用
	ErrAPIUnavailable = errors.New("节假日API服务暂时不可用")
)
