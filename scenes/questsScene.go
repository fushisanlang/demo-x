package scenes

import (
	"github.com/rivo/tview"
)

func QuestsScene() *tview.Flex {

	formatBox := formatBox()
	footBox := footBox()
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := tview.NewBox()
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()
	gameBox := tview.NewBox()
	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formatBox, 0, 1, false).
		AddItem(gameBox, 30, 0, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(footBox, 5, 0, false)
	flex := baseScene(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)

	return flex

}
