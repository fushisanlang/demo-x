package scenes

import (
	"demo-x/conf"
	"demo-x/service"

	"github.com/gogf/gf/util/gconv"
	"github.com/rivo/tview"
)

func BagScene() *tview.Flex {
	formatBox := formatBox()
	footBox := footBox()
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := tview.NewBox()
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()

	gameBox := tview.NewTable().SetBorders(false)
	bags := service.Bags

	cell := tview.NewTableCell("               ")
	cell.SetExpansion(1)
	gameBox.SetCell(0, 0, cell)
	gameBox.SetCell(0, 1, cell)
	gameBox.SetCell(0, 2, cell)
	gameBox.SetCell(0, 3, cell)
	gameBox.SetCell(0, 4, cell)
	gameBox.SetCell(0, 5, cell)
	gameBox.SetCell(0, 6, cell)
	gameBox.SetCell(0, 7, cell)
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			// Calculate the index into the data array.
			index := row*8 + col

			// If we've reached the end of the data, break out of the loop.
			if index >= 64 {
				break
			}

			// Add a new cell to the table.
			cellString := ""
			if bags.Bag[index].GoodID != 0 {
				cellString = service.NewGoodsSlice[bags.Bag[index].GoodID].GoodsName + " *" + gconv.String(bags.Bag[index].Count)
			}
			cell = tview.NewTableCell(cellString).
				SetTextColor(conf.NormalColor).
				SetAlign(tview.AlignCenter)

			gameBox.SetCell(row+1, col, cell)

		}
	}

	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formatBox, 0, 1, false).
		AddItem(gameBox, 30, 0, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(footBox, 5, 0, false)
	flex := baseSceneWithSide(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)

	return flex

}
