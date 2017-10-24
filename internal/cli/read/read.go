package read

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
	env "github.com/segmentio/go-env"
	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
)

type command struct {
	channel, token string
	users          map[string]config.User
}

func initCommand() cli.Command {
	var conf config.Config

	if _, err := toml.DecodeFile(env.MustGet("HOME")+"/.slk", &conf); err != nil {
		log.Fatalf("error reading $HOME/.slk config file: %s", err.Error())
	}

	return &command{conf.Channel, conf.Token, conf.Users}
}

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

func (c *command) Run() {
	data := url.Values{}
	data.Set("token", c.token)
	data.Set("channel", c.channel)
	// TODO: track history in a .slk_history file
	data.Set("limit", "5")

	res, err := http.PostForm("https://slack.com/api/conversations.history", data)
	if err != nil {
		log.Fatalf("error fetching slack history: %s", err.Error())
	}

	var hist history
	if err := json.NewDecoder(res.Body).Decode(&hist); err != nil {
		log.Fatalf("error reading slack history: %s", err.Error())
	}
	defer res.Body.Close()

	if !hist.Ok {
		log.Fatalf("error reading slack history")
	}

	for _, m := range hist.Messages {
		fmt.Printf("%s: %s\n", c.users[m.User].Name, m.Text)
	}
}

func (c *command) Usage() {
	fmt.Printf("Usage: %s read\n", os.Args[0])
	os.Exit(2)
}

func init() {
	cli.RegisterCommand("read", initCommand)
}
