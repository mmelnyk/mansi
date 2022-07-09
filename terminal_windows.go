package mansi

// VT Sequences for Windows 10 +
// https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences

import (
	"syscall"
)

var (
	kernel32Dll    = syscall.NewLazyDLL("Kernel32.dll")
	setConsoleMode = kernel32Dll.NewProc("SetConsoleMode")
)

func enableVirtualTerminalMode(stream syscall.Handle, enable bool) error {
	const EnableVirtualTerminalProcessing uint32 = 0x4

	var mode uint32
	err := syscall.GetConsoleMode(syscall.Stdout, &mode)
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
