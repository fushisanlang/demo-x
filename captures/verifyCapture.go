package captures

import (
	"demo-x/scenes"
	"demo-x/service"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func VerifyCapture(app *tview.Application) *tview.Flex {
	scene1 := scenes.VerifyScene()

	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			status := service.VeriftSave()
			if status {
				service.LoadData()
				app.SetRoot(VerifySaveFileSuccessCapture(app), true)
			} else {
				app.SetRoot(VerifySaveFileFailCapture(app), true)

			}

			return nil
		}
		return event
	})
	return scene1

}
