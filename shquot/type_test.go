package shquot

import (
	"fmt"
	"testing"
)

type qTest struct {
	cmdline []string
	want    string
}

func runTests(t *testing.T, tests []qTest, q Q) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s", test.cmdline), func(t *testing.T) {
			got := q(test.cmdline)
			want := test.want
			if got != want {
				t.Errorf("wrong result\ninput: %#v\ngot:   %s\nwant:  %s", test.cmdline, got, want)
			}
		})
	}
}
