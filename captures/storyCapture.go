package captures

import (
	"demo-x/scenes"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func StoryCapture(app *tview.Application) *tview.Flex {
	scene1 := scenes.StoryScene()

	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter || event.Rune() == 'n' {
			app.SetRoot(StartCapture(app), true)
			return nil
		}
		return event
	})
	return scene1

}
