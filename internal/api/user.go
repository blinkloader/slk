package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/yarikbratashchuk/slk/errors"
)

type userlist struct {
	Members []*User `json:"members"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ChanUsers returns all users that are members of the channel
// if form map[userID]username
func (c client) ChanUsers() (users map[string]string, err error) {
	data := url.Values{}
	data.Set("token", c.conf.Token)
	data.Set("channel", c.conf.Channel)

	res, err := http.PostForm("https://slack.com/api/conversations.members", data)
	if err != nil {
		return
	}

	var resParsed struct {
		Members []string `json:"members"`
	}
	if err = json.NewDecoder(res.Body).Decode(&resParsed); err != nil {
		return
	}
	defer res.Body.Close()

	members, err := c.UserList()
	if err != nil {
		return
	}

	users = make(map[string]string)
	for _, m := range resParsed.Members {
		for _, u := range members {
			if u.ID == m {
				users[m] = u.Name
			}
		}
	}

	return
}

var errNoSuchUser = errors.New("can't find user with such name")

// GetUserID returns user id by given user name
func (c client) UserID(name string) (string, error) {
	members, err := c.UserList()
	if err != nil {
		return "", err
	}

	for _, u := range members {
		if u.Name == name {
			return u.ID, nil
		}
	}
	return "", errNoSuchUser
}

// UserList returns all users
func (c client) UserList() ([]*User, error) {
	data := url.Values{}
	data.Set("token", c.conf.Token)
	res, err := http.PostForm("https://slack.com/api/users.list", data)
	if err != nil {
		return []*User{}, err
	}
	defer res.Body.Close()

	var list userlist
	if err = json.NewDecoder(res.Body).Decode(&list); err != nil {
		return []*User{}, err
	}
	return list.Members, nil
}
