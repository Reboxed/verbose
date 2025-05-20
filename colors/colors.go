package colors

import "github.com/fatih/color"

var Error = color.New(color.BgHiRed, color.FgHiWhite)
var Info = color.New(color.BgHiBlue, color.FgHiWhite)
var Warn = color.New(color.BgHiYellow, color.FgHiWhite)
var Success = color.New(color.FgBlack).AddBgRGB(39, 196, 39)
var Neutral = color.New(color.BgWhite, color.FgHiWhite)
