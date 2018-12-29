# Go Shell Quoting Utilities

This module has various helper functions for quoting and escaping an array
of strings as expected by the "exec" calls on a Unix system into single strings
that should be able to pass through shells and other intermediaries to produce
the same result.

There is no single common quoting format that works for all shells and other
layers, so this module makes no attempt to abstract away these differences.
Instead, it provides a family of functions with the same signature so that the
caller may select -- possibly dynamically -- a suitable quoting function to use.

For more information, see [the package reference documentation](https://godoc.org/github.com/apparentlymart/go-shquot/shquot).
