package read

import (
	"fmt"
	"os"

	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/log"
	"github.com/yarikbratashchuk/slk/internal/message"
	"github.com/yarikbratashchuk/slk/internal/print"
)

type command struct {
	conf config.Config
}

func initCommand() cli.Command {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	return &command{conf}
}

func (c *command) Run() {
	hist, err := api.GetChannelHistory(c.conf)
	if err != nil {
		log.Fatalf("error getting chat history: %s", err)
	}

	message.RemoveURefs(hist)

	print.Chat(c.conf.Username, c.conf.Users, hist)
}

func (c *command) Usage() {
	fmt.Printf("Usage: %s read\n", os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("read", initCommand)
}
