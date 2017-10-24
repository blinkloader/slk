package config

type Config struct {
	Channel  string `toml:"channel"`
	Token    string `toml:"token"`
	Username string `toml:"username"`

	Users map[string]User `toml:"users"`
}

type User struct {
	Name string `toml:"name"`
}
