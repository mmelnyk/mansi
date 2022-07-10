package main

import (
	"fmt"
	"time"

	"go.melnyk.org/mansi"
)

func main() {
	fmt.Print(mansi.SetWindowTitle("Example Table"))
	fmt.Print(mansi.ScreenEnableAltBuffer + mansi.CursorHome)
	fmt.Println(mansi.ColorBold+mansi.ColorTextBrightWhite, "Example Table", mansi.ResetColor)

	if !mansi.IsTerminal() {
		fmt.Println("Is terminal available?")
	}

	fmt.Println(mansi.ColorTextBrightYellow, "Bright Yellow", mansi.ColorTextBrightGreen, "Bright Green", mansi.ColorTextBrightBlue, "Bright Blue", mansi.ResetColor, mansi.ColorTextBlack)
	fmt.Println(mansi.ColorBackgroundYellow, "       Yellow", mansi.ColorBackgroundGreen, "       Green", mansi.ColorBackgroundBlue, "       Blue", mansi.ResetColor, mansi.ColorTextBrightWhite)
	fmt.Println(mansi.ColorBackgroundRed, "          Red", mansi.ColorBackgroundMagenta, "     Magenta", mansi.ColorBackgroundCyan, "       Cyan", mansi.ResetColor)

	time.Sleep(3 * time.Second)
	fmt.Print(mansi.ScreenDisableAltBuffer)
}
