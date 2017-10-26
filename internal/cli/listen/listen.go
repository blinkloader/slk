package listen

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/log"
)

type command struct {
	conf config.Config

	tflag bool
}

func initCommand() cli.Command {
	f := flag.NewFlagSet("listen", flag.ExitOnError)
	tflag := f.Bool("t", false, "terminates chat listening")
	f.Parse(os.Args[2:])

	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	return &command{conf, *tflag}
}

func (c *command) Run() {
	stopDaemon()
	if c.tflag {
		return
	}
	startDaemon()
}

func startDaemon() {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}
	cmd := exec.Command("slkd")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = cwd
	if err := cmd.Start(); err != nil {
		log.Fatalf("can't start chat listening process: %s", err)
	}

	cmd.Process.Release()
}

func stopDaemon() {
	cmd := exec.Command("killall", "slkd")
	_ = cmd.Run()
}

func (c *command) Usage() {
	fmt.Printf(`Usage: %s listen <options>

Options:
  -t  -  used to terminate chat listening process
`, os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("listen", initCommand)
}
