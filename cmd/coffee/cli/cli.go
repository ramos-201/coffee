package cli

import (
	"coffee/cmd/coffee/internal/lexer"
	"coffee/cmd/coffee/internal/parser"
	"fmt"
	"os"
	"strings"
)

func RunCli() {
	if len(os.Args) < 2 {
		MsgExitError("No command provided", 2)
	}

	cmdName := os.Args[1]
	switch cmdName {
	case "run":
		executeRunCommand(os.Args[2:])
	default:
		MsgExitError("Unknown command '"+cmdName+"'", 2)
	}
}

func executeRunCommand(cmdArgs []string) {
	if len(cmdArgs) != 1 || !strings.HasSuffix(cmdArgs[0], ".cfe") {
		MsgExitError("Expected only one '.cfe' file", 2)
		return
	}

	filename := cmdArgs[0]

	contentFile, err := os.ReadFile(filename)
	if err != nil {
		MsgExitError("The file '"+filename+"' could not be opened or does not exist", 2)
		return
	}

	// Lexer
	l := lexer.New(string(contentFile))

	// Parser
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Println("Parser errors:")
		for _, msg := range p.Errors() {
			fmt.Println(" -", msg)
		}
		MsgExitError("Parsing failed", 2)
	}

	fmt.Println("Program parsed successfully!")
	fmt.Printf("AST:\n%+v\n", program)
}
