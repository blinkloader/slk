package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"

	"github.com/blinkloader/slk/errors"
)

type chanList struct {
	Ok bool `json:"ok"`

	Channels []*Channel `json:"channels"`
	Groups   []*Channel `json:"groups"`
	Ims      []*Channel `json:"ims"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	User string `json:"user"`
}

// ChannelID returns channel id by given name
func (c client) ChannelID(name string) (id string, err error) {
	switch name[0] {
	case '@': // direct messages
		id, err = c.ImChatID(name[1:])
	default: // public channels and private groups
		id, err = c.GroupID(name)
	}
	if err != nil {
		return "", errors.Wrap(errGetChannelID, err)
	}
	return
}

// ImChatID returns IM chat id by given user name
func (c client) ImChatID(name string) (string, error) {
	l, err := getChannelList(c.conf.Token, "https://slack.com/api/im.list")
	if err != nil {
		return "", err
	}

	userID, err := c.UserID(name)
	if err != nil {
		return "", err
	}

	id, err := getChannelID("User", userID, l.Ims)
	if err != nil {
		return "", err
	}

	return id, nil
}

// GroupID returns public channel or private group id by given name
func (c client) GroupID(name string) (string, error) {
	id, err := getPubChannelID(c.conf.Token, name)
	if err == nil && id != "" {
		return id, nil
	}
	return getPrivGroupID(c.conf.Token, name)
}

// getPubChannelID returns public channel id by given name
func getPubChannelID(token, name string) (string, error) {
	l, err := getChannelList(token, "https://slack.com/api/channels.list")
	if err != nil {
		return "", err
	}

	id, err := getChannelID("Name", name, l.Channels)
	if err != nil {
		return "", err
	}

	return id, nil
}

// getPrivGroupID returns private group id by given name
func getPrivGroupID(token, name string) (string, error) {
	l, err := getChannelList(token, "https://slack.com/api/groups.list")
	if err != nil {
		return "", err
	}

	id, err := getChannelID("Name", name, l.Groups)
	if err != nil {
		return "", err
	}

	return id, nil
}

func getChannelList(token, u string) (chanList, error) {
	data := url.Values{}
	data.Set("token", token)
	res, err := http.PostForm(u, data)
	if err != nil {
		return chanList{}, err
	}

	var l chanList
	if err := json.NewDecoder(res.Body).Decode(&l); err != nil {
		return chanList{}, err
	}

	return l, nil
}

func getChannelID(fname, name string, list []*Channel) (id string, err error) {
	defer func() {
		if e := recover(); err != nil {
			if er, ok := e.(error); ok {
				err = er
			}
			return
		}
	}()
	for _, c := range list {
		n := reflect.ValueOf(c).Elem().FieldByName(fname).String()
		if n == name {
			return c.ID, nil
		}
	}
	return "", errNoSuchChannel
}
