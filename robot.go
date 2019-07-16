package grobot

import (
    "bytes"
    "errors"
    "io"
    "net/http"
)

// 消息机器人结构体, 消息处理都由 Robot 完成
type Robot struct {
    // Robot Webhook
    // dingtalk: https://oapi.dingtalk.com/robot/send?access_token=b0292a2506
    // wechatwork: https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=9b01-4b1b-abbc
    Webhook string

    // 将一段文本转化成机器人发送文本消息需要的接口参数
    ParseTextMessage func(text string) ([]byte, error)

    // 将标题和文本转化成机器人发送 Markdown 消息需要的接口参数
    ParseMarkdownMessage func(title string, text string) ([]byte, error)

    // 处理 Webhook 返回结果, 用来判断是否发送成功, 发送失败会返回 nil
    ParseResponseError func(body io.Reader) error
}

// 发送一条文本消息
func (robot Robot) SendTextMessage(text string) error {
    body, err := robot.ParseTextMessage(text)

    if err != nil {
        return errors.New("ParseTextFailed: " + err.Error())
    }

    return robot.send(body)
}

// 发送一条 Markdown 消息
func (robot Robot) SendMarkdownMessage(title string, text string) error {
    body, err := robot.ParseMarkdownMessage(title, text)

    if err != nil {
        return errors.New("ParseMarkdownFailed: " + err.Error())
    }

    return robot.send(body)
}

// 发送消息到 Webhook
func (robot Robot) send(body []byte) error {
    req, reqerr := http.NewRequest("POST", robot.Webhook, bytes.NewBuffer(body))

    if reqerr != nil {
        return errors.New("HttpRequestFailed: " + reqerr.Error())
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, resperr := client.Do(req)

    if resperr != nil {
        return errors.New("HttpResponseFailed: " + resperr.Error())
    }

    if resp != nil {
        defer resp.Body.Close()
    }

    // 判断是否发送成功
    return robot.ParseResponseError(resp.Body)
}
