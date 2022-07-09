package mansi

import (
	"syscall"
)

type Termios syscall.Termios

const ioctlReadTermios = syscall.TCGETS
