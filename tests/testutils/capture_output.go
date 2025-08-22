package testutils

import "os"

// Temporarily redirect stderr to a pipe for print capture
func CaptureStderrMock() (r *os.File, w *os.File, restore func()) {
	oldStderr := os.Stderr
	r, w, _ = os.Pipe()
	os.Stderr = w

	return r, w, func() {
		_ = w.Close()
		_ = r.Close()
		os.Stderr = oldStderr
	}
}

// Temporarily redirect stdout to a pipe for print capture
func CaptureStdoutMock() (r *os.File, w *os.File, restore func()) {
	oldStdout := os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w

	return r, w, func() {
		_ = w.Close()
		_ = r.Close()
		os.Stdout = oldStdout
	}
}
