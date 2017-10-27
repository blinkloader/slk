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

func TestTsFilterNewer(t *testing.T) {
	cases := []struct {
		test   string
		ts     string
		before []*api.Message
		after  []*api.Message
	}{{
		"simple case",
		"1234.000",
		[]*api.Message{
			&api.Message{Text: "1", Ts: "12345.000"},
			&api.Message{Text: "2", Ts: "125.000"},
			&api.Message{Text: "3", Ts: "15.000"},
			&api.Message{Text: "4", Ts: "12.000"},
			&api.Message{Text: "5", Ts: "1.000"},
		},
		[]*api.Message{
			&api.Message{Text: "1", Ts: "12345.000"},
		},
	}}
	for _, c := range cases {
		t.Run(c.test, func(t *testing.T) {
			filtered := TsFilterNewer(c.ts, c.before)
			for i, m := range filtered {
				if m.Ts != c.after[i].Ts {
					t.Error("didn't work")
				}
			}
			if len(filtered) != len(c.after) {
				t.Error("didn't work")
			}
		})
	}
}

func TestFormatLines(t *testing.T) {
	cases := []struct {
		test   string
		before []*api.Message
		after  []*api.Message
	}{{
		"simple case",
		[]*api.Message{
			&api.Message{
				Text: "random random random random random random random random random random random random random random random random random random random random random ",
			},
		},
		[]*api.Message{
			&api.Message{
				Text: "random random random random random random random random random random random random random random random random random random\nrandom random random ",
			},
		},
	}}
	for _, c := range cases {
		t.Run(c.test, func(t *testing.T) {
			FormatLines(c.before)
			for i, m := range c.before {
				if m.Text != c.after[i].Text {
					t.Error("didn't work")
				}
			}
		})
	}
}
