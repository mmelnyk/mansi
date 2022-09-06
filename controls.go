package mansi

const (
	// Cursor
	CursorUp              = CSI + "A"    // Moves the cursor one cell up. If the cursor is already at the edge of the screen, this has no effect.
	CursorDown            = CSI + "B"    // Moves the cursor one cell down. If the cursor is already at the edge of the screen, this has no effect.
	CursorForward         = CSI + "C"    // Moves the cursor one cell right. If the cursor is already at the edge of the screen, this has no effect.
	CursorBack            = CSI + "D"    // Moves the cursor one cell left. If the cursor is already at the edge of the screen, this has no effect.
	CursorHome            = CSI + "H"    // Moves the cursor to top left corner of console.
	CursorSavePosition    = CSI + "s"    // Saves the cursor position/state in SCO console mode.
	CursorRestorePosition = CSI + "u"    // Restores the cursor position/state in SCO console mode.
	CursorHide            = CSI + "?25l" // Hides the cursor.
	CursorShow            = CSI + "?25h" // Shows the cursor.
	CursorNextLine        = CSI + "E"    // Moves the cursor to the beginning of the next line. If the cursor is already at the edge of the screen, this has no effect.
	CursorPreviousLine    = CSI + "F"    // Moves the cursor to the beginning of the previous line. If the cursor is already at the edge of the screen, this has no effect.

	// Screen
	ScreenErase            = CSI + "2J"     // Clear entire screen.
	ScreenEraseAll         = CSI + "3J"     // Clear entire screen and delete all lines saved in the scrollback buffer.
	ScreenEraseToBeginning = CSI + "1J"     // Clear from cursor to beginning of the screen.
	ScreenEraseToEnd       = CSI + "J"      // Clear from cursor to end of the screen.
	ScreenEnableAltBuffer  = CSI + "?1049h" // Enable alternative screen buffer.
	ScreenDisableAltBuffer = CSI + "?1049l" // Disable alternative screen buffer.

	// Line
	LineEraseToEnd       = CSI + "K"  // Clear line from cursor to the end of the line. Cursor position does not change.
	LineEraseToBeginning = CSI + "1K" // Clear line from cursor to beginning of the line. Cursor position does not change.
	LineErase            = CSI + "2K" // Clear entire line. Cursor position does not change.

	// Tabs
	TabSet      = ESC + "H"  // Sets a tab stop in the current column the cursor is in.
	TabClear    = CSI + "0g" // Clears the tab stop in the current column, if there is one. Otherwise does nothing..
	TabClearAll = CSI + "3g" // Clears all currently set tab stops.
	TabForward  = CSI + "I"  // Advance the cursor to the next column (in the same row) with a tab stop. If there are no more tab stops, move to the last column in the row. If the cursor is in the last column, move to the first column of the next row.
	TabBackward = CSI + "Z"  // Move the cursor to the previous column (in the same row) with a tab stop. If there are no more tab stops, moves the cursor to the first column. If the cursor is in the first column, doesn’t move the cursor.

	// Reset
	Reset     = ESC + "c"  // Triggers a full reset of the terminal to its original state. This may include (if applicable): reset graphic rendition, clear tabulation stops, reset to default font, and more
	SoftReset = CSI + "!p" // Reset certain terminal settings to their defaults.

	// Scroll
	ScrollUp   = CSI + "S" // Scroll the screen up one line. The top line is lost.
	ScrollDown = CSI + "T" // Scroll the screen down one line. The bottom line is lost.
)
