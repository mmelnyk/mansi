package main

import (
	"fmt"

	"go.melnyk.org/mansi"
)

func main() {
	if !mansi.IsTerminal() {
		fmt.Println("Is terminal available?")
		return
	}

	fmt.Println(mansi.ColorBold+mansi.ColorTextBrightWhite, "Example Table", mansi.ResetColor)

	fmt.Println(mansi.TabClearAll+mansi.ColorBackgroundBlue+"Position", mansi.TabSet, "| Name          ", mansi.TabSet, "| Phone         ", mansi.TabSet, mansi.ResetColor)
	fmt.Println("#1", mansi.TabForward, "| Adam", mansi.TabForward, "| (111)222-3344", mansi.TabForward)
	fmt.Println(mansi.ColorTextBrightGreen+"#2", mansi.TabForward, "| Benjamin", mansi.TabForward, "| (111)222-3344", mansi.TabForward, mansi.ResetColor)
	fmt.Println("#3", mansi.TabForward, "| Barbara", mansi.TabForward, "| (111)222-3344", mansi.TabForward)
	fmt.Println(mansi.ColorTextBrightGreen+"#4", mansi.TabForward, "| Helena", mansi.TabForward, "| (111)222-3344", mansi.TabForward, mansi.ResetColor)
	fmt.Println("#5", mansi.TabForward, "| Jackson", mansi.TabForward, "| (111)222-3344", mansi.TabForward)
	fmt.Println(mansi.SoftReset)
}
