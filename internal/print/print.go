package print

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/config"
)

func Chat(username string, users map[string]config.User, messages []*api.Message) {
	uv := getUserWidth(users, messages)

	buf := new(bytes.Buffer)

	buf.WriteString("\n")
	for _, m := range messages {
		buf.WriteString("  ")
		buf.WriteString(Message(uv, users[m.User].Name, m.Text, users[m.User].Name == username))
	}
	buf.WriteString("\n")

	fmt.Print(buf.String())
}

func ListenChat(username string, users map[string]config.User, messages []*api.Message) {
	if len(messages) == 0 {
		return
	}

	uv := getUserWidth(users, messages)

	buf := new(bytes.Buffer)
	var my int

	buf.WriteString("\n\n")
	buf.WriteString(color.New(color.FgGreen, color.Bold).Sprintf("  Slack:\n"))
	for _, m := range messages {
		isme := users[m.User].Name == username
		if isme {
			my++
		}
		buf.WriteString("  ")
		buf.WriteString(Message(uv, users[m.User].Name, m.Text, isme))
	}
	if len(messages) == my {
		return
	}
	buf.WriteString("\n")

	fmt.Print(buf.String())
}

func Message(uv int, user, text string, me bool) string {
	space := strings.Repeat(" ", uv-len(user))
	u := space + user
	if me {
		u = color.CyanString(u)
	}

	text = strings.Replace(text, "\n", "\n    "+strings.Repeat(" ", uv), -1)

	return fmt.Sprintf("%s: %s\n", u, text)
}

func getUserWidth(users map[string]config.User, messages []*api.Message) (width int) {
	for _, m := range messages {
		if ul := len(users[m.User].Name); ul > width {
			width = ul
		}
	}
	return width
}
