package grobot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 企业微信机器人把文本转化为文本消息体
func TestWechatWork_ParseTextMessage(t *testing.T) {
	want := `{"msgtype":"text","text":{"content":"test"}}`
	body, _ := parseWechatWorkTextMessage("test")
	assert.Equal(t, want, string(body))
}

// 企业微信机器人把标题和文本转为化 Markdown 消息体
func TestWechatWork_ParseMarkdownMessage(t *testing.T) {
	want := `{"markdown":{"title":"title","text":"# H1"},"msgtype":"markdown"}`
	body, _ := parseWechatWorkMarkdownMessage("title", "# H1")
	assert.Equal(t, want, string(body))
}
