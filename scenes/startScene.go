package scenes

import (
	"github.com/rivo/tview"
)

func StartScene() *tview.Flex {
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	return flex
}

func StartScene2() *tview.Flex {
	scene1 := tview.NewFlex()
	text2 := tview.NewTextView().SetText("This is Scene 2")
	scene1.AddItem(text2, 0, 1, false)
	return scene1
}
