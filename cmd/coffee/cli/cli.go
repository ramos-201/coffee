package cli

import (
	"coffee/cmd/coffee/lexer"
	"fmt"
	"os"
	"strings"
)

func RunCli() {
	if len(os.Args) < 2 {
		ExitError("No command provided", 2)
	}

	cmdName := os.Args[1]
	switch cmdName {
	case "run":
		executeRunCommand(os.Args[2:])
	default:
		ExitError("Unknown command '"+cmdName+"'", 2)
	}
}

func executeRunCommand(cmdArgs []string) {
	if (len(cmdArgs) != 1) || (!strings.HasSuffix(cmdArgs[0], ".cfe")) {
		ExitError("Expected only one '.cfe' file", 2)
		return
	}

	filename := cmdArgs[0]

	contentFile, err := os.ReadFile(filename)
	if err != nil {
		ExitError("The file '"+filename+"' could not be opened or does not exist", 2)
		return
	}

	l := lexer.New(string(contentFile))

	for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
