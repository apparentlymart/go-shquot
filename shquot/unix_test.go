package shquot

import (
	"fmt"
	"testing"
)

func TestUnixTerminal(t *testing.T) {
	tests := []qTest{
		{
			[]string{},
			"",
		},
		{
			[]string{"echo", "\x1b[0;0HHello, world!"},
			"echo '\x14\x1b[0;0HHello, world!'",
		},
		{
			[]string{"echo", "\x14\x1b[0;0HHello, world!"},
			"echo '\x14\x14\x14\x1b[0;0HHello, world!'",
		},
	}

	// We don't use runTests for this one because we want to quote the
	// got/want in the output to ensure the control codes don't get interpreted
	// by the terminal where the tests are running.
	q := UnixTerminal(POSIXShell)
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s", test.cmdline), func(t *testing.T) {
			got := q(test.cmdline)
			want := test.want
			if got != want {
				t.Errorf("wrong result\ninput: %#v\ngot:   %q\nwant:  %q", test.cmdline, got, want)
			}
		})
	}
}
