// Package message holds logic related to []*api.Message data structure
package message

import (
	"bytes"
	"strconv"

	"github.com/yarikbratashchuk/slk/internal/api"
)

// TsFilterNewer returns messages that you didn't read yet
// its work based on comparing message timestamps.
func TsFilterNewer(ts string, messages []*api.Message) (filtered []*api.Message) {
	if ts == "" {
		return []*api.Message{}
	}

	t, _ := strconv.ParseFloat(ts, 64)
	for _, m := range messages {
		mt, _ := strconv.ParseFloat(m.Ts, 64)
		if mt > t {
			filtered = append(filtered, m)
		}
	}
	return filtered
}

// RemoveURefs removes <@USERID> references from every message text,
// which are useless for terminal environment
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

// FormatLines tries to make each message line to have around 100 characters
func FormatLines(messages []*api.Message) {
	for i, m := range messages {
		messages[i].Text = formatLines(m.Text)
	}
}

func formatLines(message string) string {
	buf := new(bytes.Buffer)
	l := len(message)
	lineEl := 0
	addN := false
	for i := 0; i < l; i++ {
		if message[i] == ' ' && addN {
			buf.WriteString("\n")
			lineEl = 0
			addN = false
			continue
		}
		if message[i] == '\n' {
			lineEl = 0
		}
		if lineEl == 120 {
			addN = true
		}
		buf.WriteByte(message[i])
		lineEl++
	}
	return buf.String()
}

// ReverseOrder does what you expect
func ReverseOrder(messages []*api.Message) []*api.Message {
	r := make([]*api.Message, len(messages))
	for i := range messages {
		r[i] = messages[len(messages)-1-i]
	}
	return r
}
