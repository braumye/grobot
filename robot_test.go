package grobot

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 机器人发送文本消息
func TestRobot_SendTextMessage(t *testing.T) {
	want := `{"msgtype":"text","text":{"content":"test"}}`
	ts := testHttp(t, want, `{"errmsg":"ok","errcode":0}`)
	defer ts.Close()

	robot := getTestRobot(ts.URL)
	err := robot.SendTextMessage("test")
	assert.Nil(t, err)
}

// 机器人发送 Markdown 消息
func TestRobot_SendMarkdownMessage(t *testing.T) {
	want := `{"markdown":{"title":"title","text":"text"},"msgtype":"markdown"}`
	ts := testHttp(t, want, `{"errmsg":"ok","errcode":0}`)
	defer ts.Close()

	robot := getTestRobot(ts.URL)
	err := robot.SendMarkdownMessage("title", "text")
	assert.Nil(t, err)
}

// 钉钉机器人发送文本消息
func TestDingTalkRobot_SendTextMessage(t *testing.T) {
	robot, _ := New("dingtalk", "token")
	err := robot.SendTextMessage("test")
	assert.Contains(t, "SendMessageFailed: token is not exist", err.Error())
}

// 钉钉机器人发送 Markdown 消息
func TestDingTalkRobot_SendMarkdownMessage(t *testing.T) {
	robot, _ := New("dingtalk", "token")
	err := robot.SendMarkdownMessage("title", "text")
	assert.Contains(t, "SendMessageFailed: token is not exist", err.Error())
}

// 企业微信机器人发送文本消息
func TestWechatWorkRobot_SendTextMessage(t *testing.T) {
	robot, _ := New("wechatwork", "token")
	err := robot.SendTextMessage("test")
	assert.Contains(t, err.Error(), "SendMessageFailed: invalid webhook url")
}

// 企业微信机器人发送 Markdown 消息
func TestWechatWorkRobot_SendMarkdownMessage(t *testing.T) {
	robot, _ := New("wechatwork", "token")
	err := robot.SendMarkdownMessage("title", "text")
	assert.Contains(t, err.Error(), "SendMessageFailed: invalid webhook url")
}

// mock http client
func testHttp(t *testing.T, want string, resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))

		assert.Equal(t, "POST", r.Method)
		b, _ := ioutil.ReadAll(r.Body)

		assert.Equal(t, want, string(b))

		r.ParseForm()
		assert.Equal(t, "token", r.Form.Get("token"))
	}))
}

func getTestRobot(api string) *Robot {
	return &Robot{
		Webhook:              api + "/send?token=token",
		ParseTextMessage:     testTestBodyFunc,
		ParseMarkdownMessage: testMarkdownParser,
	}
}

type testTextMessage struct {
	Content string `json:"content"`
}

func testTestBodyFunc(text string) ([]byte, error) {
	msg := testTextMessage{text}
	body := make(map[string]interface{})
	body["msgtype"] = "text"
	body["text"] = msg

	return json.Marshal(body)
}

type testMarkdownMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func testMarkdownParser(title string, text string) ([]byte, error) {
	msg := testMarkdownMessage{
		Title: title,
		Text:  text,
	}

	body := make(map[string]interface{})
	body["msgtype"] = "markdown"
	body["markdown"] = msg

	return json.Marshal(body)
}
