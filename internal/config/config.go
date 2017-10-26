package config

import (
	"bytes"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	env "github.com/segmentio/go-env"
	"github.com/yarikbratashchuk/slk/internal/errors"
)

var (
	configFile    = env.MustGet("HOME") + "/.slk"
	daemonPIDFile = env.MustGet("HOME") + "/.slkd"

	errReadConfig  = errors.New("reading $HOME/.slk")
	errWriteConfig = errors.New("writing $HOME/.slk")
)

type Config struct {
	Channel     string `toml:"channel"`
	ChannelName string `toml:"channel-name"`
	Token       string `toml:"token"`
	Username    string `toml:"username"`

	Users map[string]User `toml:"users"`
}

type User struct {
	Name string `toml:"name"`
}

func Read() (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		return Config{}, errors.Wrap(errReadConfig, err)
	}
	return conf, nil
}

func Write(conf Config) error {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(conf); err != nil {
		return err
	}

	if err := ioutil.WriteFile(configFile, buf.Bytes(), 0755); err != nil {
		return errors.Wrap(errWriteConfig, err)
	}
	return nil
}

func UpdateProvided(old, new Config) (conf Config, err error) {
	conf = old
	if new.Channel != "" {
		conf.Channel = new.Channel
	}
	if new.ChannelName != "" {
		conf.ChannelName = new.ChannelName
	}
	if new.Token != "" {
		conf.Token = new.Token
	}
	if new.Username != "" {
		conf.Username = new.Username
	}
	return conf, nil
}
