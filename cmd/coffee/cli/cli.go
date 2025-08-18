package cli

import (
	"os"
)

var OsExit = os.Exit // OsExit allows for mocking in tests

func RunCli() {
	if len(os.Args) < 2 {
		ExitError(
			"Error: Missing command",
			"Use 'help' to see available commands",
			"Example: 'coffee <command> <args>'",
		)
	}
}

func ExitError(msgs ...string) {
	for _, msg := range msgs {
		os.Stderr.WriteString(msg + "\n")
	}
	OsExit(1)
}
