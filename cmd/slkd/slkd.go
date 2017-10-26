package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/history"
	"github.com/yarikbratashchuk/slk/internal/message"
	"github.com/yarikbratashchuk/slk/internal/print"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-sigs
		os.Exit(0)
	}()

	for {
		newConf, err := config.Read()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if newConf.Channel != conf.Channel {
			_ = history.Clear()
			conf = newConf
		}

		fetched, err := api.GetChannelHistory(conf)
		if err != nil {
			fmt.Printf("\nslkd: %s", err)
		}

		var diff []*api.Message
		loaded, err := history.Read()
		if err == nil {
			diff = history.Diff(loaded, fetched)
		}

		if err = history.Update(loaded, diff); err != nil {
			fmt.Printf("\nslkd: %s", err)
		}

		message.RemoveURefs(diff)

		print.ListenChat(conf.Username, conf.Users, diff)

		time.Sleep(10 * time.Second)
	}

}
