//go:build darwin || dragonfly || freebsd || (linux && !appengine) || netbsd || openbsd

package mansi

import (
	"syscall"
	"unsafe"
)

func isTerminal() bool {
	var termios Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(syscall.Stdout), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
