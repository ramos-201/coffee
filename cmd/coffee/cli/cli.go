package cli

import (
	"fmt"
	"os"
)

func RunCli() {
	if len(os.Args) < 2 {
		ExitError("No command provided", 2)
	}

	cmdName := os.Args[1]
	switch cmdName {
	case "run":
		return
	default:
		ExitError("Unknown command '"+cmdName+"'", 2)
	}

	fmt.Println("Continue without error / tests")
	ErrorOutput.Write([]byte("continue without error / tests\n"))
}
