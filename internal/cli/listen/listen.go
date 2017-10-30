// Package listen is a command that starts background channel listening process (slkd)
package listen

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/blinkloader/slk/internal/cli"
	"github.com/blinkloader/slk/log"
)

type command struct{}

func initCommand() cli.Command {
	if len(os.Args) != 2 {
		usage()
	}

	return &command{}
}

func (c command) Run() {
	stopDaemon()
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

func (c command) Usage() {
	usage()
}

func usage() {
	fmt.Printf(`Usage: %s listen

Starts "slkd" background process. That process does
simple job - checks if there is any new messages
that you didn't see.
`, os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("listen", initCommand)
}
