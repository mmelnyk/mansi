//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package mansi

import (
	"syscall"
)

type Termios syscall.Termios

const ioctlReadTermios = syscall.TIOCGETA
