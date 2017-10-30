// Package out holds printing primitives.
// All slack message printing happens via this package
package out

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/blinkloader/slk/internal/api"
)

type Printer struct {
	buf *bytes.Buffer

	Out io.Writer
}

func (p *Printer) PrintChat(username string, users map[string]string, messages []*api.Message, fromDaemon bool) {
	if len(messages) == 0 {
		return
	}

	uw := getUserWidth(users, messages)

	var my int

	p.buf.WriteString("\n")
	if fromDaemon {
		p.buf.WriteString("\n")
		slkSpace := 0
		if s := uw - 5; s > 0 {
			slkSpace = s
		}
		p.buf.WriteString(color.New(color.FgGreen, color.Bold).Sprintf("  " + strings.Repeat(" ", slkSpace) + "Slack:\n"))
	}
	messagesBuf, my := writeMsgs(uw, username, users, messages)
	if fromDaemon && len(messages) == my {
		p.buf.Reset()
		return
	}
	p.buf.Write(messagesBuf)
	p.buf.WriteString("\n")

	p.Out.Write(p.buf.Bytes())
	p.buf.Reset()
}

var defPrinter = &Printer{
	new(bytes.Buffer),
	os.Stdout,
}

// PrintChat prints channel messages.
// like this:
//   yarik: hello max
//     max: hello yarik
func PrintChat(username string, users map[string]string, messages []*api.Message) {
	defPrinter.PrintChat(username, users, messages, false)
}

// PrintChatBg prints channel messages from background process (almost similar output).
func PrintChatBg(username string, users map[string]string, messages []*api.Message) {
	defPrinter.PrintChat(username, users, messages, true)
}

func writeMsgs(uw int, username string, users map[string]string, messages []*api.Message) ([]byte, int) {
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
		buf.WriteString(writeMsg(uw, users[m.User], m.Text, isme))
	}
	return buf.Bytes(), my
}

func writeMsg(uw int, user, text string, me bool) string {
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
