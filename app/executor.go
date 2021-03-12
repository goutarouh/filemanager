package app

import (
	"filemanager/component"
	"filemanager/constant"
	"filemanager/model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// ExecuteApp is the function that execute TUI app with tview and
// where I set primitives to the screen(tview.Grid)
func ExecuteApp() {

	updateTableContents := make(chan *model.TableContentsInfo)
	updatePreview := make(chan *model.UpdatePreview)
	updateCellBackground := make(chan *model.FileRow)

	defer func() {
		close(updateTableContents)
		close(updatePreview)
		close(updateCellBackground)
	}()

	eventChannel := model.EventChannel{
		UpdateTableContents:  updateTableContents,
		UpdatePreview:        updatePreview,
		UpdateCellBackground: updateCellBackground,
	}

	grid := tview.NewGrid().SetSize(10, 2, 0, 0)

	searchInputField := component.CreateSearchInputText(eventChannel)
	table := component.CreateTable(eventChannel)
	description := component.CreateDescription()
	preview := component.CreatePreview()

	grid.AddItem(searchInputField, 0, 0, 2, 1, 0, 0, false)
	grid.AddItem(table.Table, 2, 0, 10, 1, 0, 0, false)
	grid.AddItem(description, 0, 1, 2, 2, 0, 0, false)
	grid.AddItem(preview.TextView, 2, 1, 10, 2, 0, 0, false)

	app := tview.NewApplication()

	go table.EventReceiver(app, eventChannel)
	go preview.EventReceiver(app, eventChannel)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 's' {
			app.SetFocus(searchInputField)
			return nil
		}
		if event.Key() == constant.KeyUp || event.Key() == constant.KeyDown {
			app.SetFocus(table.Table)
			return event
		}
		return event
	})

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
