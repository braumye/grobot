package grobot_test

import (
	"testing"

	"github.com/braumye/grobot"
	"github.com/stretchr/testify/assert"
)

// 新建一个钉钉机器人
func TestNew_Dingtalk(t *testing.T) {
	robot, _ := grobot.New("dingtalk", "token")
	assert.Equal(t, "https://oapi.dingtalk.com/robot/send?access_token=token", robot.Webhook)
}

// 新建一个企业微信机器人
func TestNew_WechatWork(t *testing.T) {
	robot, _ := grobot.New("wechatwork", "token")
	assert.Equal(t, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=token", robot.Webhook)
}

// 新建机器人的时候 Token 不能为空
func TestNew_InvalidToken(t *testing.T) {
	_, err := grobot.New("dingtalk", "")
	assert.Equal(t, "invalid_token", err.Error())
}

// 只支持钉钉和企业微信
func TestNew_UnsupportedDriver(t *testing.T) {
	_, err := grobot.New("test", "token")
	assert.Equal(t, "driver_unsupported", err.Error())
}

// 钉钉机器人发送文本消息
func TestDingTalkRobot_SendTextMessage(t *testing.T) {
	robot, _ := grobot.New("dingtalk", "token")
	err := robot.SendTextMessage("test")
	assert.Equal(t, `SendMessageFailed:token is not exist,RawBody:{"errcode":300001,"errmsg":"token is not exist"}`, err.Error())
}

// 钉钉机器人发送 Markdown 消息
func TestDingTalkRobot_SendMarkdownMessage(t *testing.T) {
	robot, _ := grobot.New("dingtalk", "token")
	err := robot.SendMarkdownMessage("title", "text")
	assert.Equal(t, `SendMessageFailed:token is not exist,RawBody:{"errcode":300001,"errmsg":"token is not exist"}`, err.Error())
}

// 企业微信机器人发送文本消息
func TestWechatWorkRobot_SendTextMessage(t *testing.T) {
	robot, _ := grobot.New("wechatwork", "token")
	err := robot.SendTextMessage("test")
	assert.Contains(t, err.Error(), "SendMessageFailed:invalid webhook url")
}

// 企业微信机器人发送 Markdown 消息
func TestWechatWorkRobot_SendMarkdownMessage(t *testing.T) {
	robot, _ := grobot.New("wechatwork", "token")
	err := robot.SendMarkdownMessage("title", "text")
	assert.Contains(t, err.Error(), "SendMessageFailed:invalid webhook url")
}
