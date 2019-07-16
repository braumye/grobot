package grobot

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

// 钉钉机器人把一段文本转化为文本消息体
func TestDingTalk_ParseTextMessage(t *testing.T) {
    want := `{"msgtype":"text","text":{"content":"test"}}`
    body, _ := parseDingTalkTextMessage("test")
    assert.Equal(t, want, string(body))
}

// 钉钉机器人把标题和文本转化为 Markdown 消息体
func TestDingTalk_ParseMarkdownMessage(t *testing.T) {
    want := `{"markdown":{"title":"title","text":"# H1"},"msgtype":"markdown"}`
    body, _ := parseDingTalkMarkdownMessage("title", "# H1")
    assert.Equal(t, want, string(body))
}
