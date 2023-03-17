package conf

import "github.com/gdamore/tcell/v2"

var GameName string

var NormalColor tcell.Color
var AlertColor tcell.Color
var SecondColor tcell.Color
var BrightColor tcell.Color
var GuideColor tcell.Color
var ExclusiveColors tcell.Color

func init() {
	GameName = "demo-x"
	NormalColor = tcell.ColorWhite
	AlertColor = tcell.ColorOrangeRed
	SecondColor = tcell.ColorGray
	BrightColor = tcell.ColorGreenYellow
	GuideColor = tcell.ColorForestGreen
	ExclusiveColors = tcell.ColorMediumSeaGreen
}
