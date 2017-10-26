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

func GetChanUsers(conf config.Config) (users map[string]config.User, err error) {
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

	users = make(map[string]config.User)
	for _, m := range resParsed.Members {
		for _, u := range userlist.Members {
			if u.ID == m {
				users[m] = config.User{u.Name}
			}
		}
	}

	return
}

var errNoSuchUser = errors.New("can't find user with such name")

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
