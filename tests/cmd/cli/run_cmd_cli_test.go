package cli_test

import (
	"bytes"
	"coffee/cmd/coffee/cli"
	"coffee/tests/testutils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCli_RunCmd_TerminateProcessWithInvalidArgs(t *testing.T) {
	// Test: Executing 'run' command with invalid arguments
	// Expect: Print error message and process termination
	testCases := []struct {
		name string
		args []string
	}{
		// Args:
		{
			name: "It was sent without argument",
			args: []string{"coffee", "run"},
		},
		{
			name: "An invalid argument is sent",
			args: []string{"coffee", "run", "invalid_arg"},
		},
		{
			name: "The valid argument is sent with multiple additional arguments",
			args: []string{"coffee", "run", "filename.cfe", "other_arg"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rStderr, wStderr, restoreStderr := testutils.CaptureStderrMock()
			defer restoreStderr()

			cli.ErrorOutput = wStderr
			defer func() { cli.ErrorOutput = os.Stderr }()

			rStdout, wStdout, restoreStdout := testutils.CaptureStdoutMock()
			defer restoreStdout()

			exitCalled, exitCode, runCli, restoreExit := testutils.ExitCalledMock()
			defer restoreExit()

			oldArgs := os.Args
			os.Args = tc.args
			defer func() { os.Args = oldArgs }() // Reset os.Args after the test

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

			expectedOutput := "Expected only one '.cfe' file\n"
			assert.Equal(t, expectedOutput, outStderr, "Output should match exactly")

			assert.Equal(t, "", outStdout, "stdout should be empty")
			assert.True(t, *exitCalled, "Expected OsExit to be called")
			assert.Equal(t, 2, *exitCode, "A correct exit code is expected")
		})
	}
}

func TestRunCli_RunCmd_TerminateProcessWithFileOpenError(t *testing.T) {
	// Test: Executing 'run' command with nonexistent file
	// Args: "coffee run nonexistent_file.cfe"
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
	filename := "nonexistent_file.cfe"
	os.Args = []string{"coffee", "run", filename}
	defer func() { os.Args = oldArgs }() // Reset os.Args after the test

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

	expectedOutput := "The file '" + filename + "' could not be opened or does not exist\n"
	assert.Equal(t, expectedOutput, outStderr, "Output should match exactly")

	assert.Equal(t, "", outStdout, "stdout should be empty")
	assert.True(t, *exitCalled, "Expected OsExit to be called")
	assert.Equal(t, 2, *exitCode, "A correct exit code is expected")
}
