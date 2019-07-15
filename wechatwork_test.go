package grobot

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestWechatWork_ParseTextMessage(t *testing.T) {
    want := `{"msgtype":"text","text":{"content":"test"}}`
    body, _ := parseWechatWorkTextMessage("test")
    assert.Equal(t, want, string(body))
}

func TestWechatWork_ParseMarkdownMessage(t *testing.T) {
    want := `{"markdown":{"title":"title","text":"# H1"},"msgtype":"markdown"}`
    body, _ := parseWechatWorkMarkdownMessage("title", "# H1")
    assert.Equal(t, want, string(body))
}
