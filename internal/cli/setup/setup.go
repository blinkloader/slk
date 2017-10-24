package setup

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	env "github.com/segmentio/go-env"
	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
)

type command struct {
	flag         *flag.FlagSet
	tflag, cflag *string
}

func parseCommand() cli.Command {
	if len(os.Args) < 2 {
		return &command{}
	}

	f := flag.NewFlagSet("setup", flag.ExitOnError)

	tflag := f.String("t", "", "slack API token")
	cflag := f.String("c", "", "channel, private group, or IM channel to send message to")

	f.Parse(os.Args[2:])

	return &command{f, tflag, cflag}
}

func (c *command) Run() {
	if c.flag.NFlag() < 2 {
		c.Usage()
	}

	// TODO: get users here
	conf := config.Config{*c.cflag, *c.tflag, nil}

	log.SetPrefix("setup: ")

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(conf); err != nil {
		log.Fatalf("error writing $HOME/.slk config file: %s", err.Error())
	}

	if err := ioutil.WriteFile(env.MustGet("HOME")+"/.slk", buf.Bytes(), 0755); err != nil {
		log.Fatalf("error writing $HOME/.slk config file: %s", err.Error())
	}
}

func (c *command) Usage() {
	fmt.Printf(`Usage: %s setup -t=<TOKEN> -c=<channel>
	
<TOKEN>   - slack API token
<channel> - channel, private group, or IM channel to send message to
`, os.Args[0])
	os.Exit(2)
}

func init() {
	cli.RegisterCommand("setup", parseCommand)
}
