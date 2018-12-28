package shquot

import (
	"strings"
)

// PowerShell quotes the given command line for interpretation as a command
// line in Microsoft PowerShell.
//
// This function is suitable for use with command lines that run PowerShell
// cmdlets. It is not suitable for using PowerShell to run a separate Windows
// process because in that situation the command line will be processed both by
// PowerShell and by the target program itself, and must therefore have two
// levels of quoting applied to it. Windows applications can parse their
// command line strings any way they like, but most applications use the
// algorithm targeted by the WindowsArgv function in this package, in which
// case the PowerShellWindowsArgv function can be used to apply both levels of
// quoting at once.
func PowerShell(cmdline []string) string {
	if len(cmdline) == 0 {
		return ""
	}

	var buf strings.Builder
	powershellSingle(cmdline[0], &buf)
	for _, a := range cmdline[1:] {
		buf.WriteByte(' ')
		powershellSingle(a, &buf)
	}
	return buf.String()
}

func powershellSingle(a string, to *strings.Builder) {
}

// PowerShellWindowsArgv applies both WindowsArgv quoting and PowerShell
// quoting to the given command line, thus allowing PowerShell to launch
// the given command line as a separate process on a Windows system.
//
// This strange function is necessary because on Windows command line parsing
// is handled by the application itself, but PowerShell does Unix-shell-like
// tokenization of the command line first, before gluing the parts back together
// to pass to the launched program.
func PowerShellWindowsArgv(cmdline []string) string {
	if len(cmdline) == 0 {
		return ""
	}

	var buf strings.Builder
	buf.WriteByte('&')
	for i, a := range cmdline[1:] {
		if i > 0 {
			buf.WriteByte(' ')
		}
		var thisBuf strings.Builder
		windowsArgvSingle(a, &thisBuf)
		powershellSingle(thisBuf.String(), &buf)
	}
	return buf.String()
}
