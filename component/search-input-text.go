package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CreateSearchInputText は検索の箱を作成する。
func CreateSearchInputText(searchInputTextChannel chan string) *tview.InputField {
	searchInputField := tview.NewInputField()
	searchInputField.SetTitle("search")
	searchInputField.SetBorder(true)
	searchInputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			searchInputTextChannel <- searchInputField.GetText()
			return nil
		}
		return event
	})

	return searchInputField
}
