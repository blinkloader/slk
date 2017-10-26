package message

import (
	"bytes"

	"github.com/yarikbratashchuk/slk/internal/api"
)

func RemoveURefs(messages []*api.Message) {
	for i, m := range messages {
		messages[i].Text = removeUserRef(m.Text)
	}
}

func removeUserRef(message string) string {
	buf := new(bytes.Buffer)
	l := len(message)
charloop:
	for i := 0; i < l; i++ {
		if message[i] == '<' && message[i+1] == '@' {
			for j := i + 2; j < l; j++ {
				if message[j] == '>' && j < l-1 && message[j+1] == ' ' {
					i = j + 1
					continue charloop
				}
			}
		} else {
			buf.WriteByte(message[i])
		}
	}
	return buf.String()
}
