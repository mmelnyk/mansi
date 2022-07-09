package mansi

const (
	// Controls
	EraseDisplay = CSI + "2J"
	EraseLine    = CSI + "K"
	MoveTop      = CSI + "H" // Format is CSI Line;Column H
	MoveBegin    = CSI + "G" // Format is CSI Column G
	SavePos      = CSI + "s"
	RestorePos   = CSI + "u"

	// Tabs
	TabSet   = ESC + "H"
	TabClear = CSI + "3g"
	Tab      = CSI + "I"
	TabBack  = CSI + "Z"
)
