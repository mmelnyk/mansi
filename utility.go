package mansi

import "strconv"

// CursorUpBy returns an escape sequence to move the cursor lines cells in the given direction. If the cursor is already at the edge of the screen, this has no effect.
func CursorUpBy(lines int) string {
	if lines >= 1 {
		return CSI + strconv.Itoa(lines) + "A"
	} else {
		return ""
	}
}

// CursorDownBy returns an escape sequence to move the cursor lines cells in the given direction. If the cursor is already at the edge of the screen, this has no effect.
func CursorDownBy(lines int) string {
	if lines >= 1 {
		return CSI + strconv.Itoa(lines) + "B"
	} else {
		return ""
	}
}

// CursorForwardBy returns an escape sequence to move the cursor lines cells in the given direction. If the cursor is already at the edge of the screen, this has no effect.
func CursorForwardBy(columns int) string {
	if columns >= 1 {
		return CSI + strconv.Itoa(columns) + "C"
	} else {
		return ""
	}
}

// CursorBackBy returns an escape sequence to move the cursor lines cells in the given direction. If the cursor is already at the edge of the screen, this has no effect.
func CursorBackBy(columns int) string {
	if columns >= 1 {
		return CSI + strconv.Itoa(columns) + "D"
	} else {
		return ""
	}
}

// CursorPosition returns an escape sequence to move the cursor to a coordinate pair,
// where (1, 1) is the origin (top-left corner).
func CursorPosition(column, line int) string {
	if column <= 0 {
		column = 1
	}
	if line <= 0 {
		line = 1
	}

	return CSI + strconv.Itoa(line) + ";" + strconv.Itoa(column) + "H"
}

// SetWindowTitle returns an escape sequence to set console window title.
func SetWindowTitle(title string) string {
	return OSC + "0;" + title + BEL
}
