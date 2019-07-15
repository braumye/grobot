package grobot

import (
    "encoding/json"
)

type DingTalkTextMessage struct {
    Content string `json:"content"`
}

type DingTalkMarkdownMessage struct {
    Title string `json:"title"`
    Text  string `json:"text"`
}

func newDingTalkRobot(token string) *Robot {
    return &Robot{
        Webhook:          "https://oapi.dingtalk.com/robot/send?access_token=" + token,
        ParseTextMessage: parseDingTalkTextMessage,
    }
}

func parseDingTalkTextMessage(text string) ([]byte, error) {
    msg := DingTalkTextMessage{
        Content: text,
    }

    body := make(map[string]interface{})
    body["msgtype"] = "text"
    body["text"] = msg

    return json.Marshal(body)
}

func parseDingTalkMarkdownMessage(title string, text string) ([]byte, error) {
    msg := DingTalkMarkdownMessage{
        Title: title,
        Text:  text,
    }

    body := make(map[string]interface{})
    body["msgtype"] = "markdown"
    body["markdown"] = msg

    return json.Marshal(body)
}
