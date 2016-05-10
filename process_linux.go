package phantomjs

import (
	"syscall"
)

func killProcess(pid int, handlePtr uintptr) {
	syscall.Kill(pid, syscall.SIGKILL)
}
