// Package api holds interface to slack api.
package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/errors"
)

var (
	errGetHistory  = errors.New("error getting channel history")
	errPostMessage = errors.New("error posting message")

	errGetChannelID  = errors.New("error getting channel-id")
	errNoSuchChannel = errors.New("no such channel")
)

type history struct {
	Ok       bool       `json:"ok"`
	Messages []*Message `json:"messages"`
}

type Message struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
	Ts   string `json:"ts"`
}

// GetChannelHistory returns 10 last messages in the channel
func GetChannelHistory(conf config.Config) ([]*Message, error) {
	data := url.Values{}
	data.Set("token", conf.Token)
	data.Set("channel", conf.Channel)
	data.Set("limit", "10")

	res, err := http.PostForm("https://slack.com/api/conversations.history", data)
	if err != nil {
		return []*Message{}, errors.Wrap(errGetHistory, err)
	}

	var h history
	if err = json.NewDecoder(res.Body).Decode(&h); err != nil {
		return []*Message{}, errors.Wrap(errGetHistory, err)
	}
	defer res.Body.Close()

	if !h.Ok {
		return []*Message{}, errGetHistory
	}

	return h.Messages, nil
}

// SendMessage sends message to channel
func SendMessage(c config.Config, message string) error {
	data := url.Values{}
	data.Set("token", c.Token)
	data.Set("channel", c.Channel)
	data.Set("text", message)
	data.Set("username", c.Username)
	data.Set("as_user", "1")

	_, err := http.PostForm("https://slack.com/api/chat.postMessage", data)
	if err != nil {
		return errors.Wrap(errPostMessage, err)
	}
	return nil
}
