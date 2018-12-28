package shquot

import (
	"testing"
)

func TestPOSIXShell(t *testing.T) {
	tests := []qTest{
		{
			[]string{},
			``,
		},
		{
			[]string{`echo`, ``},
			`echo ''`,
		},
		{
			[]string{`ls`},
			`ls`,
		},
		{
			[]string{`ls`, `-lah`},
			`ls -lah`,
		},
		{
			[]string{`echo`, `hello world`},
			`echo 'hello world'`,
		},
		{
			[]string{`echo`, `"hello world"`},
			`echo '"hello world"'`,
		},
		{
			[]string{`echo`, `I'm alive!`},
			`echo 'I'\''m alive!'`,
		},
		{
			[]string{`echo`, `Hello!`},
			`echo Hello\!`,
		},
		{
			[]string{`echo`, `hello`, `>foo.txt`},
			`echo hello \>foo.txt`,
		},
		{
			[]string{`echo`, `hello`, `>~/foo.txt`},
			`echo hello \>\~/foo.txt`,
		},
		{
			[]string{`cat`, `baz`, `|`, `grep`, `foo`},
			`cat baz \| grep foo`,
		},
		{
			[]string{`curl`, `https://example.com/?q=a&p=x`},
			`curl https://example.com/\?q\=a\&p\=x`,
		},
		{
			[]string{`sleep`, `50`, `&`},
			`sleep 50 \&`,
		},
		{
			[]string{`echo`, "hello\nworld\n"},
			"echo 'hello\nworld\n'",
		},
		{
			[]string{`bash`, `-c`, `foo | bar | baz`},
			`bash -c 'foo | bar | baz'`,
		},
	}

	runTests(t, tests, POSIXShell)
}
