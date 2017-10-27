package message

import (
	"testing"

	"github.com/yarikbratashchuk/slk/internal/api"
)

func TestReverseOrder(t *testing.T) {
	cases := []struct {
		test   string
		before []*api.Message
		after  []*api.Message
	}{{
		"even number of elements",
		[]*api.Message{
			&api.Message{Text: "1"},
			&api.Message{Text: "2"},
			&api.Message{Text: "3"},
			&api.Message{Text: "4"},
		},
		[]*api.Message{
			&api.Message{Text: "4"},
			&api.Message{Text: "3"},
			&api.Message{Text: "2"},
			&api.Message{Text: "1"},
		},
	}, {
		"odd number of elements",
		[]*api.Message{
			&api.Message{Text: "1"},
			&api.Message{Text: "2"},
			&api.Message{Text: "3"},
		},
		[]*api.Message{
			&api.Message{Text: "3"},
			&api.Message{Text: "2"},
			&api.Message{Text: "1"},
		},
	}}
	for _, c := range cases {
		t.Run(c.test, func(t *testing.T) {
			rev := ReverseOrder(c.before)
			for i, m := range rev {
				if m.Text != c.after[i].Text {
					t.Error("didn't work")
				}
			}
		})
	}
}
