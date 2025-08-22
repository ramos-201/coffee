package cli

import (
	"fmt"
	"os"
)

func RunCli() {
	if len(os.Args) < 2 {
		ExitError("No command provided", 1)
	}

	fmt.Println("Continue without error / tests")
	ErrorOutput.Write([]byte("continue without error / tests\n"))
}
