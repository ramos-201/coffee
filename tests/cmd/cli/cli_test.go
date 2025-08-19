package cli_test

import (
	"bytes"
	"os"
	"testing"

	"coffee/cmd/coffee/cli"
	"coffee/tests/testutils"

	"github.com/stretchr/testify/assert"
)

func TestRunCli_TerminateProcessWithEmptyArgs(t *testing.T) {
	// Tests running the CLI with no arguments
	// Args: "coffee" [empty arg]
	// Return: Wait for an error message and for the process to terminate
	rStderr, wStderr, restoreStderr := testutils.CaptureStderr()
	defer restoreStderr()

	rStdout, wStdout, restoreStdout := testutils.CaptureStdout()
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

	wStdout.Close()
	var bufOut bytes.Buffer
	bufOut.ReadFrom(rStdout)
	outStdout := bufOut.String()

	expectedOutput := "Error: No command provided\n"
	assert.Equal(t, expectedOutput, outStderr, "Output should match exactly")

	assert.Equal(t, "", outStdout, "stdout should be empty")
	assert.True(t, *exitCalled, "Expected OsExit to be called")
}

func TestRunCli_TerminateProcessWhitInvalidArgs(t *testing.T) {
	// Tests running the CLI with invalid arguments
	// Args: "coffee" invalid_test_arg
	// Return: Wait for an error message and for the process to terminate
	rStderr, wStderr, restoreStderr := testutils.CaptureStderr()
	defer restoreStderr()

	rStdout, wStdout, restoreStdout := testutils.CaptureStdout()
	defer restoreStdout()

	exitCalled, restoreExit := testutils.MockExitCalled()
	defer restoreExit()

	oldArgs := os.Args
	os.Args = []string{"coffee", "invalid_arg"}
	defer func() { os.Args = oldArgs }() // Reset os.Args after the test

	cli.RunCli()

	wStderr.Close()
	var buf bytes.Buffer
	buf.ReadFrom(rStderr)
	outStderr := buf.String()

	wStdout.Close()
	var bufOut bytes.Buffer
	bufOut.ReadFrom(rStdout)
	outStdout := bufOut.String()

	expectedOutput := "Error: Unknown command 'invalid_arg'\n"
	assert.Equal(t, expectedOutput, outStderr, "Output should match exactly")

	assert.Equal(t, "", outStdout, "stdout should be empty")
	assert.True(t, *exitCalled, "Expected OsExit to be called")
}
