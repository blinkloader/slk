package write

import (
	"flag"
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

	message string
}

func initCommand() cli.Command {
	if len(os.Args) < 2 {
		return &command{}
	}

	f := flag.NewFlagSet("write", flag.ExitOnError)
	mflag := f.String("m", "", "text message")
	f.Parse(os.Args[2:])

	var conf config.Config
	log.SetPrefix("write: ")
	if _, err := toml.DecodeFile(env.MustGet("HOME")+"/.slk", &conf); err != nil {
		log.Fatalf("error reading $HOME/.slk config file: %s", err.Error())
	}

	return &command{conf.Channel, conf.Token, conf.Users, *mflag}
}

func (c *command) Run() {
	data := url.Values{}
	data.Set("token", c.token)
	data.Set("channel", c.channel)
	data.Set("text", c.message)
	data.Set("as_user", "1")
	data.Set("username", "yarik")

	_, err := http.PostForm("https://slack.com/api/chat.postMessage", data)
	if err != nil {
		log.Fatalf("error sending message: %s", err.Error())
	}
}

func (c *command) Usage() {
	fmt.Printf("Usage: %s write -m <message>\n", os.Args[0])
	os.Exit(2)
}

func init() {
	cli.RegisterCommand("write", initCommand)
}
