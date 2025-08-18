package cli_test

import (
	"bytes"
	"os"
	"testing"

	"coffee/cmd/coffee/cli"
	"coffee/tests/testutils"

	"github.com/stretchr/testify/assert"
)

func TestRunCli_ExitsWhenNoCommandProvided(t *testing.T) {
	rStderr, wStderr, restoreStderr := testutils.CaptureStderr()
	defer restoreStderr()

	rStout, wStout, restoreStdout := testutils.CaptureStdout()
	defer restoreStdout()

	exitCalled, restoreExit := testutils.MockExitCalled()
	defer restoreExit()

	oldArgs := os.Args
	os.Args = []string{"coffee"}
	defer func() { os.Args = oldArgs }() // Reset os.Args after the test

	cli.RunCli()

	wStderr.Close()
	var buf bytes.Buffer
	buf.ReadFrom(rStderr)
	outStderr := buf.String()

	wStout.Close()
	var bufOut bytes.Buffer
	bufOut.ReadFrom(rStout)
	outStout := bufOut.String()

	expectedOutput := "Error: Missing command\n" +
		"Use 'help' to see available commands\n" +
		"Example: 'coffee <command> <args>'\n"

	assert.Equal(t, expectedOutput, outStderr, "Output should match exactly")
	assert.Equal(t, "", outStout, "stdout should be empty")
	assert.True(t, *exitCalled, "Expected OsExit to be called")
}
