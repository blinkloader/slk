package main

import (
	"fmt"
	"os"

	"github.com/yarikbratashchuk/slk/internal/cli"

	_ "github.com/yarikbratashchuk/slk/internal/cli/listen"
	_ "github.com/yarikbratashchuk/slk/internal/cli/read"
	_ "github.com/yarikbratashchuk/slk/internal/cli/setup"
	_ "github.com/yarikbratashchuk/slk/internal/cli/write"
)

func main() {
	if len(os.Args) < 2 {
		usage(os.Args[0])
	}

	commands := cli.InitCommands()

	if len(os.Args) >= 2 && os.Args[1] == "help" {
		if len(os.Args) > 2 {
			if commands[os.Args[2]] != nil {
				commands[os.Args[2]]().Usage()
			}
		}
		usage(os.Args[0])
	}

	commands[os.Args[1]]().Run()
}

func usage(command string) {
	fmt.Printf("Usage: %s <command> <options>\n", command)
	os.Exit(2)
}
