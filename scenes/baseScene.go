package scenes

import (
	"demo-x/conf"
	"demo-x/model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func createAlertBox(title string, content []model.TextPrintStruct) *tview.Box {
	box := tview.NewBox().SetTitle(title).SetBackgroundColor(tcell.ColorDarkGray)
	box.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		for i, line := range content {
			tview.Print(screen, line.Line, x+2, y+i+1, width, line.Align, line.Color)
		}
		return x, y, width, height
	})
	return box
}
func createBoxWithoutSider(content []model.TextPrintStruct) *tview.Box {
	box := tview.NewBox().SetBorder(false)
	box.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		for i, line := range content {
			tview.Print(screen, line.Line, x+2, y+i+1, width, line.Align, line.Color)
		}
		return x, y, width, height
	})
	return box
}
func createBox(title string, content []model.TextPrintStruct) *tview.Box {
	box := tview.NewBox().SetBorder(true).SetTitle(title)
	box.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		for i, line := range content {
			tview.Print(screen, line.Line, x+2, y+i+1, width, line.Align, line.Color)
		}
		return x, y, width, height
	})
	return box
}
func formatBox() *tview.Box {
	//用来创建一个空box，占位使用，确保显示内容上下剧中
	formatBox := createBoxWithoutSider([]model.TextPrintStruct{
		{},
	})
	return formatBox
}
func footBox() *tview.Box {
	footBox := createBoxWithoutSider([]model.TextPrintStruct{
		{},
		{Line: "点击 \"回车\" 键继续", Align: tview.AlignCenter, Color: conf.GuideColor},
		{},
	})
	return footBox
}
func footBoxEsc() *tview.Box {
	footBox := createBoxWithoutSider([]model.TextPrintStruct{
		{},
		{Line: "点击 \"esc\" 键返回", Align: tview.AlignCenter, Color: conf.GuideColor},
		{},
	})
	return footBox
}
func baseSceneWithSide(statusBox, questsBox, mapBox, helpBox, newsBox *tview.Box, gameFlex *tview.Flex) *tview.Flex {
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(statusBox, 16, 0, false).
			//AddItem(questsBox, 8, 0, false).
			AddItem(questsBox, 0, 1, true), 10, 0, false).
		AddItem(gameFlex, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(mapBox, 13, 0, false).
			AddItem(helpBox, 16, 0, false).
			AddItem(newsBox, 0, 1, true), 10, 0, false)

	return flex
}
func baseScene(statusBox, questsBox, mapBox, helpBox, newsBox *tview.Box, gameFlex *tview.Flex) *tview.Flex {

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(statusBox, 16, 0, false).
			//AddItem(questsBox, 8, 0, false).
			AddItem(questsBox, 0, 1, true), 30, 0, false).
		AddItem(gameFlex, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(mapBox, 13, 0, false).
			AddItem(helpBox, 16, 0, false).
			AddItem(newsBox, 0, 1, true), 30, 0, false)

	return flex
}
