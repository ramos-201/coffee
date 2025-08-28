package cli

import (
	"os"
	"sync"
)

var (
	exitStatus  int
	exitMutex   sync.Mutex
	OsExit      = os.Exit // OsExit allows for mocking in tests
	ErrorOutput = os.Stderr
)

func MsgExitError(msg string, code int) {
	ErrorOutput.Write([]byte(msg + "\n"))
	SetExitStatus(code)
	Exit()
}

func SetExitStatus(code int) {
	exitMutex.Lock()
	if code > exitStatus {
		exitStatus = code
	}
	exitMutex.Unlock()
}

func Exit() {
	exitMutex.Lock()
	status := exitStatus
	exitMutex.Unlock()
	OsExit(status)
}
