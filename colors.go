package mansi

const (
	//Colors
	ResetColor  = CSI + "0m" // Reset color settings to default
	ColorBold   = CSI + "1m"
	ColorFaint  = CSI + "2m"
	ColorNormal = CSI + "22m"

	ColorTextBlack   = CSI + "30m" // Black color for text
	ColorTextRed     = CSI + "31m" // Red color for text
	ColorTextGreen   = CSI + "32m" // Green color for text
	ColorTextYellow  = CSI + "33m" // Yellow color for text
	ColorTextBlue    = CSI + "34m" // Blue color for text
	ColorTextMagenta = CSI + "35m" // Magenta color for text
	ColorTextCyan    = CSI + "36m" // Cyan color for text
	ColorTextWhite   = CSI + "37m" // White color for text

	ColorBackgroundBlack   = CSI + "40m" // Black color for background
	ColorBackgroundRed     = CSI + "41m" // Red color for background
	ColorBackgroundGreen   = CSI + "42m" // Green color for background
	ColorBackgroundYellow  = CSI + "43m" // Yellow color for background
	ColorBackgroundBlue    = CSI + "44m" // Blue color for background
	ColorBackgroundMagenta = CSI + "45m" // Magenta color for background
	ColorBackgroundCyan    = CSI + "46m" // Cyan color for background
	ColorBackgroundWhite   = CSI + "47m" // White color for background

	ColorTextBrightBlack   = CSI + "90m" // Bright black color for text
	ColorTextBrightRed     = CSI + "91m" // Bright red color for text
	ColorTextBrightGreen   = CSI + "92m" // Bright green color for text
	ColorTextBrightYellow  = CSI + "93m" // Bright yellow color for text
	ColorTextBrightBlue    = CSI + "94m" // Bright blue color for text
	ColorTextBrightMagenta = CSI + "95m" // Bright magenta color for text
	ColorTextBrightCyan    = CSI + "96m" // Bright cyan color for text
	ColorTextBrightWhite   = CSI + "97m" // Bright white color for text

	ColorBackgroundBrightBlack   = CSI + "100m" // Bright black color for background
	ColorBackgroundBrightRed     = CSI + "101m" // Bright red color for background
	ColorBackgroundBrightGreen   = CSI + "102m" // Bright green color for background
	ColorBackgroundBrightYellow  = CSI + "103m" // Bright yellow color for background
	ColorBackgroundBrightBlue    = CSI + "104m" // Bright blue color for background
	ColorBackgroundBrightMagenta = CSI + "105m" // Bright magenta color for background
	ColorBackgroundBrightCyan    = CSI + "106m" // Bright cyan color for background
	ColorBackgroundBrightWhite   = CSI + "107m" // Bright white color for background
)
