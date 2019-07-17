package grobot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Robot 消息机器人结构体, 消息处理都由 Robot 完成
type Robot struct {
	// Robot Webhook
	// dingtalk: https://oapi.dingtalk.com/robot/send?access_token=b0292a2506
	// wechatwork: https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=9b01-4b1b-abbc
	Webhook string

	// 将一段文本转化成机器人发送文本消息需要的接口参数
	ParseTextMessage func(text string) map[string]interface{}

	// 将标题和文本转化成机器人发送 Markdown 消息需要的接口参数
	ParseMarkdownMessage func(title string, text string) map[string]interface{}
}

// SendTextMessage 发送一条文本消息
func (robot Robot) SendTextMessage(text string) error {
	return robot.send(robot.ParseTextMessage(text))
}

// SendMarkdownMessage 发送一条 Markdown 消息
func (robot Robot) SendMarkdownMessage(title string, text string) error {
	return robot.send(robot.ParseMarkdownMessage(title, text))
}

// WebhookResponse 调用 webhook 之后返回的消息体
type WebhookResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// 发送消息到 Webhook
func (robot Robot) send(body map[string]interface{}) error {
	message, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", robot.Webhook, bytes.NewBuffer(message))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, respErr := client.Do(req)

	if respErr != nil {
		return newError("HttpResponseFailed", respErr)
	}

	bodyBytes := []byte("")
	if resp != nil {
		defer resp.Body.Close()

		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}

	rawBody := string(bodyBytes)

	if resp.StatusCode != http.StatusOK {
		return newError("HttpResponseStatusCode", resp.StatusCode, ",RawBody:", rawBody)
	}

	jsonResp := WebhookResponse{}
	decodeErr := json.Unmarshal(bodyBytes, &jsonResp)

	if decodeErr != nil {
		return newError("HttpResponseBodyDecodeFailed", decodeErr, ",RawBody:", rawBody)
	}

	if jsonResp.ErrMsg != "ok" {
		return newError("SendMessageFailed", jsonResp.ErrMsg, ",RawBody:", rawBody)
	}

	return nil
}
