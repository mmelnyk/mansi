package mansi

var (
	isTerminalCached = isTerminal()
)

// IsTerminal is returning true if StdOut is terminal
func IsTerminal() bool {
	return isTerminalCached
}
