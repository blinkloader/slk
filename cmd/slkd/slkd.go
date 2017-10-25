package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/history"
	"github.com/yarikbratashchuk/slk/internal/print"
)

func main() {
	conf := config.Read()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-sigs
		if err := config.RemoveDaemonPID(); err != nil {
			log.Println(err)
		}

		os.Exit(1)
	}()

	for {
		fetched, err := api.GetChatHistory(conf)
		if err != nil {
			log.Println(err)
			continue
		}

		var diff []*api.Message
		loaded, err := history.Read()
		if err == nil {
			diff = history.Diff(loaded, fetched)
		}

		err = history.Update(loaded, diff)

		print.ListenChat(conf.Username, conf.Users, diff)

		time.Sleep(10 * time.Second)
	}

}
