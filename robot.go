package grobot

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// Robot 消息机器人结构体, 消息处理都由 Robot 完成
type Robot struct {
	// Robot Webhook
	// dingtalk: https://oapi.dingtalk.com/robot/send?access_token=b0292a2506
	// wechatwork: https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=9b01-4b1b-abbc
	Webhook string

	// 将一段文本转化成机器人发送文本消息需要的接口参数
	ParseTextMessage func(text string) ([]byte, error)

	// 将标题和文本转化成机器人发送 Markdown 消息需要的接口参数
	ParseMarkdownMessage func(title string, text string) ([]byte, error)
}

// SendTextMessage 发送一条文本消息
func (robot Robot) SendTextMessage(text string) error {
	body, err := robot.ParseTextMessage(text)

	if err != nil {
		return errors.New("ParseTextFailed: " + err.Error())
	}

	return robot.send(body)
}

// SendMarkdownMessage 发送一条 Markdown 消息
func (robot Robot) SendMarkdownMessage(title string, text string) error {
	body, err := robot.ParseMarkdownMessage(title, text)

	if err != nil {
		return errors.New("ParseMarkdownFailed: " + err.Error())
	}

	return robot.send(body)
}

// WebhookResponse 调用 webhook 之后返回的消息体
type WebhookResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// 发送消息到 Webhook
func (robot Robot) send(body []byte) error {
	req, reqErr := http.NewRequest("POST", robot.Webhook, bytes.NewBuffer(body))

	if reqErr != nil {
		return errors.New("HttpRequestFailed: " + reqErr.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, respErr := client.Do(req)

	if respErr != nil {
		return errors.New("HttpResponseFailed: " + respErr.Error())
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	jsonResp := WebhookResponse{}
	decodeErr := json.NewDecoder(resp.Body).Decode(&jsonResp)

	if decodeErr != nil {
		return errors.New("HttpResponseBodyDecodeFailed: " + decodeErr.Error())
	}

	if jsonResp.ErrMsg != "ok" {
		return errors.New("SendMessageFailed: " + jsonResp.ErrMsg)
	}

	return nil
}
