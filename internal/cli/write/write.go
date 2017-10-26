package write

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

	message string
}

func initCommand() cli.Command {
	if len(os.Args) < 2 {
		return &command{}
	}

	f := flag.NewFlagSet("write", flag.ExitOnError)
	mflag := f.String("m", "", "text message")
	f.Parse(os.Args[2:])

	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	return &command{conf, *mflag}
}

func (c *command) Run() {
	if err := api.SendMessage(c.conf, c.message); err != nil {
		log.Fatalf("error sending message: %s", err)
	}
}

func (c *command) Usage() {
	fmt.Printf(`Usage: %s write <options>
	
Options:
  -m  -  message
`, os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("write", initCommand)
}
