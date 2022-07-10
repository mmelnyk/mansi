package mansi

// See ANSI Escape Codes - https://en.wikipedia.org/wiki/ANSI_escape_code

const (
	// ESCape codes
	Bell           = "\007" // Makes an audible noise.
	Backspace      = "\010" // Moves the cursor left (but may "backwards wrap" if cursor is at start of line).
	Tab            = "\011" // Moves the cursor right to next multiple of 8.
	LineFeed       = "\012" // Moves to next line, scrolls the display up if at bottom of the screen. Usually does not move horizontally, though programs should not rely on this.
	FormFeed       = "\014" // Move a printer to top of next page. Usually does not move horizontally, though programs should not rely on this. Effect on video terminals varies.
	CarriageReturn = "\015" // Moves the cursor to column zero.
	Escape         = "\033" // Starts all the escape sequences
	BEL            = Bell
	BS             = Backspace
	HT             = Tab
	LF             = LineFeed
	FF             = FormFeed
	ESC            = Escape    // Starts all the escape sequences
	CSI            = ESC + "[" // Control Sequence Introducer
	OSC            = ESC + "]" // Operating System Command
)
