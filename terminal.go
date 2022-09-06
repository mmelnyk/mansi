package mansi

var (
	isTerminalCached = isTerminal()
)

// IsTerminal is returning true if StdOut is terminal
func IsTerminal() bool {
	return isTerminalCached
}

// GetConsoleSize returns the visible dimensions of the given terminal.
// These dimensions don't include any scrollback buffer height.
func GetConsoleSize() (width, height int, err error) {
	return getConsoleSize()
}
