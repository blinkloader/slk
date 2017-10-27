// on command prints name of the current channel
package on

import (
	"fmt"
	"log"
	"os"

	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
)

type command struct {
	conf config.Config
}

func initCommand() cli.Command {
	if len(os.Args) != 2 {
		usage()
	}

	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	return &command{conf}
}

func (c command) Run() {
	fmt.Println(c.conf.ChannelName)
}

func (c command) Usage() {
	usage()
}

func usage() {
	fmt.Printf(`Usage: %s on

Returns name of the current channel
`, os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("on", initCommand)
	cli.RegisterCommand("at", initCommand)
	cli.RegisterCommand("in", initCommand)
}
