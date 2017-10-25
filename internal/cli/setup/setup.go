package setup

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/history"
)

type command struct {
	conf config.Config

	flag *flag.FlagSet
}

func parseCommand() cli.Command {
	f := flag.NewFlagSet("setup", flag.ExitOnError)

	tflag := f.String("t", "", "slack API token")
	cflag := f.String("c", "", "channel, private group, or IM channel to send message to")
	uflag := f.String("u", "", "your username in that channel")

	f.Parse(os.Args[2:])

	return &command{config.Config{*cflag, *tflag, *uflag, nil}, f}
}

func (c *command) Run() {
	if c.flag.NFlag() < 3 {
		c.Usage()
	}

	users, err := api.GetChatUsers(c.conf)
	if err != nil {
		log.Fatalf("error getting chat users: %s", err.Error())
	}
	c.conf.Users = users

	if err := config.Write(c.conf); err != nil {
		log.Fatalf("error saving config: %s", err.Error())
	}

	history.Clear()
}

func (c *command) Usage() {
	fmt.Printf(`Usage: %s setup -t=<slack-token> -c=<channel-id> -u=<channel-username>
	
<slack-token>      - slack API token
<channel-id>       - IM channel id to send message to
<channel-username> - your username in that chat
`, os.Args[0])
	os.Exit(2)
}

func init() {
	cli.RegisterCommand("setup", parseCommand)
}
