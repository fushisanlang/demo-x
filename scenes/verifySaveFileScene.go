package scenes

import (
	"demo-x/conf"
	"demo-x/model"

	"github.com/rivo/tview"
)

func VeriftSaveFileScene() *tview.Flex {
	formatBox := formatBox()
	footBox := footBox()
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := tview.NewBox()
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()
	gameBox := createBoxWithoutSider([]model.TextPrint{
		{},
		{Line: "正在检查存档，请稍候。。。", Align: tview.AlignCenter, Color: conf.BrightColor},
		{},
	})
	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formatBox, 0, 1, false).
		AddItem(gameBox, 5, 0, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(footBox, 5, 0, false)
	flex := baseScene(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)

	return flex
}
func VeriftSaveFileSuccessScene() *tview.Flex {
	formatBox := formatBox()
	footBox := footBox()
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := tview.NewBox()
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()
	gameBox := createBoxWithoutSider([]model.TextPrint{
		{},
		{Line: "存档加载完毕", Align: tview.AlignCenter, Color: conf.BrightColor},
		{},
	})

	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formatBox, 0, 1, false).
		AddItem(gameBox, 5, 0, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(footBox, 5, 0, false)
	flex := baseScene(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)

	return flex
}
func VeriftSaveFileFailScene() *tview.Flex {
	formatBox := formatBox()
	footBox := footBox()
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := tview.NewBox()
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()
	gameBox := createBoxWithoutSider([]model.TextPrint{
		{},
		{Line: "未找到存档或存档异常，新建存档中。。。", Align: tview.AlignCenter, Color: conf.BrightColor},
		{Line: "如欲修复存档，请在" + conf.SaveBakDir + "文件夹中查找存档备份。", Align: tview.AlignCenter, Color: conf.BrightColor},
		{Line: "可以联系开发者获取帮助。", Align: tview.AlignCenter, Color: conf.BrightColor},

		{},
	})
	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formatBox, 0, 1, false).
		AddItem(gameBox, 7, 0, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(footBox, 5, 0, false)
	flex := baseScene(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)

	return flex
}
