package cli

import (
	"os"
	"strings"
)

var OsExit = os.Exit // OsExit allows for mocking in tests

var commands string = "run"

func RunCli() {
	if len(os.Args) < 2 {
		ExitError("Error: No command provided")
		return
	}

	cmdName := os.Args[1]
	if cmdName != commands {
		ExitError("Error: Unknown command '" + cmdName + "'")
		return
	}

	cmdArgs := os.Args[2:]
	if (len(cmdArgs) != 1) || (!strings.HasSuffix(cmdArgs[0], ".cfe")) {
		ExitError("Error: Expected only one '.cfe' file")
	}
}

func ExitError(msg string) {
	os.Stderr.WriteString(msg + "\n")
	OsExit(1)
}
