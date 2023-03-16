package scenes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func createBox(title string, content []TextPrint) *tview.Box {
	box := tview.NewBox().SetBorder(true).SetTitle(title)
	box.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		for i, line := range content {
			tview.Print(screen, line.Line, x+2, y+i+1, width, line.Align, line.Color)
		}
		return x, y, width, height
	})
	return box
}

func StartScene() *tview.Flex {

	statusBox := createBox("角色", []TextPrint{
		{"角色 符十三郎", tview.AlignLeft, tcell.ColorGreen},
		{"等级 100", tview.AlignLeft, tcell.ColorWhite},
		{"经验 100", tview.AlignLeft, tcell.ColorWhite},
		{"生命 100", tview.AlignLeft, tcell.ColorWhite},
		{"法力 100", tview.AlignLeft, tcell.ColorWhite},
		{"精力 100", tview.AlignLeft, tcell.ColorWhite},
	})

	goodBox := createBox("功法", []TextPrint{
		{"心法 九阳神功", tview.AlignLeft, tcell.ColorWhite},
		{"功法 乾坤大挪移", tview.AlignLeft, tcell.ColorWhite},
		{"轻功 梯云纵", tview.AlignLeft, tcell.ColorWhite},
		{"吐纳 搬山诀", tview.AlignLeft, tcell.ColorWhite},
		{"灵器 翻天印", tview.AlignLeft, tcell.ColorWhite},
		{"法袍 皂色玄衣", tview.AlignLeft, tcell.ColorWhite},
	})

	bagBox := createBox("背包", []TextPrint{
		{"a", tview.AlignLeft, tcell.ColorWhite},
	})
	helpBox := createBox("帮助", []TextPrint{
		{"a 属性面板翻页", tview.AlignLeft, tcell.ColorWhite},
		{"b 运转功法", tview.AlignLeft, tcell.ColorWhite},
		{"c 运转功法", tview.AlignLeft, tcell.ColorWhite},
	})

	mapBox := createBox("地图", []TextPrint{
		{},
		{"  ##########......######", tview.AlignLeft, tcell.ColorWhite},
		{"  ########...##....#####", tview.AlignLeft, tcell.ColorWhite},
		{"  #######...#####...####", tview.AlignLeft, tcell.ColorWhite},
		{"  ######..##.######.####", tview.AlignLeft, tcell.ColorWhite},
		{"  ######..#...######.###", tview.AlignLeft, tcell.ColorWhite},
		{"  ######......######.###", tview.AlignLeft, tcell.ColorWhite},
		{"  #########...##########", tview.AlignLeft, tcell.ColorWhite},
		{"  #########..###########", tview.AlignLeft, tcell.ColorWhite},
		{"  #########@############", tview.AlignLeft, tcell.ColorWhite},
	})

	newsBox := tview.NewBox().SetBorder(true).SetTitle("新闻")
	newsBox.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		tview.Print(screen, "xxx年xxx月xxx日", x+2, y+1, width, tview.AlignLeft, tcell.ColorWhite)
		tview.Print(screen, "a打了b", x+2, y+2, width, tview.AlignLeft, tcell.ColorWhite)
		lines := tview.WordWrap("轻功梯云纵是生生世世生生世世生生世世生生世世生生世世生生世世123213123", width-4)
		for i := 0; i < len(lines); i++ {
			tview.Print(screen, lines[i], x+2, y+3+i, width, tview.AlignLeft, tcell.ColorWhite)
		}
		return x, y, width, height
	})

	gameBox := tview.NewBox().SetBorder(true).SetTitle("游戏")
	gameBox.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {

		return x, y, width, height
	})
	flex := baseScene(statusBox, goodBox, bagBox, mapBox, helpBox, newsBox, gameBox)
	return flex
}

func StartScene2() *tview.Flex {
	scene1 := tview.NewFlex()
	text2 := tview.NewTextView().SetText("This is Scene 2")
	scene1.AddItem(text2, 0, 1, false)
	return scene1
}
