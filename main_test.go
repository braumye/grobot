package grobot

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNew_Dingtalk(t *testing.T) {
    robot, _ := New("dingtalk", "token")
    assert.Equal(t, "https://oapi.dingtalk.com/robot/send?access_token=token", robot.Webhook)
}

func TestNew_WechatWork(t *testing.T) {
    robot, _ := New("wechatwork", "token")
    assert.Equal(t, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=token", robot.Webhook)
}

func TestNew_InvalidToken(t *testing.T) {
    _, err := New("dingtalk", "")
    assert.Equal(t, "invalid_token", err.Error())
}

func TestNew_UnsupportedDriver(t *testing.T) {
    _, err := New("test", "token")
    assert.Equal(t, "driver_unsupported", err.Error())
}
