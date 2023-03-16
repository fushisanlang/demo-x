package scenes

import (
	"demo-x/conf"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func VerifyScene() *tview.Flex {
	GameName := conf.GameName
	statusBox := tview.NewBox().SetBorder(false).SetTitle("用户需知")
	statusBox.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		tview.Print(screen, "玩家您好", x, y+1, width, tview.AlignCenter, tcell.ColorWhite)
		tview.Print(screen, "你正在打开的游戏", x, y+2, width, tview.AlignCenter, tcell.ColorWhite)
		tview.Print(screen, GameName, x, y+3, width, tview.AlignCenter, tcell.ColorRed)
		tview.Print(screen, "是一款基于tui的游戏", x, y+4, width, tview.AlignCenter, tcell.ColorWhite)
		tview.Print(screen, "由于技术原因", x, y+6, width, tview.AlignCenter, tcell.ColorWhite)
		tview.Print(screen, "需要您手动调节窗口大小", x, y+7, width, tview.AlignCenter, tcell.ColorGreen)
		tview.Print(screen, "来保证您的游戏体验", x, y+8, width, tview.AlignCenter, tcell.ColorWhite)

		tview.Print(screen, "请尝试改变窗口大小", x, y+10, width, tview.AlignCenter, tcell.ColorWhite)
		tview.Print(screen, "直到正好看不到“多余窗口”", x, y+11, width, tview.AlignCenter, tcell.ColorYellow)

		tview.Print(screen, "如果您发现这段文字没有左右居中", x, y+13, width, tview.AlignCenter, tcell.ColorGray)
		tview.Print(screen, "可能是窗口较小", x, y+14, width, tview.AlignCenter, tcell.ColorGray)
		tview.Print(screen, "请尝试扩大窗口", x, y+15, width, tview.AlignCenter, tcell.ColorGreen)

		tview.Print(screen, "在调整完成后", x, y+17, width, tview.AlignCenter, tcell.ColorWhite)
		tview.Print(screen, "点击回车开始游戏", x, y+18, width, tview.AlignCenter, tcell.ColorGreenYellow)

		tview.Print(screen, "在游戏中也可以随时调整窗口大小", x, y+20, width, tview.AlignCenter, tcell.ColorWhite)
		tview.Print(screen, "祝您游戏愉快", x, y+21, width, tview.AlignCenter, tcell.ColorWhite)

		return x, y, width, height
	})
	goodBox := tview.NewBox().SetBorder(true).SetTitle("多余窗口")

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(statusBox, 50, 0, false).
			AddItem(goodBox, 0, 1, true), 160, 0, false).
		AddItem(goodBox, 0, 1, true)
	return flex

}
