package shquot

import (
	"testing"
)

func TestDockerfile(t *testing.T) {
	tests := []qTest{
		{
			[]string{},
			`[]`,
		},
		{
			[]string{"tar", "zxvf", "foo.tar"},
			`["tar","zxvf","foo.tar"]`,
		},
		{
			[]string{"echo", `hello, "world"`},
			`["echo","hello, \"world\""]`,
		},
	}

	runTests(t, tests, Dockerfile)
}
