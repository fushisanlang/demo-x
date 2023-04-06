package main

import (
	"demo-x/model"
	"demo-x/service"
	"github.com/rivo/tview"
	"strconv"
)

type CaseStruct struct {
	ID     int
	GoodID int
	Count  int
}

type BagsStruct struct {
	Bag   [64]CaseStruct
	UseID int
}

func main() {
	// Create the application and the table.
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)

	// Create the bags data.
	bags := service.Bags
	for i := 0; i < 64; i++ {
		bags[i] = model.CaseStruct(CaseStruct{ID: i, GoodID: i % 10, Count: i % 5})
	}

	// Add the bags data to the table.
	for row := 0; row < 8; row++ {
		for col := 0; col < 16; col += 2 {
			// Calculate the index into the bags array.
			index := row*16 + col/2

			// If we've reached the end of the bags data, break out of the loop.
			if index >= len(bags) {
				break
			}

			// Add a new cell to the table.
			goodID := tview.NewTableCell(
				strconv.Itoa(bags[index].GoodID)).
				SetTextColor(tview.Styles.PrimaryTextColor).
				SetAlign(tview.AlignCenter)
			count := tview.NewTableCell(
				strconv.Itoa(bags[index].Count)).
				SetTextColor(tview.Styles.PrimaryTextColor).
				SetAlign(tview.AlignCenter)
			table.SetCell(row, col, goodID)
			table.SetCell(row, col+1, count)
		}
	}

	// Set the table as the root component of the application.
	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
