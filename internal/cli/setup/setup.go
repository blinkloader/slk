// setup command sets slk configuration
// it must be run before any other commands
package setup

import (
	"flag"
	"fmt"
	"os"

	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/log"
)

type command struct {
	conf config.Config

	flag *flag.FlagSet
}

func parseCommand() cli.Command {
	if len(os.Args[2:]) == 0 {
		os.Exit(0)
	}

	conf, _ := config.Read()

	f := flag.NewFlagSet("setup", flag.ExitOnError)

	tflag := f.String("t", conf.Token, "slack API token")
	cflag := f.String("c", conf.Channel, "channel, private group, or IM channel to send message to")
	uflag := f.String("u", conf.Username, "your username")

	f.Parse(os.Args[2:])

	if len(f.Args()) != 0 {
		usage()
	}

	if *tflag == "" || *cflag == "" || *uflag == "" {
		usage()
	}

	return &command{config.Config{"", *cflag, *tflag, *uflag, nil, conf.ChannelTs}, f}
}

func (c *command) Run() {
	channelID, err := api.GetChannelID(c.conf.Token, c.conf.ChannelName)
	if err != nil {
		log.Fatal(err)
	}

	c.conf.Channel = channelID

	users, err := api.GetChanUsers(c.conf)
	if err != nil {
		log.Fatalf("error getting channel users: %s", err)
	}
	c.conf.Users = users

	if err := config.Write(c.conf); err != nil {
		log.Fatalf("error saving config: %s", err)
	}
}

func (c *command) Usage() {
	usage()
}

func usage() {
	fmt.Printf(`Usage: %s setup <options>

Set up configuration. You need to run this command before you can use "slk".
Options -t, -c, -u are required if you run "slk setup" for the first time.
Next time you run "slk setup" all flags are optional (it takes default values from $HOME/.slk)

Options:
  -t  -  slack API token
  -c  -  channel, private group, or IM channel to send message to
  -u  -  your username
`, os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("setup", parseCommand)
}
