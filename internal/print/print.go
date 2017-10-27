// Package print holds printing primitives.
// All slack message printing happens via this package
package print

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/yarikbratashchuk/slk/internal/api"
)

// Chat prints channel messages.
// like this:
//   yarik: hello max
//     max: hello yarik
func Chat(username string, users map[string]string, messages []*api.Message) {
	chat(username, users, messages, false)
}

// ListenChat prints channel messages from background process (almost similar output).
func ListenChat(username string, users map[string]string, messages []*api.Message) {
	chat(username, users, messages, true)
}

func chat(username string, users map[string]string, messages []*api.Message, fromDaemon bool) {
	if len(messages) == 0 {
		return
	}

	uw := getUserWidth(users, messages)

	buf := new(bytes.Buffer)
	var my int

	buf.WriteString("\n")
	if fromDaemon {
		buf.WriteString("\n")
		buf.WriteString(color.New(color.FgGreen, color.Bold).Sprintf("  " + strings.Repeat(" ", uw-5) + "Slack:\n"))
	}
	messagesBuf, my := Messages(uw, username, users, messages)
	if fromDaemon && len(messages) == my {
		return
	}
	buf.Write(messagesBuf)
	buf.WriteString("\n")

	fmt.Print(buf.String())
}

func Messages(uw int, username string, users map[string]string, messages []*api.Message) ([]byte, int) {
	buf := new(bytes.Buffer)
	var my int
	for _, m := range messages {
		if m.Text == "" {
			continue
		}
		isme := users[m.User] == username
		if isme {
			my++
		}
		buf.WriteString("  ")
		buf.WriteString(Message(uw, users[m.User], m.Text, isme))
	}
	return buf.Bytes(), my
}

func Message(uw int, user, text string, me bool) string {
	space := strings.Repeat(" ", uw-len(user))
	u := space + user
	if me {
		u = color.CyanString(u)
	}

	text = strings.Replace(text, "\n\n", "\n", -1) // my secret trick
	text = strings.Replace(text, "\n", "\n    "+strings.Repeat(" ", uw), -1)

	return fmt.Sprintf("%s: %s\n", u, text)
}

func getUserWidth(users map[string]string, messages []*api.Message) (width int) {
	for _, m := range messages {
		if ul := len(users[m.User]); ul > width {
			width = ul
		}
	}
	return width
}
