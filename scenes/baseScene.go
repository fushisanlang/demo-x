package scenes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func baseScene(statusBox, goodBox, bagBox, mapBox, helpBox, newsBox, gameBox *tview.Box) *tview.Flex {

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(statusBox, 8, 0, false).
			AddItem(goodBox, 8, 0, false).
			AddItem(bagBox, 0, 1, true), 30, 0, false).
		AddItem(gameBox, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(mapBox, 13, 0, false).
			AddItem(helpBox, 15, 0, false).
			AddItem(newsBox, 0, 1, true), 30, 0, false)

	return flex
}

type TextPrint struct {
	Line  string
	Align int
	Color tcell.Color
}
