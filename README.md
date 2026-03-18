好的，根据您提供的【链接内容】（timor.tech 节假日 API 的使用说明和完整接口文档），以及“README.md 这个文件需要重写”的需求，我将为您撰写一份适用于开源 SDK 项目的、详尽且规范的 README.md 文档。

这份文档将清晰阐明：
1.  **项目用途**：此 Go SDK 是什么，解决什么问题。
2.  **出处**：明确 SDK 所封装的 API 来源。
3.  **使用限制**：准确转述原 API 的服务条款、限额和重要通知。
4.  **VIP开通**：指引商业用户。
5.  **开源声明**：为您的代码仓库添加合适的许可证。

---

# timor-holiday-sdk-go

https://pkg.go.dev/badge/github.com/yourusername/timor-holiday-sdk-go.svg](https://pkg.go.dev/github.com/yourusername/timor-holiday-sdk-go)
https://img.shields.io/badge/License-MIT-yellow.svg](https://opensource.org/licenses/MIT)
https://goreportcard.com/badge/github.com/yourusername/timor-holiday-sdk-go](https://goreportcard.com/report/github.com/yourusername/timor-holiday-sdk-go)

> 一个非官方的、功能完整的 Go 语言 SDK，用于访问 https://timor.tech/api/holiday/ 的免费节假日 API 服务。

## 📋 项目简介 (Purpose)

`timor-holiday-sdk-go` 是一款专为 Go 开发者设计的客户端库。它将 https://timor.tech/api/holiday/ 提供的、功能强大的免费节假日查询 API 进行了完整的、类型安全的封装。

**它能帮你做什么？**
- 在应用中轻松判断任意日期是工作日、周末、法定节假日还是调休日。
- 获取下一个节假日或工作日的详细信息。
- 批量查询多个日期的状态，高效处理排期、考勤、薪资计算等场景。
- 获得人性化的放假安排语音文本（TTS），适用于智能语音助手或消息推送。
- 查询整年或整月的节假日安排，用于生成日历或规划。

**核心价值**：开发者无需再手动处理 HTTP 请求、解析 JSON 或记忆复杂的 API 参数，通过简洁的 Go 方法和结构体即可使用所有功能。

## ⚠️ 重要声明：关于底层 API 服务 (Source & Attribution)

**本 SDK 封装的服务并非由本仓库维护者提供。**

本 SDK 完全依赖并调用由 **timor.tech 站长（QQ: 1260022720）** 个人开发并免费提供的节假日 API 服务。

- **官方服务地址**：https://timor.tech/api/holiday/
- **服务性质**：这是一个由个人利用业余时间开发和维护的免费公共服务。数据会根据国务院公告每年更新。
- **致谢**：感谢原作者的辛勤付出和无私分享。请大家在使用时尊重其服务条款（见下文“使用限制”），合理使用，切勿滥用。

**本 SDK 项目仅作为该 API 的 Go 语言客户端存在，不保证底层 API 服务的永久可用性、数据绝对准确性或响应速度。相关问题请直接联系原服务作者。**

## 🚦 使用限制 (Usage Limits & Terms)

**以下限制条款直接来源于 timor.tech 官方说明，使用本 SDK 即视为同意这些条款。** 本 SDK 的调用行为将直接消耗您在原服务的调用额度。

### 1. 调用频率限制 (Rate Limits)
为了保障服务稳定，原服务对所有匿名用户（通过IP识别）实施以下24小时滚动窗口限额：

- **节假日接口**（如 `/info`, `/batch`, `/next` 等）：`10,000 次/IP/24小时`
- **文字（TTS）接口**（如 `/tts`, `/tts/next` 等）：`100 次/IP/24小时`

**重要**：如果您处于小区宽带等共享IP的网络环境，您将与同一出口IP下的其他用户共享此限额。

### 2. 服务状态与稳定性
- **攻击与防御**：该服务长期遭受 CC（ChallengeCollapsar）攻击。尽管已接入 CDN 加速和防御，访问速度可能受影响，且部分用户可能会被要求进行浏览器验证。
- **VIP 服务**：对于对稳定性和速度有**100%要求**的企业用户，原服务提供付费 VIP 方案（见下文），可免受限速和验证影响。

### 3. 数据范围
- 目前最多配置了比当前时间**往后一年**的节假日数据。
- 历史数据可追溯。

### 4. 错误处理
- 超过调用限额后，API 将返回 **HTTP 429 (Too Many Requests)** 状态码。本 SDK 会将其转换为 `ErrRateLimitExceeded` 错误。
- 服务端错误返回 `code: -1`。本 SDK 会检查此字段并返回相应错误。

## 💼 VIP / 企业服务开通 (VIP & Enterprise)

如果您的企业应用有高并发、高稳定性需求，或希望获得更快的响应速度和无验证的访问体验，**请直接联系原服务作者开通 VIP**。

**VIP开通地址与方式**：
1.  **联系站长**：
    - **QQ**: `1260022720` （添加时请注明“购买节假日服务”）
    - **邮箱**: `bluescode@outlook.com`
2.  **服务版本**：
    - **无忧企业版**：获取 Token 或设置 IP 白名单，绕过限速。**Token 严禁在客户端（如App、网页）代码中使用**，仅限在内网服务器代理访问。
    - **内网部署版**：获取独立的二进制可执行文件，部署在您的内网，平均响应时间约 2ms。数据更新后会主动通知。

**购买 VIP 是对原服务作者最直接的支持，能帮助这个优秀的免费服务长期稳定运行下去。**

## 🚀 快速开始 (Quick Start)

### 安装
```bash
go get github.com/yourusername/timor-holiday-sdk-go
```

### 基础用法
```go
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/timor-holiday-sdk-go/holiday"
)

func main() {
    // 1. 创建客户端
    client := holiday.NewClient() // 默认使用 HTTPS

    // 2. 查询今天是否需要上班
    resp, err := client.GetInfo("today")
    if err != nil {
        log.Fatalf("查询失败: %v", err)
    }

    // 3. 解析结果
    // resp.Type.Type 含义: 0=工作日, 1=周末, 2=节假日, 3=调休
    switch resp.Type.Type {
    case 0, 3:
        fmt.Printf("今天是%s，要上班。\n", resp.Type.Name)
    case 1, 2:
        fmt.Printf("今天是%s，休息啦！\n", resp.Type.Name)
        if resp.Holiday != nil {
            fmt.Printf("  节日快乐！今天是%s，加班有%d倍工资哦。\n", resp.Holiday.Name, resp.Holiday.Wage)
        }
    }
}
```

## 📖 功能特性与接口覆盖 (Features)

本 SDK 完整实现了 timor.tech 节假日 API 文档中列出的所有接口：

| 接口功能 | 方法 | 说明 |
| :--- | :--- | :--- |
| **查询单日信息** | `client.GetInfo(date)` | 获取指定日期的工作日/节假日状态及详情。 |
| **批量查询** | `client.GetBatch(req)` | 最多批量查询50个日期的状态。 |
| **下一个节假日** | `client.GetNextHoliday(req)` | 获取下一个节假日，包含调休信息。 |
| **下一个工作日** | `client.GetNextWorkday(date)` | 获取下一个工作日（含调休）。 |
| **全年/月节假日** | `client.GetYear(date, withType, withWeek)` | 获取指定年或月的所有节假日列表。 |
| **放假安排描述** | `client.GetTTS()` | 获取人性化的近期放假语音文本。 |
| **下一个节日描述** | `client.GetTTSNext()` | 获取下一个节假日的描述文本。 |
| **明天是否放假** | `client.GetTTSTomorrow()` | 获取关于明天是否放假的回答文本。 |

## 🔧 高级配置

### 自定义 HTTP 客户端
```go
import (
    "net/http"
    "time"
    "github.com/yourusername/timor-holiday-sdk-go/holiday"
)

client := holiday.NewClient().
    WithHTTPClient(&http.Client{
        Timeout: 30 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns: 10,
        },
    })
```

### 批量查询示例
```go
// 查询多个重要日期
batchResp, err := client.GetBatch(holiday.BatchRequest{
    Dates: []string{"2026-01-01", "2026-02-16", "2026-10-01"},
})
if err != nil {
    log.Fatal(err)
}
for date, info := range batchResp.Holiday {
    if info != nil {
        fmt.Printf("%s: %s (Wage: %dx)\n", date, info.Name, info.Wage)
    }
}
```

## 📄 开源协议 (Open Source License)

本项目采用 **MIT 许可证** 发布。

```text
MIT License

Copyright (c) 2025 YourName

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

**重要**：此许可证仅适用于 **本 SDK 的代码**（即 `*.go` 文件及其组织方式）。您对 https://timor.tech/api/holiday/ API 的调用行为，仍需完全遵守其网站上声明的使用条款。

## 🤝 贡献与支持 (Contributing)

欢迎提交 Issue 和 Pull Request 来完善此 SDK。

**注意**：关于底层 API 的数据错误、功能建议、VIP 购买或服务稳定性问题，**请直接联系原服务作者**（QQ: 1260022720），因为本仓库无法解决此类问题。

---
**最后，请再次合理、友善地使用此免费 API 服务。如果觉得有用，可以考虑https://timor.tech/api/holiday/以支持其长期维护。**