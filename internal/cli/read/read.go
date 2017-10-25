package read

import (
	"fmt"
	"log"
	"os"

	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/print"
)

type command struct {
	conf config.Config
}

func initCommand() cli.Command {
	return &command{config.Read()}
}

func (c *command) Run() {
	hist, err := api.GetChatHistory(c.conf)
	if err != nil {
		log.Fatalf("error getting chat history: %s", err.Error())
	}

	print.Chat(c.conf.Username, c.conf.Users, hist)
}

func (c *command) Usage() {
	fmt.Printf("Usage: %s read\n", os.Args[0])
	os.Exit(2)
}

func init() {
	cli.RegisterCommand("read", initCommand)
}
