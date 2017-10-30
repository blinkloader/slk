// Package ignore is a command that kills background process (slkd)
package ignore

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/blinkloader/slk/internal/cli"
)

type command struct{}

func initCommand() cli.Command {
	if len(os.Args) != 2 {
		usage()
	}
	return command{}
}

func (c command) Run() {
	stopDaemon()
}

func stopDaemon() {
	cmd := exec.Command("killall", "slkd")
	_ = cmd.Run()
}

func (c command) Usage() {
	usage()
}

func usage() {
	fmt.Printf(`Usage: %s ignore

Kills "slkd" background process. That process does
simple job - checks if there is any new messages
that you didn't see.
`, os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("ignore", initCommand)
}
