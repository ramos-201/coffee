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
	// Test: Run CLI with no arguments
	// Values: "coffee"
	// Expect: Print error message and process termination
	rStderr, wStderr, restoreStderr := testutils.CaptureStderrMock()
	defer restoreStderr()

	cli.ErrorOutput = wStderr
	defer func() { cli.ErrorOutput = os.Stderr }()

	rStdout, wStdout, restoreStdout := testutils.CaptureStdoutMock()
	defer restoreStdout()

	exitCalled, exitCode, runCli, restoreExit := testutils.ExitCalledMock()
	defer restoreExit()

	oldArgs := os.Args
	os.Args = []string{"coffee"}
	defer func() { os.Args = oldArgs }()

	cli.SetExitStatus(0)

	runCli(cli.RunCli)

	wStderr.Close()
	var buf bytes.Buffer
	buf.ReadFrom(rStderr)
	outStderr := buf.String()

	wStdout.Close()
	var bufOut bytes.Buffer
	bufOut.ReadFrom(rStdout)
	outStdout := bufOut.String()

	expectedOutput := "No command provided\n"
	assert.Equal(t, expectedOutput, outStderr, "Output should match exactly")

	assert.Equal(t, "", outStdout, "stdout should be empty")
	assert.True(t, *exitCalled, "Expected OsExit to be called")
	assert.Equal(t, 1, *exitCode, "Expected exit code to be 1")
}
