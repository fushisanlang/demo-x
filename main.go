package main

import (
	"demo-x/captures"
	"demo-x/tools"
	"fmt"
	"sort"

	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	startScene := captures.VerifyCapture(app)
	//@todo: 需要判断是否有存档

	app.SetRoot(startScene, true)
	//// Run the application.
	err := app.Run()
	tools.CheckErr(err)

}
func main2() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 999, 999999}
	d := sort.Search(len(a), func(i int) bool { return a[i] >= 99 })
	fmt.Println(d)
}
