package scenes

import (
	"demo-x/conf"
	"demo-x/model"

	"github.com/rivo/tview"
)

func StoryScene() *tview.Flex {
	formatBox := aFormatBox
	footBox := aFootBox
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := tview.NewBox()
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()
	gameBox := createBoxWithoutSider([]model.TextPrintStruct{
		{},
		{Line: "嘉德十八年，景帝痴迷长生之术，听信国师何继谗言。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "在全国抓捕三千名童男童女准备炼制不老丹。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "此令一出，举国哗然。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "又恰逢黄河决堤，南方大旱。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "民间传言景帝自断气数，董家江山命不久矣。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: " ", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "时任冀州将军那河离主张清君侧斩妖道起兵造反。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "只用了十余日便兵临京畿。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "奈何妖道何继在城门外设下五行三刹阵，将冀州军硬生生困在城下七天，军心大乱。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "三才宗推算出凉朝气数已尽，便派出当代行走刘不空下山祝那河离铲除妖道。", Align: tview.AlignCenter, Color: conf.NormalColor},

		{Line: "刘不空道法高深，轻松地破了何继的大阵。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "还用搬山填海之术治理了黄河的水患和南方的旱灾。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: " ", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "就在他回宗门复命的途中，于泰山脚下遇到一名弃婴。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "这婴儿双目闪烁，颅顶放光，乃是万中无一的修真奇才。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "便将其带回门中，取名丹阳子，收为亲传弟子。", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: " ", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: " ", Align: tview.AlignCenter, Color: conf.NormalColor},
		//{"按 \"n\" 键或 \"回车键\" 翻页。", tview.AlignCenter, conf.GuideColor},
	})
	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formatBox, 0, 1, false).
		AddItem(gameBox, 30, 0, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(footBox, 5, 0, false)
	flex := baseScene(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)

	return flex

}
