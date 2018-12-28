package shquot_test

import (
	"fmt"

	"github.com/apparentlymart/go-shquot/shquot"
)

func Example() {
	cmdline := []string{`echo`, `Hello, world!`}
	fmt.Println("POSIXShell:", shquot.POSIXShell(cmdline))
	fmt.Println("WindowsArgv:", shquot.WindowsArgv(cmdline))
	fmt.Println("WindowsCmdExe+WindowsArgv:", shquot.WindowsCmdExe(shquot.WindowsArgv)(cmdline))
	fmt.Println("Dockerfile:", shquot.Dockerfile(cmdline))

	// Output:
	// POSIXShell: echo 'Hello, world!'
	// WindowsArgv: echo "Hello, world!"
	// WindowsCmdExe+WindowsArgv: echo ^"Hello, world^!^"
	// Dockerfile: ["echo","Hello, world!"]
}
