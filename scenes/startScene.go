package scenes

import (
	"demo-x/conf"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func StartScene() *tview.Flex {

	statusBox := createBox("角色", []TextPrint{
		{"角色 符十三郎", tview.AlignLeft, conf.NormalColor},
		{"等级 100", tview.AlignLeft, conf.NormalColor},
		{"经验 100", tview.AlignLeft, conf.NormalColor},
		{"生命 100", tview.AlignLeft, conf.NormalColor},
		{"法力 100", tview.AlignLeft, conf.NormalColor},
		{"精力 100", tview.AlignLeft, conf.NormalColor},
	})

	goodBox := createBox("功法", []TextPrint{
		{"心法 九阳神功", tview.AlignLeft, conf.NormalColor},
		{"功法 乾坤大挪移", tview.AlignLeft, conf.NormalColor},
		{"轻功 梯云纵", tview.AlignLeft, conf.NormalColor},
		{"吐纳 搬山诀", tview.AlignLeft, conf.NormalColor},
		{"灵器 翻天印", tview.AlignLeft, conf.NormalColor},
		{"法袍 皂色玄衣", tview.AlignLeft, conf.NormalColor},
	})

	bagBox := createBox("背包", []TextPrint{
		{"a", tview.AlignLeft, conf.NormalColor},
	})
	helpBox := createBox("帮助", []TextPrint{
		{"a 属性面板翻页", tview.AlignLeft, conf.NormalColor},
		{"b 运转功法", tview.AlignLeft, conf.NormalColor},
		{"c 运转功法", tview.AlignLeft, conf.NormalColor},
	})

	mapBox := createBox("地图", []TextPrint{
		{},
		{"  ##########......######", tview.AlignLeft, conf.NormalColor},
		{"  ########...##....#####", tview.AlignLeft, conf.NormalColor},
		{"  #######...#####...####", tview.AlignLeft, conf.NormalColor},
		{"  ######..##.######.####", tview.AlignLeft, conf.NormalColor},
		{"  ######..#...######.###", tview.AlignLeft, conf.NormalColor},
		{"  ######......######.###", tview.AlignLeft, conf.NormalColor},
		{"  #########...##########", tview.AlignLeft, conf.NormalColor},
		{"  #########..###########", tview.AlignLeft, conf.NormalColor},
		{"  #########@############", tview.AlignLeft, conf.NormalColor},
	})

	newsBox := tview.NewBox().SetBorder(true).SetTitle("新闻")
	newsBox.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		tview.Print(screen, "xxx年xxx月xxx日", x+2, y+1, width, tview.AlignLeft, conf.NormalColor)
		tview.Print(screen, "a打了b", x+2, y+2, width, tview.AlignLeft, conf.NormalColor)
		lines := tview.WordWrap("轻功梯云纵是生生世世生生世世生生世世生生世世生生世世生生世世123213123", width-4)
		for i := 0; i < len(lines); i++ {
			tview.Print(screen, lines[i], x+2, y+3+i, width, tview.AlignLeft, conf.NormalColor)
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
