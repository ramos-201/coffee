package testutils

import (
	"coffee/cmd/coffee/cli"
)

// Mock cli.OsExit to prevent the program from terminating
func ExitCalledMock() (called *bool, code *int, run func(f func()), restore func()) {
	originalExit := cli.OsExit

	exitCalled := false
	exitCode := 0

	cli.OsExit = func(c int) {
		exitCalled = true
		exitCode = c
		panic("exit") // Stop Cli simulating a real exit
	}

	restoreFunc := func() {
		cli.OsExit = originalExit
	}

	runFunc := func(runF func()) {
		defer func() {
			if r := recover(); r != nil && r != "exit" {
				panic(r)
			}
		}()
		runF()
	}

	return &exitCalled, &exitCode, runFunc, restoreFunc
}
