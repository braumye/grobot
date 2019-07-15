package grobot

import (
    "encoding/json"
    "errors"
    "io"
)

type WechatWorkTextMessage struct {
    Content string `json:"content"`
}

type WechatWorkMarkdownMessage struct {
    Title string `json:"title"`
    Text  string `json:"text"`
}

func newWechatWorkRobot(token string) *Robot {
    return &Robot{
        Webhook:              "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + token,
        ParseTextMessage:     parseDingTalkTextMessage,
        ParseMarkdownMessage: parseWechatWorkMarkdownMessage,
        ParseResponseError:   parseWechatWorkResponse,
    }
}

func parseWechatWorkTextMessage(text string) ([]byte, error) {
    msg := WechatWorkTextMessage{
        Content: text,
    }

    body := make(map[string]interface{})
    body["msgtype"] = "text"
    body["text"] = msg

    return json.Marshal(body)
}

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

type WechatWorkResponse struct {
    ErrCode int    `json:"errcode"`
    ErrMsg  string `json:"errmsg"`
}

func parseWechatWorkResponse(body io.Reader) error {
    jsonResp := WechatWorkResponse{}
    decodeErr := json.NewDecoder(body).Decode(&jsonResp)

    if decodeErr != nil {
        return errors.New("HttpResponseBodyDecodeFailed: " + decodeErr.Error())
    }

    if jsonResp.ErrMsg != "ok" {
        return errors.New("SendMessageFailed: " + jsonResp.ErrMsg)
    }

    return nil
}
