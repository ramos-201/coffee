package cli

import (
	"bufio"
	"os"
	"strings"
)

var OsExit = os.Exit // OsExit allows for mocking in tests

func RunCli() {
	if len(os.Args) < 2 {
		ExitError("No command provided")
		return
	}

	cmdName := os.Args[1]

	switch cmdName {
	case "run":
		executeRunCommand(os.Args[2:])
		return
	default:
		ExitError("Unknown command '" + cmdName + "'")
		return
	}
}

func executeRunCommand(cmdArgs []string) {
	if (len(cmdArgs) != 1) || (!strings.HasSuffix(cmdArgs[0], ".cfe")) {
		ExitError("Expected only one '.cfe' file")
		return
	}

	filename := cmdArgs[0]
	file, err := os.Open(filename)

	if err != nil {
		ExitError("The file '" + filename + "' could not be opened or does not exist")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineStr := strings.TrimSpace(line)
		if lineStr == "" {
			continue
		}
	}
}

func ExitError(msg string) {
	os.Stderr.WriteString(msg + "\n")
	OsExit(1)
}
