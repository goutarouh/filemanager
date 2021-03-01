package main

import (
	"filemanager/component"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	searchInputTextChannel := make(chan string)

	//main flex
	flex := tview.NewFlex().SetDirection(tview.FlexColumn)

	// left side of screen
	searchInputField := component.CreateSearchInputText(searchInputTextChannel)
	table := component.CreateTable()

	go table.EventReceiver(searchInputTextChannel)

	flex.AddItem(
		tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(searchInputField, 0, 2, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("test2"), 0, 2, false).
			AddItem(table.Table, 0, 6, false),
		0, 1, false)

	// right side of screen
	flex.AddItem(tview.NewBox().SetBorder(true).SetTitle("preview"), 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
