package verbose

import "github.com/fatih/color"

var ErrorColor = color.New(color.BgHiRed, color.FgHiWhite)
var InfoColor = color.New(color.BgHiBlue, color.FgHiWhite)
var WarnColor = color.New(color.BgHiYellow, color.FgHiWhite)
var SuccessColor = color.New(color.FgBlack).AddBgRGB(39, 196, 39)
var NeutralColor = color.New(color.BgWhite, color.FgHiWhite)

var Error = "ERR"
var Info = "INFO"
var Warn = "WARN"
var Success = "OK"
var Print = "PRINT"
