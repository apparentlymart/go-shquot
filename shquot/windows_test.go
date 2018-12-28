package shquot

import (
	"testing"
)

func TestWindowsArgv(t *testing.T) {
	tests := []qTest{
		{
			[]string{},
			"",
		},
		{
			[]string{"cmd.exe"},
			`"cmd.exe"`,
		},
		{
			[]string{`echo`, `hello`},
			`"echo" hello`,
		},
		{
			[]string{`echo`, ``},
			`"echo" ""`,
		},
		{
			[]string{`echo`, `hello world`},
			`"echo" "hello world"`,
		},
		{
			[]string{`echo`, `hello\ world`},
			`"echo" "hello\ world"`,
		},
		{
			[]string{`echo`, `hello world\`},
			`"echo" "hello world\\"`,
		},
		{
			[]string{`echo`, `hello world\\`},
			`"echo" "hello world\\\\"`,
		},
		{
			[]string{`echo`, `hello world\"`},
			`"echo" "hello world\\\""`,
		},
		{
			[]string{`echo`, `hello "world"`},
			`"echo" "hello \"world\""`,
		},
		{
			[]string{`"hello" world`, `hello "world"`},
			`"hello world" "hello \"world\""`,
		},
	}

	runTests(t, tests, WindowsArgv)
}

func TestWindowsCmdExe(t *testing.T) {
	tests := []qTest{
		{
			[]string{},
			"",
		},
		{
			[]string{"cmd.exe"},
			`^"cmd.exe^"`,
		},
		{
			[]string{`echo`, `hello`},
			`^"echo^" hello`,
		},
		{
			[]string{`echo`, `hello world`},
			`^"echo^" ^"hello world^"`,
		},
		{
			[]string{`echo`, `hello\ world`},
			`^"echo^" ^"hello\ world^"`,
		},
		{
			[]string{`echo`, `hello world\`},
			`^"echo^" ^"hello world\\^"`,
		},
		{
			[]string{`echo`, `hello world\\`},
			`^"echo^" ^"hello world\\\\^"`,
		},
		{
			[]string{`echo`, `hello world\"`},
			`^"echo^" ^"hello world\\\^"^"`,
		},
		{
			[]string{`echo`, `hello "world"`},
			`^"echo^" ^"hello \^"world\^"^"`,
		},
		{
			[]string{`type`, `foo`, `|`, `sort`},
			`^"type^" foo ^| sort`,
		},
		{
			[]string{`type`, `foo`, `&`, `sort`},
			`^"type^" foo ^& sort`,
		},
		{
			[]string{`hello`, `(world)`},
			`^"hello^" ^(world^)`,
		},
		{
			[]string{`echo`, `hello`, `>foo.txt`},
			`^"echo^" hello ^>foo.txt`,
		},
		{
			[]string{`echo`, `^2`},
			`^"echo^" ^^2`,
		},
		{
			[]string{`echo`, `^^`},
			`^"echo^" ^^^^`,
		},
		{
			[]string{`echo`, `^"`},
			`^"echo^" ^"^^\^"^"`,
		},
	}

	runTests(t, tests, WindowsCmdExe(WindowsArgv))
}
