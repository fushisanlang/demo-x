package scenes

import "github.com/rivo/tview"

var aFormatBox *tview.Box
var aFootBox *tview.Box

func init() {
	aFormatBox = formatBox()
	aFootBox = footBox()
}
