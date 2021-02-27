package main

import "github.com/rivo/tview"

func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Search Word"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Sort or Filter"), 0, 2, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("file list"), 0, 7, false), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("preview"), 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
