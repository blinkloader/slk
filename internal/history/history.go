package history

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	env "github.com/segmentio/go-env"
	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/errors"
)

var (
	historyFile = env.MustGet("HOME") + "/.slk_history"

	errReadingHistory  = errors.New("error reading $HOME/.slk_history")
	errWritingHistory  = errors.New("error writing $HOME/.slk_history")
	errRemovingHistory = errors.New("error removing $HOME/.slk_history")
)

type History struct {
	Messages []*api.Message `toml:"messages"`
}

func Read() ([]*api.Message, error) {
	var hist History
	if _, err := toml.DecodeFile(historyFile, &hist); err != nil {
		return nil, errors.Wrap(errReadingHistory, err)
	}
	return hist.Messages, nil
}

func Update(old, dif []*api.Message) error {
	dif = append(dif, old...)

	if len(dif) > 10 {
		dif = dif[:10]
	}

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(History{dif}); err != nil {
		return errors.Wrap(errWritingHistory, err)
	}

	if err := ioutil.WriteFile(historyFile, buf.Bytes(), 0755); err != nil {
		return errors.Wrap(errWritingHistory, err)
	}
	return nil
}

func Diff(old, new []*api.Message) (diff []*api.Message) {
newloop:
	for _, m := range new {
		for _, n := range old {
			if n.Text == m.Text && n.Ts == m.Ts {
				continue newloop
			}
		}
		diff = append(diff, m)
	}
	return diff
}

func Clear() error {
	if err := os.Remove(historyFile); err != nil {
		return errors.Wrap(errRemovingHistory, err)
	}
	return nil
}
