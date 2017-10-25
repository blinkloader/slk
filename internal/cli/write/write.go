package write

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
)

type command struct {
	conf config.Config

	message string
}

func initCommand() cli.Command {
	if len(os.Args) < 2 {
		return &command{}
	}

	f := flag.NewFlagSet("write", flag.ExitOnError)
	mflag := f.String("m", "", "text message")
	f.Parse(os.Args[2:])

	return &command{config.Read(), *mflag}
}

func (c *command) Run() {
	if err := api.SendMessage(c.conf, c.message); err != nil {
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
