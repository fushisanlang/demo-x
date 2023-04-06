package scenes

import (
	"demo-x/conf"
	"demo-x/data"
	"demo-x/model"
	"demo-x/service"

	"github.com/rivo/tview"
)

func StartScene(startCode int) *tview.Flex {
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := createBox("地图", []model.TextPrint{
		{},
		{Line: "#######", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "###*###", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "###@###", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "###.###", Align: tview.AlignCenter, Color: conf.NormalColor},
		{},
	})
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()

	if startCode > 5 {
		questsBox = createBox("任务", []model.TextPrint{
			{Line: "- 修炼" + data.GoodsMap[0].GoodsName, Align: tview.AlignLeft, Color: conf.BrightColor},
		})
		statusBox = createBox("角色", []model.TextPrint{
			{Line: "角色 丹阳子", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "等级 0%", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "经验 0%", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "生命 100%", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "法力 100%", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "精力 100%", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: " ", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: " ", Align: tview.AlignLeft, Color: conf.NormalColor},

			{Line: "心法 ", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "功法 ", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "轻功 ", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "吐纳 ", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "灵器 ", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "法袍 ", Align: tview.AlignLeft, Color: conf.NormalColor},
		})
		helpBox = createBox("帮助", []model.TextPrint{
			{Line: "F1  帮助菜单", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "F2  打开背包", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "F3  技能列表", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "F4  任务列表", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: " ", Align: tview.AlignLeft, Color: conf.NormalColor},
			//{Line: "F5  任务列表", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "F6  打开地图", Align: tview.AlignLeft, Color: conf.NormalColor},
			//{Line: "F7  任务列表", Align: tview.AlignLeft, Color: conf.NormalColor},
			//{Line: "F8  任务列表", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: " ", Align: tview.AlignLeft, Color: conf.NormalColor},
			//{Line: "F9  关闭对话", Align: tview.AlignLeft, Color: conf.NormalColor},
			//{Line: "F10 挂机界面", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "F11 保存游戏", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "F12 退出游戏 ", Align: tview.AlignLeft, Color: conf.NormalColor},
		})
		newsBox = createBox("世界", []model.TextPrint{
			{Line: "甲子年 甲子月 甲子日", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "泰山-三才宫", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "内门-阅空阁", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: " ", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: " ", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: " ", Align: tview.AlignLeft, Color: conf.NormalColor},

			{Line: "十八日后：", Align: tview.AlignLeft, Color: conf.NormalColor},
			{Line: "将在黑水举办荡魔大会", Align: tview.AlignLeft, Color: conf.NormalColor},
		})
	}
	gameFlex := startScene(startCode)
	flex := baseScene(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)
	return flex
}

func startScene(startCode int) *tview.Flex {

	formatBox := formatBox()
	footBox := footBox()
	box1 := createAlertBox("", []model.TextPrint{

		{Line: "十八年后。。。", Align: tview.AlignCenter, Color: conf.NormalColor},
	})
	box2 := createBox("", []model.TextPrint{
		{},
		{Line: "刘不空:", Align: tview.AlignLeft, Color: conf.BrightColor},
		{Line: "  徒儿，距离我带你回门已有一十八年了。这几年你外功底子打得不错，已经可以修炼内功了。", Align: tview.AlignLeft, Color: conf.NormalColor},
		{},
	})
	box3 := createBox("", []model.TextPrint{
		{},
		{Line: "刘不空:", Align: tview.AlignLeft, Color: conf.BrightColor},
		{Line: "  徒儿，距离我带你回门已有一十八年了。这几年你外功底子打得不错，已经可以修炼内功了。", Align: tview.AlignLeft, Color: conf.NormalColor},
		{Line: "  这是本门的入门内功" + data.GoodsMap[1].GoodsName + "，你从今日就开始修炼吧。", Align: tview.AlignLeft, Color: conf.NormalColor},
		{},
	})
	box4 := createAlertBox("", []model.TextPrint{

		{Line: "获得物品", Align: tview.AlignCenter, Color: conf.NormalColor},

		{Line: data.GoodsMap[0].GoodsName, Align: tview.AlignCenter, Color: conf.ExclusiveColors},
	})
	box5 := createBox("", []model.TextPrint{
		{},
		{Line: "丹阳子:", Align: tview.AlignLeft, Color: conf.ExclusiveColors},
		{Line: "  谢师父。", Align: tview.AlignLeft, Color: conf.NormalColor},
		{},
	})
	box6 := createBox("", []model.TextPrint{
		{},
		{Line: "刘不空:", Align: tview.AlignLeft, Color: conf.BrightColor},
		{Line: "  我近日要去黑水参加荡魔大会。如果" + data.GoodsMap[0].GoodsName + "修炼中有不解之处，可以去找你师兄。", Align: tview.AlignLeft, Color: conf.NormalColor},
		{},
	})

	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(formatBox, 0, 1, false)
	switch startCode {
	case 0:
		gameFlex.AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(formatBox, 0, 1, false).
			AddItem(box1, 0, 1, false).
			AddItem(formatBox, 0, 1, false), 3, 0, false)

	case 1:
		gameFlex.AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(formatBox, 0, 1, false).
			AddItem(box1, 0, 1, false).
			AddItem(formatBox, 0, 1, false), 3, 0, false).
			AddItem(box2, 6, 0, false)
	case 2:
		gameFlex.AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(formatBox, 0, 1, false).
			AddItem(box1, 0, 1, false).
			AddItem(formatBox, 0, 1, false), 3, 0, false).
			AddItem(box3, 7, 0, false)

	case 3:
		gameFlex.AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(formatBox, 0, 1, false).
			AddItem(box1, 0, 1, false).
			AddItem(formatBox, 0, 1, false), 3, 0, false).
			AddItem(box3, 7, 0, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(formatBox, 0, 1, false).
				AddItem(box4, 0, 1, false).
				AddItem(formatBox, 0, 1, false), 4, 0, false)
		service.AddGoods(1, 1)

	case 4:
		gameFlex.AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(formatBox, 0, 1, false).
			AddItem(box1, 0, 1, false).
			AddItem(formatBox, 0, 1, false), 3, 0, false).
			AddItem(box3, 7, 0, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(formatBox, 0, 1, false).
				AddItem(box4, 0, 1, false).
				AddItem(formatBox, 0, 1, false), 4, 0, false).
			AddItem(box5, 6, 0, false)
	case 5:

		gameFlex.AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(formatBox, 0, 1, false).
			AddItem(box1, 0, 1, false).
			AddItem(formatBox, 0, 1, false), 3, 0, false).
			AddItem(box3, 6, 0, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(formatBox, 0, 1, false).
				AddItem(box4, 0, 1, false).
				AddItem(formatBox, 0, 1, false), 4, 0, false).
			AddItem(box5, 6, 0, false).
			AddItem(box6, 6, 0, false)

	default:

		gameFlex.AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(formatBox, 0, 1, false).
			AddItem(box1, 0, 1, false).
			AddItem(formatBox, 0, 1, false), 3, 0, false).
			AddItem(box3, 6, 0, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(formatBox, 0, 1, false).
				AddItem(box4, 0, 1, false).
				AddItem(formatBox, 0, 1, false), 4, 0, false).
			AddItem(box5, 6, 0, false).
			AddItem(box6, 6, 0, false)
		footBox = formatBox

	}
	gameFlex.AddItem(formatBox, 0, 1, false).AddItem(footBox, 5, 0, false)
	return gameFlex
}
