package grobot

import (
	"encoding/json"
)

// WechatWorkTextMessage 企业微信机器人文本消息体
type WechatWorkTextMessage struct {
	Content string `json:"content"`
}

// WechatWorkMarkdownMessage 企业微信机器人 Markdown 消息体
type WechatWorkMarkdownMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// newWechatWorkRobot 初始化企业微信机器人
// @see https://work.weixin.qq.com/api/doc#90000/90136/91770
func newWechatWorkRobot(token string) *Robot {
	return &Robot{
		Webhook:              "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + token,
		ParseTextMessage:     parseDingTalkTextMessage,
		ParseMarkdownMessage: parseWechatWorkMarkdownMessage,
	}
}

// 请求参数示例
// {
//     "msgtype": "text",
//     "text": {
//         "content": "广州今日天气：29度，大部分多云，降雨概率：60%",
//         "mentioned_list":["wangqing","@all"],
//         "mentioned_mobile_list":["13800001111","@all"]
//     }
// }
func parseWechatWorkTextMessage(text string) ([]byte, error) {
	msg := WechatWorkTextMessage{
		Content: text,
	}

	body := make(map[string]interface{})
	body["msgtype"] = "text"
	body["text"] = msg

	return json.Marshal(body)
}

// 请求参数示例
// {
//     "msgtype": "markdown",
//     "markdown": {
//         "content": "实时新增用户反馈<font color=\"warning\">132例</font>，请相关同事注意。\n
//          >类型:<font color=\"comment\">用户反馈</font> \n
//          >普通用户反馈:<font color=\"comment\">117例</font> \n
//          >VIP用户反馈:<font color=\"comment\">15例</font>"
//     }
// }
func parseWechatWorkMarkdownMessage(title string, text string) ([]byte, error) {
	msg := WechatWorkMarkdownMessage{
		Title: title,
		Text:  text,
	}

	body := make(map[string]interface{})
	body["msgtype"] = "markdown"
	body["markdown"] = msg

	return json.Marshal(body)
}
