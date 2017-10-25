package config

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
	env "github.com/segmentio/go-env"
)

var (
	configFile    = env.MustGet("HOME") + "/.slk"
	daemonPIDFile = env.MustGet("HOME") + "/.slkd"
)

type Config struct {
	Channel  string `toml:"channel"`
	Token    string `toml:"token"`
	Username string `toml:"username"`

	Users map[string]User `toml:"users"`
}

type User struct {
	Name string `toml:"name"`
}

func Read() Config {
	var conf Config
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		log.Fatalf("error reading $HOME/.slk config file: %s", err.Error())
	}
	return conf
}

func Write(conf Config) error {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(conf); err != nil {
		return err
	}

	if err := ioutil.WriteFile(configFile, buf.Bytes(), 0755); err != nil {
		return err
	}
	return nil
}
