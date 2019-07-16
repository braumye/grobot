package grobot

import (
    "errors"
)

// New 创建一个消息机器人,
// driver 当前支持 dingtalk 和 wechatwork,
// dingtalk 对应钉钉 webhook 的 access_token 参数,
// wechatwork 对应企业微信 webhook 的 key 参数
func New(driver string, token string) (*Robot, error) {
    if token == "" {
        return nil, errors.New("invalid_token")
    }

    if driver == "dingtalk" {
        return newDingTalkRobot(token), nil
    }

    if driver == "wechatwork" {
        return newWechatWorkRobot(token), nil
    }

    return nil, errors.New("driver_unsupported")
}
