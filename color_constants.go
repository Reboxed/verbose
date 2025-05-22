package verbose

import "github.com/fatih/color"

var errorColor = color.New(color.BgHiRed, color.FgHiWhite)
var infoColor = color.New(color.BgHiBlue, color.FgHiWhite)
var warnColor = color.New(color.BgHiYellow, color.FgHiWhite)
var successColor = color.New(color.FgBlack).AddBgRGB(39, 196, 39)
var neutralColor = color.New(color.BgWhite, color.FgHiWhite)
