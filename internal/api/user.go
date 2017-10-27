package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/errors"
)

type userList struct {
	Members []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"members"`
}

// GetChanUsers returns all users that are members of the channel
// if form map[userID]username
func GetChanUsers(conf config.Config) (users map[string]string, err error) {
	data := url.Values{}
	data.Set("token", conf.Token)
	data.Set("channel", conf.Channel)

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

	userlist, err := getUserList(conf.Token)
	if err != nil {
		return
	}

	users = make(map[string]string)
	for _, m := range resParsed.Members {
		for _, u := range userlist.Members {
			if u.ID == m {
				users[m] = u.Name
			}
		}
	}

	return
}

var errNoSuchUser = errors.New("can't find user with such name")

// GetUserID returns user id by given user name
func GetUserID(token, name string) (string, error) {
	ulist, err := getUserList(token)
	if err != nil {
		return "", err
	}

	for _, u := range ulist.Members {
		if u.Name == name {
			return u.ID, nil
		}
	}
	return "", errNoSuchUser
}

func getUserList(token string) (list userList, err error) {
	data := url.Values{}
	data.Set("token", token)
	res, err := http.PostForm("https://slack.com/api/users.list", data)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&list); err != nil {
		return
	}
	return
}
