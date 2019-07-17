package grobot

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 机器人发送文本消息
func TestRobot_SendTextMessage(t *testing.T) {
	want := `{"msgtype":"text","text":{"content":"test"}}`
	ts := testHttp(t, want, `{"errmsg":"ok","errcode":0}`, http.StatusOK)
	defer ts.Close()

	robot := getTestRobot(ts.URL)
	err := robot.SendTextMessage("test")
	assert.Nil(t, err)
}

// 机器人发送 Markdown 消息
func TestRobot_SendMarkdownMessage(t *testing.T) {
	want := `{"markdown":{"title":"title","text":"text"},"msgtype":"markdown"}`
	ts := testHttp(t, want, `{"errmsg":"ok","errcode":0}`, http.StatusOK)
	defer ts.Close()

	robot := getTestRobot(ts.URL)
	err := robot.SendMarkdownMessage("title", "text")
	assert.Nil(t, err)
}

func TestRobotSendMessageFailed_WebhookReturnEmpty(t *testing.T) {
	want := `{"markdown":{"title":"title","text":"text"},"msgtype":"markdown"}`
	ts := testHttp(t, want, "", http.StatusOK)
	defer ts.Close()

	robot := getTestRobot(ts.URL)
	err := robot.SendMarkdownMessage("title", "text")
	assert.Equal(t, "HttpResponseBodyDecodeFailed:unexpected end of JSON input,RawBody:", err.Error())
}

func TestRobotSendMessageFailed_ResponseUnsuccessfully(t *testing.T) {
	want := `{"markdown":{"title":"title","text":"text"},"msgtype":"markdown"}`
	ts := testHttp(t, want, `{"errcode":400,"errmsg":"bas request"}`, http.StatusBadRequest)
	defer ts.Close()

	robot := getTestRobot(ts.URL)
	err := robot.SendMarkdownMessage("title", "text")
	assert.Equal(t, `HttpResponseStatusCode:400,RawBody:{"errcode":400,"errmsg":"bas request"}`, err.Error())
}

// mock http client
func testHttp(t *testing.T, want string, resp string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
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
		Webhook:              stringBuilder(api, "/send?token=token"),
		ParseTextMessage:     testTestBodyFunc,
		ParseMarkdownMessage: testMarkdownParser,
	}
}

type testTextMessage struct {
	Content string `json:"content"`
}

func testTestBodyFunc(text string) map[string]interface{} {
	msg := testTextMessage{text}
	body := make(map[string]interface{})
	body["msgtype"] = "text"
	body["text"] = msg

	return body
}

type testMarkdownMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func testMarkdownParser(title string, text string) map[string]interface{} {
	msg := testMarkdownMessage{
		Title: title,
		Text:  text,
	}

	body := make(map[string]interface{})
	body["msgtype"] = "markdown"
	body["markdown"] = msg

	return body
}
