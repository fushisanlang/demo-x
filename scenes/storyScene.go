package scenes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func StoryScene() *tview.Flex {

	statusBox := tview.NewBox()
	goodBox := tview.NewBox()
	bagBox := tview.NewBox()
	mapBox := tview.NewBox()
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()
	gameBox := createBox("背景", []TextPrint{
		{},
		{"嘉德十八年，景帝痴迷长生之术，听信国师何继谗言。", tview.AlignCenter, tcell.ColorWhite},
		{"在全国抓捕三千名童男童女准备炼制不老丹。", tview.AlignCenter, tcell.ColorWhite},
		{"此令一出，举国哗然。", tview.AlignCenter, tcell.ColorWhite},
		{"又恰逢黄河决堤，南方大旱。", tview.AlignCenter, tcell.ColorWhite},
		{"民间传言景帝自断气数，董家江山命不久矣。", tview.AlignCenter, tcell.ColorWhite},
		{" ", tview.AlignCenter, tcell.ColorWhite},
		{"时任冀州将军那河离主张清君侧斩妖道起兵造反。", tview.AlignCenter, tcell.ColorWhite},
		{"只用了十余日便兵临京畿。", tview.AlignCenter, tcell.ColorWhite},
		{"奈何妖道何继在城门外设下五行三刹阵，将冀州军硬生生困在城下七天，军心大乱。", tview.AlignCenter, tcell.ColorWhite},
		{"三才宗推算出凉朝气数已尽，便派出当代行走刘不空下山祝那河离铲除妖道。", tview.AlignCenter, tcell.ColorWhite},

		{"刘不空道法高深，轻松地破了何继的大阵。", tview.AlignCenter, tcell.ColorWhite},
		{"还用了搬山填海之术治理了黄河的水患和南方的旱灾。", tview.AlignCenter, tcell.ColorWhite},
		{" ", tview.AlignCenter, tcell.ColorWhite},
		{"就在他回宗门复命的途中，于泰山脚下遇到一名弃婴。", tview.AlignCenter, tcell.ColorWhite},
		{"这婴儿双目闪烁，颅顶放光，乃是万中无一的修真奇才。", tview.AlignCenter, tcell.ColorWhite},
		{"便将其带回门中，收为弟子。", tview.AlignCenter, tcell.ColorWhite},
		{" ", tview.AlignCenter, tcell.ColorWhite},
		{" ", tview.AlignCenter, tcell.ColorWhite},
		{"按 \"n\" 键或 \"回车键\" 翻页。", tview.AlignCenter, tcell.ColorGreenYellow},
	})
	flex := baseScene(statusBox, goodBox, bagBox, mapBox, helpBox, newsBox, gameBox)
	return flex

}
