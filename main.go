package grobot

import (
    "errors"
)

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
