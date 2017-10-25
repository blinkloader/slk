package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/yarikbratashchuk/slk/internal/config"
)

type chatHistory struct {
	Ok       bool       `json:"ok"`
	Messages []*Message `json:"messages"`
}

type Message struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
	Ts   string `json:"ts"`
}

func GetChatHistory(conf config.Config) (hist []*Message, err error) {
	data := url.Values{}
	data.Set("token", conf.Token)
	data.Set("channel", conf.Channel)
	data.Set("limit", "10")

	res, err := http.PostForm("https://slack.com/api/conversations.history", data)
	if err != nil {
		return
	}

	var h chatHistory
	if err = json.NewDecoder(res.Body).Decode(&h); err != nil {
		return
	}
	defer res.Body.Close()

	if !h.Ok {
		err = errors.New("error reading slack history")
		return
	}

	return h.Messages, nil
}

func SendMessage(c config.Config, message string) error {
	data := url.Values{}
	data.Set("token", c.Token)
	data.Set("channel", c.Channel)
	data.Set("text", message)
	data.Set("username", c.Username)
	data.Set("as_user", "1")

	_, err := http.PostForm("https://slack.com/api/chat.postMessage", data)
	if err != nil {
		return err
	}
	return nil
}
