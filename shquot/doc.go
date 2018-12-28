// Package shquot contains various functions for quoting sequences of
// command arguments for literal interpretation by different shells and other
// similar intermediaries that process incoming command lines.
//
// The functions all have the same signature, defined as type Q in this package,
// taking an array of arguments in the usual style passed to "execve" on a
// Unix/POSIX system.
//
// While calling "execve" directly is always preferable to avoid
// misinterpretation by intermediaries, sometimes such preprocessing cannot
// be avoided. For example, remote command execution protocols like SSH often
// expect a single string to be interpreted by a shell.
//
// Since each shell or intermediary has different details, it's important to
// select the correct quoting function for the target system or else the
// result may be misinterpreted.
package shquot
