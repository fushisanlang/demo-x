package captures

import (
	"demo-x/scenes"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func StartCaptures(app *tview.Application) *tview.Flex {
	scene1 := scenes.StartScene()

	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'n' || event.Key() == tcell.KeyRight {
			app.SetRoot(StartCaptures2(app), true)
			return nil
		}
		return event
	})
	return scene1

}

func StartCaptures2(app *tview.Application) *tview.Flex {
	scene1 := scenes.StartScene2()

	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'b' || event.Key() == tcell.KeyLeft {
			app.SetRoot(StartCaptures(app), true)
			return nil
		}
		return event
	})
	return scene1

}
