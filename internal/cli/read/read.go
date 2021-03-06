// Package read is a command that returns 10 last messages from channel
package read

import (
	"fmt"
	"os"

	"github.com/blinkloader/slk/internal/api"
	"github.com/blinkloader/slk/internal/cli"
	"github.com/blinkloader/slk/internal/config"
	"github.com/blinkloader/slk/internal/message"
	"github.com/blinkloader/slk/internal/out"
	"github.com/blinkloader/slk/log"
)

type command struct {
	conf *config.Config
	api  api.Client
}

func initCommand() cli.Command {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	return &command{conf, api.New(conf).NumMessages(10)}
}

func (c *command) Run() {
	hist, err := c.api.ChannelHistory()
	if err != nil {
		log.Fatalf("error getting chat history: %s", err)
	}

	hist = message.ReverseOrder(hist)

	message.RemoveURefs(hist)
	message.FormatLines(hist)

	out.PrintChat(c.conf.Username, c.conf.Users, hist)

	if len(hist) == 0 {
		return
	}

	c.conf.ChannelTs[c.conf.Channel] = hist[len(hist)-1].Ts

	if err := c.conf.Write(); err != nil {
		log.Fatal(err)
	}
}

func (c *command) Usage() {
	fmt.Printf(`Usage: %s read

Gets 10 last messages from channel (you are currently on).
Configuration is stored at $HOME/.slk.
`, os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("read", initCommand)
}
