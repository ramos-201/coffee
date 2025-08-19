package cli

import (
	"os"
)

var OsExit = os.Exit // OsExit allows for mocking in tests

var commands string = "example"

func RunCli() {
	if len(os.Args) < 2 {
		ExitError("Error: No command provided")
		return
	}

	cmdName := os.Args[1]

	if cmdName != commands {
		ExitError("Error: Unknown command '" + cmdName + "'")
	}
}

func ExitError(msg string) {
	os.Stderr.WriteString(msg + "\n")
	OsExit(1)
}
