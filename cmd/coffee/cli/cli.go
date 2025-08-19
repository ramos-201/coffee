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
		return
	}

	// RunCommand ...
	filename := cmdArgs[0]
	file, err := os.Open(filename)

	if err != nil {
		ExitError("Error: The file '" + filename + "' could not be opened or does not exist")
		return
	}
	defer file.Close()
}

func ExitError(msg string) {
	os.Stderr.WriteString(msg + "\n")
	OsExit(1)
}
