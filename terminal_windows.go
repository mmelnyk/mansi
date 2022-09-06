package mansi

// VT Sequences for Windows 10 +
// https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences

import (
	"syscall"
	"unsafe"
)

var (
	kernel32Dll                = syscall.NewLazyDLL("Kernel32.dll")
	setConsoleMode             = kernel32Dll.NewProc("SetConsoleMode")
	getConsoleScreenBufferInfo = kernel32Dll.NewProc("GetConsoleScreenBufferInfo")
)

func enableVirtualTerminalMode(stream syscall.Handle, enable bool) error {
	const EnableVirtualTerminalProcessing uint32 = 0x4

	var mode uint32
	err := syscall.GetConsoleMode(stream, &mode)
	if err != nil {
		return err
	}

	if enable {
		mode |= EnableVirtualTerminalProcessing
	} else {
		mode &^= EnableVirtualTerminalProcessing
	}

	if ret, _, err := setConsoleMode.Call(uintptr(stream), uintptr(mode)); ret == 0 {
		return err
	}

	return nil
}

func isTerminal() bool {
	return enableVirtualTerminalMode(syscall.Stdout, true) == nil
}

type coord struct {
	X int16
	Y int16
}

type smallRect struct {
	Left   int16
	Top    int16
	Right  int16
	Bottom int16
}

// Used with GetConsoleScreenBuffer to retrieve information about a console
// screen buffer. See
// https://docs.microsoft.com/en-us/windows/console/console-screen-buffer-info-str
// for details.

type consoleScreenBufferInfo struct {
	Size              coord
	CursorPosition    coord
	Attributes        uint16
	Window            smallRect
	MaximumWindowSize coord
}

func getConsoleSize() (int, int, error) {
	var info consoleScreenBufferInfo

	if ret, _, err := getConsoleScreenBufferInfo.Call(uintptr(syscall.Stdout), uintptr(unsafe.Pointer(&info))); ret == 0 {
		return -1, -1, err
	}

	return int(info.Window.Right - info.Window.Left + 1), int(info.Window.Bottom - info.Window.Top + 1), nil
}
