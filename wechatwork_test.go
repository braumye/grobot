package grobot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 企业微信机器人把文本转化为文本消息体
func TestWechatWork_ParseTextMessage(t *testing.T) {
	want := `{"msgtype":"text","text":{"content":"test"}}`
	body := parseWechatWorkTextMessage("test")
	message, _ := json.Marshal(body)
	assert.Equal(t, want, string(message))
}

// 企业微信机器人把标题和文本转为化 Markdown 消息体
func TestWechatWork_ParseMarkdownMessage(t *testing.T) {
	want := `{"markdown":{"title":"title","text":"# H1"},"msgtype":"markdown"}`
	body := parseWechatWorkMarkdownMessage("title", "# H1")
	message, _ := json.Marshal(body)
	assert.Equal(t, want, string(message))
}
