# mANSI
[![License][license-img]][license] [![Actions Status][action-img]][action] [![GoDoc][godoc-img]][godoc] [![Go Report Card][goreport-img]][goreport] [![Coverage Status][codecov-img]][codecov]

mANSI has some basic definitions for ANSI Escape Codes

## Features

* Defined constants for ANSI escape sequences
* Helper functions for ANSI sequences with parameters
* Support for Linux/BSD, MacOS and Windows 10+

## Usage
Import mansi into your golang project and embed provided ANSI escape sequences into your code output.

```
import "go.melnyk.org/mansi"
```
## Example
```
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
```

## Development Status: In active development
All APIs are in active development and not finalized, and breaking changes will be made in the 0.x series.


[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[license]: https://github.com/mmelnyk/mansi/blob/master/LICENSE
[action-img]: https://github.com/mmelnyk/mansi/workflows/Test/badge.svg
[action]: https://github.com/mmelnyk/mansi/actions
[godoc-img]: https://godoc.org/go.melnyk.org/mansi?status.svg
[godoc]: https://godoc.org/go.melnyk.org/mansi
[goreport-img]: https://goreportcard.com/badge/go.melnyk.org/mansi
[goreport]: https://goreportcard.com/report/go.melnyk.org/mansi
[codecov-img]: https://codecov.io/gh/mmelnyk/mansi/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/mmelnyk/mansi
