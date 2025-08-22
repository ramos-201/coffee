package testutils

import (
	"coffee/cmd/coffee/cli"
	"os"
)

// Mock cli.OsExit to prevent the program from terminating
func ExitCalledMock() (called *bool, restore func()) {
	exitCalled := false
	cli.OsExit = func(code int) { exitCalled = true }

	return &exitCalled, func() { cli.OsExit = os.Exit }
}
