package verbose

import "github.com/fatih/color"

var ErrorColor = color.New(color.BgHiRed, color.FgHiWhite)
var InfoColor = color.New(color.BgHiBlue, color.FgHiWhite)
var WarnColor = color.New(color.BgHiYellow, color.FgBlack)
var SuccessColor = color.New(color.BgHiGreen, color.FgBlack)
var NeutralColor = color.New(color.BgHiWhite, color.FgBlack)
var DebugColor = color.New(color.BgHiCyan, color.FgBlack)

var ErrorTrueColor = color.New().AddBgRGB(252, 62, 53).AddRGB(0, 0, 0)
var InfoTrueColor = color.New().AddBgRGB(69, 165, 249).AddRGB(0, 0, 0)
var WarnTrueColor = color.New().AddBgRGB(255, 199, 33).AddRGB(0, 0, 0)
var SuccessTrueColor = color.New().AddBgRGB(39, 196, 39).AddRGB(0, 0, 0)
var NeutralTrueColor = color.New().AddBgRGB(242, 242, 242).AddRGB(0, 0, 0)
var DebugTrueColor = color.New().AddBgRGB(36, 226, 166).AddRGB(0, 0, 0)
