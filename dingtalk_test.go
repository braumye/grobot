package grobot

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestDingTalk_ParseTextMessage(t *testing.T) {
    want := `{"msgtype":"text","text":{"content":"test"}}`
    body, _ := parseDingTalkTextMessage("test")
    assert.Equal(t, want, string(body))
}

func TestDingTalk_ParseMarkdownMessage(t *testing.T) {
    want := `{"markdown":{"title":"title","text":"# H1"},"msgtype":"markdown"}`
    body, _ := parseDingTalkMarkdownMessage("title", "# H1")
    assert.Equal(t, want, string(body))
}
