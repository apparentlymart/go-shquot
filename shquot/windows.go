package shquot

import (
	"strings"
)

// WindowsArgv quotes arguments using the conventions expected by the Windows
// API function CommandLineToArgvW.
//
// On Windows the final parsing of a command line string is the responsibility
// of the application itself, and so some applications may not honor these
// conventions but most modern command line applications do, due to this step
// being handled automatically by the Microsoft Visual C++ runtime library
// prior to calling the application's entry point.
func WindowsArgv(cmdline []string) string {
	if len(cmdline) == 0 {
		return ""
	}

	var buf strings.Builder
	windowsArgvSingle(cmdline[0], &buf)
	for _, a := range cmdline[1:] {
		buf.WriteByte(' ')
		windowsArgvSingle(a, &buf)
	}
	return buf.String()
}

func windowsArgvSingle(a string, to *strings.Builder) {
	if len(a) > 0 && !strings.ContainsAny(a, " \t\n\v\"") {
		// No quoting required, then.
		to.WriteString(a)
		return
	}

	to.WriteByte('"')
	bs := 0
	for _, c := range a {
		switch c {
		case '\\':
			bs++
			continue
		case '"':
			// All of the backslashes we saw so far must be escaped, and then
			// we need one more backslash for the quote character.
			to.WriteString(strings.Repeat("\\", bs*2+1))
			to.WriteRune(c)
			bs = 0
		default:
			// If we encounter anything other than a quote or a backslash
			// then any preceding backslashes we've seen are _not_ special and
			// so we must write them out literally first.
			if bs > 0 {
				to.WriteString(strings.Repeat("\\", bs))
			}
			to.WriteRune(c)
			bs = 0
		}
	}
	// If any backslashes are pending once we exit then we need to double them
	// all up so that the closing quote will _not_ be interpreted as an escape.
	if bs > 0 {
		to.WriteString(strings.Repeat("\\", bs*2))
	}
	to.WriteByte('"')
}

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
