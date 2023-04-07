package conf

import "github.com/gdamore/tcell/v2"

var GameName string

var NormalColor tcell.Color
var AlertColor tcell.Color
var SecondColor tcell.Color
var BrightColor tcell.Color
var GuideColor tcell.Color
var ExclusiveColors tcell.Color
var SaveFileList []string
var SaveBakDir string
var SqlFile string
var SaveFile string

func init() {
	GameName = "demo-x"
	SaveBakDir = "data_bak"
	NormalColor = tcell.ColorWhite       //普通显示
	AlertColor = tcell.ColorOrangeRed    //警告显示
	SecondColor = tcell.ColorGray        //非必要显示
	BrightColor = tcell.ColorGreenYellow //高亮显示
	GuideColor = tcell.ColorForestGreen  //操作指引
	ExclusiveColors = tcell.ColorMediumSeaGreen
	SaveFileList = []string{"bag.bin"}
	SaveFile = "data/save.bat"
}
