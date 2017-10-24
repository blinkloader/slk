package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/yarikbratashchuk/slk/internal/config"
)

type userList struct {
	Members []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"members"`
}

func GetChatUsers(conf config.Config) (users map[string]config.User, err error) {
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

	res, err = http.PostForm("https://slack.com/api/users.list", data)
	if err != nil {
		return
	}
	defer res.Body.Close()

	var userlist userList
	if err = json.NewDecoder(res.Body).Decode(&userlist); err != nil {
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
