package shquot

import (
	"strings"
)

// WindowsCmdExe produces a quoting function that prepares a command line to
// pass through the Windows command interpreter cmd.exe.
//
// Since cmd.exe is just an intermediary, the caller must provide another
// quoting function that deals with the subsequent layer of quoting. On
// Windows most of the command line processing is actually delegated to the
// application itself rather than the command interpreter, and so which
// wrapped quoting function to select depends on the target program.
// Most modern command line applications use the CommandLineToArgvW function
// for argument processing, and its escaping rules are implemented by
// WindowsArgv in this package.
//
// Note that this extra level of quoting is necessary only for command lines
// that will pass through the command interpreter, such as generated command
// scripts. If you're calling the Windows CreateProcess API directly then you
// must not apply cmd.exe quoting, or the result will be incorrectly parsed.
func WindowsCmdExe(wrapped Q) Q {
	r := strings.NewReplacer(
		"(", "^(",
		")", "^)",
		"%", "^%",
		"!", "^!",
		"^", "^^",
		`"`, `^"`,
		"<", "^<",
		">", "^>",
		"&", "^&",
		"|", "^|",
	)
	return func(cmdline []string) string {
		s := wrapped(cmdline)
		return r.Replace(s)
	}
}
