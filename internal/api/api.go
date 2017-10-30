// Package api holds interface to slack api.
package api

import "github.com/yarikbratashchuk/slk/internal/config"

// Client is an api interface
type Client interface {
	// ChannelHistory returns last messages in the channel
	ChannelHistory() ([]*Message, error)

	// ChannelID returns channel id by given name
	ChannelID(string) (string, error)
	// GroupID returns public channel or private group id by given name
	GroupID(string) (string, error)
	// ImChatID returns IM chat id by given user name
	ImChatID(string) (string, error)

	// ChanUsers returns all users that are members of the channel
	// if form map[userID]username
	ChanUsers() (map[string]string, error)

	// UserID returns user id by given user name
	UserID(string) (string, error)
	// UserList returns all users
	UserList() ([]*User, error)

	// SendMessage sends message to channel
	SendMessage(string) error

	// NumMessages sets number of last messages being returned
	NumMessages(int) Client
}

// client implements Client
type client struct {
	conf *config.Config

	msglimit int
}

// New returns new api client
func New(conf *config.Config) Client {
	return client{conf, 10}
}

// NumMessages sets messages limit
func (c client) NumMessages(l int) Client {
	c.msglimit = l
	return c
}
