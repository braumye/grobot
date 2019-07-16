package grobot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 钉钉机器人把一段文本转化为文本消息体
func TestDingTalk_ParseTextMessage(t *testing.T) {
	want := `{"msgtype":"text","text":{"content":"test"}}`
	body := parseDingTalkTextMessage("test")
	message, _ := json.Marshal(body)
	assert.Equal(t, want, string(message))
}

// 钉钉机器人把标题和文本转化为 Markdown 消息体
func TestDingTalk_ParseMarkdownMessage(t *testing.T) {
	want := `{"markdown":{"title":"title","text":"# H1"},"msgtype":"markdown"}`
	body := parseDingTalkMarkdownMessage("title", "# H1")
	message, _ := json.Marshal(body)
	assert.Equal(t, want, string(message))
}
