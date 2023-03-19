package scenes

import (
	"demo-x/conf"
	"demo-x/model"

	"github.com/rivo/tview"
)

func VerifyScene() *tview.Flex {
	statusBox := createBoxWithoutSider([]model.TextPrint{{},
		{Line: "玩家您好", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "你正在打开的游戏", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: conf.GameName, Align: tview.AlignCenter, Color: conf.BrightColor},
		{Line: "是一款基于tui的游戏", Align: tview.AlignCenter, Color: conf.NormalColor}, {},
		{Line: "由于技术原因", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "需要您手动调节窗口大小", Align: tview.AlignCenter, Color: conf.GuideColor},
		{Line: "来保证您的游戏体验", Align: tview.AlignCenter, Color: conf.NormalColor}, {},
		{Line: "请尝试改变窗口大小", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "直到正好看不到“多余窗口”", Align: tview.AlignCenter, Color: conf.GuideColor}, {},
		{Line: "如果您发现这段文字没有左右居中", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "可能是窗口较小", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "请尝试扩大窗口", Align: tview.AlignCenter, Color: conf.GuideColor}, {},
		{Line: "如果您发现汉字显示异常或者想调整字体", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "请调整窗口的字体显示或者给开发者留言", Align: tview.AlignCenter, Color: conf.GuideColor},
		{Line: "If you find that the Chinese characters are displayed abnormally", Align: tview.AlignCenter, Color: conf.SecondColor},
		{Line: "Please adjust the font display of the window or leave a message to the developer", Align: tview.AlignCenter, Color: conf.SecondColor},
		{Line: "在游戏中也可以随时调整", Align: tview.AlignCenter, Color: conf.SecondColor}, {}, {}, {}, {}, {},
		{Line: "邮箱/mail ", Align: tview.AlignCenter, Color: conf.NormalColor},
		{Line: "13@fushisanlang.cn", Align: tview.AlignCenter, Color: conf.ExclusiveColors}, {}, {}, {},

		{Line: "点击 回车 进入游戏", Align: tview.AlignCenter, Color: conf.GuideColor}, {},

		{Line: "祝您游戏愉快", Align: tview.AlignCenter, Color: conf.BrightColor},
	})

	goodBox := tview.NewBox().SetBorder(true).SetTitle("多余窗口")

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(statusBox, 50, 0, false).
			AddItem(goodBox, 0, 1, true), 160, 0, false).
		AddItem(goodBox, 0, 1, true)
	return flex

}
