package scenes

import (
	"demo-x/conf"
	"github.com/rivo/tview"
)

func HelpScene() *tview.Flex {

	formatBox := aFormatBox
	footBox := aFootBox
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := tview.NewBox()

	table := tview.NewTable().SetBorders(false)

	cells := [][]string{
		{"快捷按键", "战斗指令", "移动指令"},
		{" ", "", ""},
		{"F1  帮助菜单", "q  技能1", "w  向上移动"},
		{"F2  打开背包", "w  技能2", "a  向左移动"},
		{"F3  技能列表", "e  技能3", "s  向下移动"},
		{"F4  任务列表", "r  技能4", "d  向右移动"},
		{" ", "", ""},
		{"F5  打开地图", "1  道具1", "j  返回据点"},
		{"F6  打开地图", "2  道具2", "i  交互场景"},
		{"F7  自动操作", "3  道具3", "k  关闭自动"},
		{"F8  关闭自动", "4  道具4", "l         "},
		{" ", "", ""},
		{"F9  关闭对话", "z  跳过战斗", "b  返回据点"},
		{"F10 挂机界面", "x  跳过结算", "n  快速对话"},
		{"F11 保存游戏", "c  逃离战斗", "m         "},
		{"F12 退出游戏", "v  重新开始", "<         "},
	}
	for row, rowCells := range cells {
		for col, cellText := range rowCells {
			cell := tview.NewTableCell(cellText).
				SetTextColor(conf.NormalColor).
				SetAlign(tview.AlignCenter)

			if col == 1 {
				cell.SetExpansion(1)
			} else {
				cell.SetExpansion(0)
			}

			table.SetCell(row, col, cell)
		}
	}
	newsBox := tview.NewBox()
	helpBox := tview.NewBox()
	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formatBox, 0, 1, false).
		AddItem(table, 30, 0, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(footBox, 5, 0, false)
	flex := baseSceneWithSide(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)

	return flex

}
