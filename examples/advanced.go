package main

import (
	"fmt"
	"github.com/smailpengzai/timor-holiday-sdk-go/holiday"
	"log"
)

func main() {
	client := holiday.NewClient()

	fmt.Println("=== 节假日API SDK 高级示例 ===")

	// 1. 查询今天是否上班
	fmt.Println("\n1. 查询今天状态:")
	today, err := client.GetInfo("today")
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Printf("  今天是 %s, 类型代码: %d\n", today.Type.Name, today.Type.Type)
		if today.Holiday != nil {
			fmt.Printf("  节假日: %s (%d倍工资)\n",
				today.Holiday.Name, today.Holiday.Wage)
		}
	}

	// 2. 查询清明节
	fmt.Println("\n2. 查询2026年清明节:")
	qingming, err := client.GetInfo("2026-04-05")
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else if qingming.Holiday != nil {
		fmt.Printf("  %s: %s, 工资倍数: %d\n",
			qingming.Holiday.Date,
			qingming.Holiday.Name,
			qingming.Holiday.Wage)
	}

	// 3. 批量查询
	fmt.Println("\n3. 批量查询多个日期:")
	batchResp, err := client.GetBatch(holiday.BatchRequest{
		Dates: []string{"2026-01-01", "2026-02-16", "2026-10-01"},
	})
	if err != nil {
		log.Printf("批量查询失败: %v", err)
	} else {
		for date, info := range batchResp.Holiday {
			if info != nil {
				fmt.Printf("  %s: %s (放假: %v)\n", date, info.Name, info.Holiday)
			} else {
				fmt.Printf("  %s: 非节假日\n", date)
			}
		}
	}

	// 4. 获取下一个节假日
	fmt.Println("\n4. 下一个节假日:")
	nextResp, err := client.GetNextHoliday(holiday.NextRequest{
		WithType: true,
	})
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else if nextResp.Holiday != nil {
		fmt.Printf("  下一个节日: %s (日期: %s, 还有%d天)\n",
			nextResp.Holiday.Name,
			nextResp.Holiday.Date,
			nextResp.Holiday.Rest)

		if nextResp.Workday != nil {
			fmt.Printf("  注意: 之前有调休: %s (日期: %s)\n",
				nextResp.Workday.Name,
				nextResp.Workday.Date)
		}
	}

	// 5. 获取TTS描述
	fmt.Println("\n5. 放假安排描述:")
	ttsResp, err := client.GetTTS()
	if err != nil {
		log.Printf("获取TTS失败: %v", err)
	} else {
		fmt.Printf("  %s\n", ttsResp.TTS)
	}

	// 6. 查询明年节假日
	fmt.Println("\n6. 2026年主要节假日:")
	yearResp, err := client.GetYear("2026", false, false)
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Println("  重要节假日:")
		for monthDay, info := range yearResp.Holiday {
			if info != nil && info.Wage == 3 {
				// 显示3倍工资的重要节日
				fmt.Printf("  %s: %s (%d倍工资)\n",
					monthDay, info.Name, info.Wage)
			}
		}
	}
}
