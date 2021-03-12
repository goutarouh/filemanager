package component

import (
	"filemanager/model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CreateSearchInputText は検索の箱を作成する。
func CreateSearchInputText(eventChannel model.EventChannel) *tview.InputField {
	searchInputField := tview.NewInputField()
	searchInputField.SetLabel("search: ")
	searchInputField.SetBorder(true)

	searchInputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			tableContentsInfo := &model.TableContentsInfo{
				Word:            searchInputField.GetText(),
				Path:            "",
				TableUpdateType: model.SearchInputText,
			}
			eventChannel.UpdateTableContents <- tableContentsInfo
			return nil
		}
		return event
	})

	return searchInputField
}
