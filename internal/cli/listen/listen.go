package listen

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/internal/config"
)

type command struct {
	conf config.Config

	tflag bool
}

func initCommand() cli.Command {
	f := flag.NewFlagSet("listen", flag.ExitOnError)
	tflag := f.Bool("t", false, "terminates chat listening")
	f.Parse(os.Args[2:])
	return &command{config.Read(), *tflag}
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
		log.Fatalf("can't start chat listening process: %s", err.Error())
	}

	cmd.Process.Release()
}

func stopDaemon() {
	cmd := exec.Command("killall", "slkd")
	_ = cmd.Run()
}

func (c *command) Usage() {
	fmt.Printf("Usage: %s listen [-t]\n", os.Args[0])
	os.Exit(2)
}

func init() {
	cli.RegisterCommand("listen", initCommand)
}
