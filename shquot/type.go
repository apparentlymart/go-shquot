package shquot

// Q is the signature of all command line quoting functions in this package.
// This may be useful for calling applications that select dynamically which
// quoting mechanism to use and store a reference to the appropriate function
// to call later.
//
// cmdline is a slice of string arguments where the first element is
// conventionally the command itself and any remaining elements are arguments
// to that command. This mimics the way command lines are passed to the execve
// function on a Unix (POSIX) system.
type Q func(cmdline []string) string
