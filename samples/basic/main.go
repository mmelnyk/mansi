package main

import (
	"fmt"

	"go.melnyk.org/mansi"
)

func main() {
	fmt.Println(mansi.ColorBold+mansi.ColorTextBrightWhite, "Example Table", mansi.ResetColor)

	if !mansi.IsTerminal() {
		fmt.Println("Is terminal available?")
	}

	fmt.Println(mansi.ColorTextBrightYellow, "Bright Yellow", mansi.ColorTextBrightGreen, "Bright Green", mansi.ColorTextBrightBlue, "Bright Blue", mansi.ResetColor, mansi.ColorTextBlack)
	fmt.Println(mansi.ColorBackgroundYellow, "       Yellow", mansi.ColorBackgroundGreen, "       Green", mansi.ColorBackgroundBlue, "       Blue", mansi.ResetColor, mansi.ColorTextBrightWhite)
	fmt.Println(mansi.ColorBackgroundRed, "          Red", mansi.ColorBackgroundMagenta, "     Magenta", mansi.ColorBackgroundCyan, "       Cyan", mansi.ResetColor)
}
