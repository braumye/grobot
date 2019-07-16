# GROBOT

[![Build Status](https://travis-ci.org/braumye/grobot.svg?branch=master)](https://travis-ci.org/braumye/grobot)
[![codecov](https://codecov.io/gh/braumye/grobot/branch/master/graph/badge.svg)](https://codecov.io/gh/braumye/grobot)
[![Go Report Card](https://goreportcard.com/badge/github.com/braumye/grobot)](https://goreportcard.com/report/github.com/braumye/grobot)

支持钉钉和企业微信的消息机器人

## DingTalk Example

```golang
robot, err := grobot.New("dingtalk", "your_dingtalk_access_token")

// 发送文本消息
err = robot.SendTextMessage("test message")

// 发送 Markdown 消息
err = robot.SendMarkdownMessage("markdown title", "# Markdown Text")
```

## Wechat Work Example

```golang
robot, err := grobot.New("wechatwork", "your_wechat_work_key")

// 发送文本消息
err = robot.SendTextMessage("test message")

// 发送 Markdown 消息
err = robot.SendMarkdownMessage("markdown title", "# Markdown Text")
```

