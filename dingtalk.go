package grobot

// DingTalkTextMessage 钉钉机器人文本消息体
type DingTalkTextMessage struct {
	Content string `json:"content"`
}

// DingTalkMarkdownMessage 钉钉机器人 Markdown 消息体
type DingTalkMarkdownMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// 初始化钉钉机器人
// @see https://open-doc.dingtalk.com/microapp/serverapi2/qf2nxq
func newDingTalkRobot(token string) *Robot {
	return &Robot{
		Webhook:              stringBuilder("https://oapi.dingtalk.com/robot/send?access_token=", token),
		ParseTextMessage:     parseDingTalkTextMessage,
		ParseMarkdownMessage: parseDingTalkMarkdownMessage,
	}
}

// 钉钉机器人文本消息请求参数示例
// {
//     "msgtype": "text",
//     "text": {
//         "content": "我就是我, 是不一样的烟火@156xxxx8827"
//     },
//     "at": {
//         "atMobiles": [
//             "156xxxx8827",
//             "189xxxx8325"
//         ],
//         "isAtAll": false
//     }
// }
func parseDingTalkTextMessage(text string) map[string]interface{} {
	msg := DingTalkTextMessage{
		Content: text,
	}

	body := make(map[string]interface{})
	body["msgtype"] = "text"
	body["text"] = msg

	return body
}

// 钉钉机器人发送 Markdown 消息的接口参数示例
// {
//      "msgtype": "markdown",
//      "markdown": {
//          "title":"杭州天气",
//          "text": "#### 杭州天气 @156xxxx8827\n" +
//                  "> 9度，西北风1级，空气良89，相对温度73%\n\n" +
//                  "> ![screenshot](https://gw.alcts.com/zos/skylark-tools/public/files/84111bb.png)\n"  +
//                  "> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n"
//      },
//     "at": {
//         "atMobiles": [
//             "156xxxx8827",
//             "189xxxx8325"
//         ],
//         "isAtAll": false
//     }
//  }
func parseDingTalkMarkdownMessage(title string, text string) map[string]interface{} {
	msg := DingTalkMarkdownMessage{
		Title: title,
		Text:  text,
	}

	body := make(map[string]interface{})
	body["msgtype"] = "markdown"
	body["markdown"] = msg

	return body
}
