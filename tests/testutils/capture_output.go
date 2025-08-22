package testutils

import "os"

// Temporarily redirect stderr to a pipe for print capture
func CaptureStderrMock() (r *os.File, w *os.File, restore func()) {
	oldStderr := os.Stderr
	r, w, _ = os.Pipe()
	os.Stderr = w

	restoreFunc := func() {
		_ = w.Close()
		_ = r.Close()
		os.Stderr = oldStderr
	}

	return r, w, restoreFunc
}

// Temporarily redirect stdout to a pipe for print capture
func CaptureStdoutMock() (r *os.File, w *os.File, restore func()) {
	oldStdout := os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w

	restoreFunc := func() {
		_ = w.Close()
		_ = r.Close()
		os.Stdout = oldStdout
	}

	return r, w, restoreFunc
}
