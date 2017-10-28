// to command switches channel (you are currently on)
package to

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/yarikbratashchuk/slk/internal/cli"
	"github.com/yarikbratashchuk/slk/log"
)

type command struct{}

func initCommand() cli.Command {
	if len(os.Args) != 3 {
		usage()
	}
	return command{}
}

func (c command) Run() {
	cmd := exec.Command("slk", "setup", "-c", os.Args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("slk: %s", err)
	}
}

func (c command) Usage() {
	usage()
}

func usage() {
	fmt.Printf(`Usage: %s to <channel>

Used to switch between channels. Equivalent to "slk setup -c <channel>".
Use "slk on" to see what channel you are currently on.
`, os.Args[0])
	os.Exit(0)
}

func init() {
	cli.RegisterCommand("to", initCommand)
}
