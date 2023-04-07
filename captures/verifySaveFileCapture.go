package captures

import (
	"demo-x/scenes"
	"demo-x/service"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func VerifySaveFileCapture(app *tview.Application) *tview.Flex {
	scene1 := scenes.VeriftSaveFileScene()
	time.Sleep(3 * time.Second)

	app.SetRoot(VerifySaveFileSuccessCapture(app), true)

	return scene1
}

func VerifySaveFileSuccessCapture(app *tview.Application) *tview.Flex {
	scene1 := scenes.VeriftSaveFileSuccessScene()
	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			app.SetRoot(CaptureMap[service.GameData.UserSaveId](app), true)
			return nil
		}
		return event
	})
	return scene1

}
func VerifySaveFileFailCapture(app *tview.Application) *tview.Flex {
	scene1 := scenes.VeriftSaveFileFailScene()
	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			service.InitData()
			app.SetRoot(StoryCapture(app), true)
			return nil
		}
		return event
	})
	return scene1

}
