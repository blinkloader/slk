// Package write is a command that writes message to channel
package write

import (
	"fmt"
	"os"

	"github.com/blinkloader/slk/internal/api"
	"github.com/blinkloader/slk/internal/cli"
	"github.com/blinkloader/slk/internal/config"
	"github.com/blinkloader/slk/log"
)

type command struct {
	conf *config.Config
	api  api.Client

	message string
}

func initCommand() cli.Command {
	if len(os.Args) != 3 {
		usage()
	}

	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	return &command{conf, api.New(conf), os.Args[2]}
}

func (c *command) Run() {
	if err := c.api.SendMessage(c.message); err != nil {
		log.Fatalf("error sending message: %s", err)
	}
}

func (c *command) Usage() {
	usage()
}

func usage() {
	fmt.Printf(`Usage: %s write <message>

Writes message to the channel (you are currently on).
Configuration is stored at $HOME/.slk.
`, os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("write", initCommand)
}
