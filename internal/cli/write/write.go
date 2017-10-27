// write command writes message to channel
package write

import (
	"fmt"
	"os"

	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/log"
)

type command struct {
	conf config.Config

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

	return &command{conf, os.Args[2]}
}

func (c *command) Run() {
	if err := api.SendMessage(c.conf, c.message); err != nil {
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
