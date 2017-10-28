// Package config holds all stuff related to slk configuration ($HOME/.slk)
package config

import (
	"bytes"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	env "github.com/segmentio/go-env"
	"github.com/yarikbratashchuk/slk/errors"
)

var (
	configFile = env.MustGet("HOME") + "/.slk"

	errReadConfig  = errors.New("reading $HOME/.slk")
	errWriteConfig = errors.New("writing $HOME/.slk")
)

// Config holds slk configuration
type Config struct {
	Channel     string `toml:"channel"`
	ChannelName string `toml:"channel-name"`
	Token       string `toml:"token"`
	Username    string `toml:"username"`

	Users map[string]string `toml:"users"`

	ChannelTs map[string]string `toml:"channel-ts"`
}

// Read reads config from file ($HOME/.slk)
func (c *Config) Read() error {
	if _, err := toml.DecodeFile(configFile, c); err != nil {
		return errors.Wrap(errReadConfig, err)
	}
	if c.ChannelTs == nil {
		c.ChannelTs = make(map[string]string)
	}
	return nil
}

// Write writes config to file ($HOME/.slk)
func (c *Config) Write() error {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		return err
	}

	if err := ioutil.WriteFile(configFile, buf.Bytes(), 0755); err != nil {
		return errors.Wrap(errWriteConfig, err)
	}
	return nil
}

var conf = &Config{}

// Read reads config from file ($HOME/.slk)
func Read() (*Config, error) {
	if err := conf.Read(); err != nil {
		return &Config{}, err
	}
	return conf, nil
}
