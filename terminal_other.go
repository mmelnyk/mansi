//go:build darwin || dragonfly || freebsd || (linux && !appengine) || netbsd || openbsd

package mansi

import (
	"errors"
	"syscall"
	"unsafe"
)

type windowsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func isTerminal() bool {
	var termios Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(syscall.Stdout), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

func getConsoleSize() (int, int, error) {
	var cs windowsize
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(syscall.Stdout), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&cs)), 0, 0, 0)
	if err != 0 {
		return -1, -1, errors.New("syscall error")
	}

	return int(cs.Col), int(cs.Row), nil
}
