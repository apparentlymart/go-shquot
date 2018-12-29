package shquot

import (
	"testing"
)

func TestViaPowerShell(t *testing.T) {
	tests := []qTest{
		{
			[]string{},
			"",
		},
		{
			[]string{"cmd.exe"},
			`& Start-Process -FilePath "cmd.exe"`,
		},
		{
			[]string{"cmd.exe", "/c", "echo hello"},
			"& Start-Process -FilePath \"cmd.exe\" -ArgumentList \"/c `\"echo hello`\"\"",
		},
		{
			[]string{"echo", "hello $name"},
			"& Start-Process -FilePath \"echo\" -ArgumentList \"`\"hello `$name`\"\"",
		},
	}

	runTests(t, tests, ViaPowerShell(WindowsArgvSplit))
}
