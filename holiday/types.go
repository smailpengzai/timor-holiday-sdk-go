package holiday

// 根据文档和链接内容定义的所有数据结构

// HolidayInfo 节假日详细信息 (对应 /info 接口的 holiday 字段)
type HolidayInfo struct {
	Holiday bool   `json:"holiday"`          // 是否为节假日
	Name    string `json:"name"`             // 节假日名称
	Wage    int    `json:"wage"`             // 工资倍数
	Date    string `json:"date"`             // 日期
	Rest    int    `json:"rest"`             // 距离天数 (可能为0)
	After   *bool  `json:"after,omitempty"`  // 调休方向: false=先调休, true=后调休
	Target  string `json:"target,omitempty"` // 关联的节假日
}

// TypeInfo 日期类型信息
type TypeInfo struct {
	Type int    `json:"type"` // 0:工作日, 1:周末, 2:节假日, 3:调休
	Name string `json:"name"` // 类型名称
	Week int    `json:"week"` // 星期几 (1-7)
}

// InfoResponse /info 接口返回结构
type InfoResponse struct {
	Code    int          `json:"code"`    // 0=成功, -1=失败
	Type    TypeInfo     `json:"type"`    // 日期类型
	Holiday *HolidayInfo `json:"holiday"` // 节假日详情 (可能为nil)
}

// BatchResponse /batch 接口返回结构
type BatchResponse struct {
	Code    int                     `json:"code"`
	Holiday map[string]*HolidayInfo `json:"holiday"`        // 日期为key
	Type    map[string]TypeInfo     `json:"type,omitempty"` // 当type=Y时存在
}

// NextResponse /next 接口返回结构
type NextResponse struct {
	Code    int          `json:"code"`
	Holiday *HolidayInfo `json:"holiday"` // 下一个节假日
	Workday *HolidayInfo `json:"workday"` // 调休信息 (可能为nil)
	Type    struct {
		Holiday *TypeInfo `json:"holiday,omitempty"`
		Workday *TypeInfo `json:"workday,omitempty"`
	} `json:"type,omitempty"` // 当type=Y时存在
}

// WorkdayResponse /workday/next 接口返回结构
type WorkdayResponse struct {
	Code    int `json:"code"`
	Workday *struct {
		Type int    `json:"type"` // 0或3
		Name string `json:"name"`
		Week int    `json:"week"`
		Date string `json:"date"`
		Rest int    `json:"rest"`
	} `json:"workday"` // 可能为nil
}

// YearResponse /year 接口返回结构
type YearResponse struct {
	Code    int                     `json:"code"`
	Holiday map[string]*HolidayInfo `json:"holiday"`        // key格式: "MM-DD"
	Type    map[string]TypeInfo     `json:"type,omitempty"` // 当type=Y时存在
}

// TTSResponse tts 系列接口返回结构
type TTSResponse struct {
	Code int    `json:"code"`
	TTS  string `json:"tts"`
}
