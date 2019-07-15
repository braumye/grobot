package grobot

import (
    "bytes"
    "errors"
    "io"
    "net/http"
)

type Robot struct {
    Webhook              string
    ParseTextMessage     func(text string) ([]byte, error)
    ParseMarkdownMessage func(title string, text string) ([]byte, error)
    ParseResponseError   func(body io.Reader) error
}

func (this Robot) SendTextMessage(text string) error {
    body, err := this.ParseTextMessage(text)

    if err != nil {
        return errors.New("ParseTextFailed: " + err.Error())
    }

    return this.send(body)
}

func (this Robot) SendMarkdownMessage(title string, text string) error {
    body, err := this.ParseMarkdownMessage(title, text)

    if err != nil {
        return errors.New("ParseMarkdownFailed: " + err.Error())
    }

    return this.send(body)
}

func (this Robot) send(body []byte) error {
    req, reqerr := http.NewRequest("POST", this.Webhook, bytes.NewBuffer(body))

    if reqerr != nil {
        return errors.New("HttpRequestFailed: " + reqerr.Error())
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, resperr := client.Do(req)

    if resperr != nil {
        return errors.New("HttpResponseFailed: " + resperr.Error())
    }

    if resp != nil {
        defer resp.Body.Close()
    }

    return this.ParseResponseError(resp.Body)
}
