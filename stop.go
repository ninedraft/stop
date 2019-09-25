package stop

import (
	"os"
	"runtime"
	"sync/atomic"
)

var exitStatus = new(int32)

// Stop sets global exit status and exits current goroutine with runtime.Goexit().
func Stop(status int) {
	atomic.StoreInt32(exitStatus, int32(status))
	runtime.Goexit()
}

// Status returns current exit status.
func Status() int {
	return int(atomic.LoadInt32(exitStatus))
}

// Exit calls os.Exit with stored non-zero global status.
func Exit() {
	var status = Status()
	if status != 0 {
		os.Exit(status)
	}
}
