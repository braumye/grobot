package grobot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 新建一个钉钉机器人
func TestNew_Dingtalk(t *testing.T) {
	robot, _ := New("dingtalk", "token")
	assert.Equal(t, "https://oapi.dingtalk.com/robot/send?access_token=token", robot.Webhook)
}

// 新建一个企业微信机器人
func TestNew_WechatWork(t *testing.T) {
	robot, _ := New("wechatwork", "token")
	assert.Equal(t, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=token", robot.Webhook)
}

// 新建机器人的时候 Token 不能为空
func TestNew_InvalidToken(t *testing.T) {
	_, err := New("dingtalk", "")
	assert.Equal(t, "invalid_token", err.Error())
}

// 只支持钉钉和企业微信
func TestNew_UnsupportedDriver(t *testing.T) {
	_, err := New("test", "token")
	assert.Equal(t, "driver_unsupported", err.Error())
}
