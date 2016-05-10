package phantomjs

import (
	"syscall"
)

func killProcess(pid int, handlePtr uintptr) {
	syscall.TerminateProcess(syscall.Handle(handlePtr), 0)
}
