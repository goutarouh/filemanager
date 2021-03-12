package component

import (
	"filemanager/constant"
	"filemanager/data"
	"filemanager/log"
	"filemanager/model"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// FileTable contains the data that is used for file-table.
type FileTable struct {
	Table             *tview.Table
	path              string
	files             []os.FileInfo
	item              []os.FileInfo
	beforeSelectedRow int
	selectedFileRow   *model.FileRow
}

// CreateTable makes *FileTable which struct is used for file-table-area.
func CreateTable(eventChannel model.EventChannel) *FileTable {
	table := tview.NewTable()
	table.SetTitle("file list")
	table.SetBorder(true)

	initDir, err := os.Getwd()
	if err != nil {
		log.Log.Fatal("error at getting current directory")
	}

	files := data.ReadAllFileName(initDir)

	fileRow := &model.FileRow{
		FileInfo: nil,
		Row:      constant.ItemNotSelected,
	}

	fileTable := &FileTable{
		Table:             table,
		path:              initDir,
		files:             files,
		item:              files,
		beforeSelectedRow: constant.ItemNotSelected,
		selectedFileRow:   fileRow,
	}

	fileTable.Table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		oldRow := fileTable.beforeSelectedRow
		switch event.Key() {
		case constant.KeyUp:
			var fileRow *model.FileRow
			if oldRow == constant.ItemNotSelected {
				fileInfo := fileTable.item[len(fileTable.item)-1]
				fileRow = &model.FileRow{
					FileInfo: fileInfo,
					Row:      len(fileTable.item) - 1,
				}
			} else if oldRow > 0 {
				fileInfo := fileTable.item[oldRow-1]
				fileRow = &model.FileRow{
					FileInfo: fileInfo,
					Row:      oldRow - 1,
				}
			} else {
				return nil
			}
			updatePreview := &model.UpdatePreview{
				FileRow: fileRow,
				AbsPath: filepath.Join(fileTable.path, fileRow.FileInfo.Name()),
			}
			eventChannel.UpdatePreview <- updatePreview
			eventChannel.UpdateCellBackground <- fileRow
			return nil
		case constant.KeyDown:
			var fileRow *model.FileRow
			if oldRow == constant.ItemNotSelected {
				fileRow = &model.FileRow{
					FileInfo: fileTable.item[0],
					Row:      0,
				}
			} else if oldRow < (len(fileTable.item) - 1) {
				fileRow = &model.FileRow{
					FileInfo: fileTable.item[oldRow+1],
					Row:      oldRow + 1,
				}
			} else {
				return nil
			}
			updatePreview := &model.UpdatePreview{
				FileRow: fileRow,
				AbsPath: filepath.Join(fileTable.path, fileRow.FileInfo.Name()),
			}
			eventChannel.UpdatePreview <- updatePreview
			eventChannel.UpdateCellBackground <- fileRow
			return nil
		case constant.KeyLeft:
			tableContentsInfo := &model.TableContentsInfo{
				Word:            "",
				Path:            filepath.Dir(fileTable.path),
				TableUpdateType: model.LeftRightPressed,
			}
			eventChannel.UpdateTableContents <- tableContentsInfo
			return nil
		case constant.KeyRight, constant.KeyEnter:
			if fileTable.selectedFileRow.FileInfo != nil && fileTable.selectedFileRow.FileInfo.IsDir() {
				tableContentsInfo := &model.TableContentsInfo{
					Word:            "",
					Path:            filepath.Join(fileTable.path, fileTable.selectedFileRow.FileInfo.Name()),
					TableUpdateType: model.LeftRightPressed,
				}
				eventChannel.UpdateTableContents <- tableContentsInfo
			}
			return nil
		}

		if event.Rune() == 'r' {
			tableContentsInfo := &model.TableContentsInfo{
				Word:            "",
				Path:            initDir,
				TableUpdateType: model.Other,
			}
			eventChannel.UpdateTableContents <- tableContentsInfo
			return nil
		}
		return event
	})

	tableContentsInfo := &model.TableContentsInfo{
		Word:            "",
		Path:            initDir,
		TableUpdateType: model.Other,
	}

	fileTable.ShowFiles(eventChannel, tableContentsInfo)
	return fileTable
}

// EventReceiver is the function where you receive various event.
// here, you should update primitives.
func (fileTable *FileTable) EventReceiver(app *tview.Application, eventChannel model.EventChannel) {
	for {
		select {
		case updateTableContents := <-eventChannel.UpdateTableContents:
			fileTable.Table.Clear()
			app.QueueUpdateDraw(func() {
				fileTable.ShowFiles(eventChannel, updateTableContents)
			})
		case fileRow := <-eventChannel.UpdateCellBackground:
			app.QueueUpdateDraw(func() {
				fileTable.updateTableRowColor(fileRow)
			})
		}
	}
}

// ShowFiles shows file-table.
func (fileTable *FileTable) ShowFiles(eventChannel model.EventChannel, updateTableContents *model.TableContentsInfo) {

	if updateTableContents.TableUpdateType != model.SearchInputText {
		fileTable.files = data.ReadAllFileName(updateTableContents.Path)
		fileTable.path = updateTableContents.Path
	}
	if updateTableContents.TableUpdateType == model.LeftRightPressed ||
		updateTableContents.TableUpdateType == model.SearchInputText {

		fileTable.beforeSelectedRow = constant.ItemNotSelected
		fileRow := &model.FileRow{
			FileInfo: nil,
			Row:      constant.ItemNotSelected,
		}
		fileTable.selectedFileRow = fileRow
	}

	filteredList := data.FilterFile(fileTable.files, updateTableContents.Word)
	fileTable.item = filteredList

	cols, rows := 1, len(filteredList)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			fileName := filteredList[r].Name()
			cell := tview.NewTableCell(fileName)

			if filteredList[r].IsDir() {
				cell.SetTextColor(tcell.NewRGBColor(100, 40, 190))
			}

			fileRow := &model.FileRow{
				FileInfo: filteredList[r],
				Row:      r,
			}

			cell.SetReference(fileRow)

			cell.SetClickedFunc(func() bool {
				reference := cell.GetReference()
				if fileRow, ok := reference.(*model.FileRow); ok {

					updatePreview := &model.UpdatePreview{
						FileRow: fileRow,
						AbsPath: filepath.Join(fileTable.path, fileRow.FileInfo.Name()),
					}
					eventChannel.UpdateCellBackground <- fileRow
					eventChannel.UpdatePreview <- updatePreview
				}
				return true
			})
			fileTable.Table.SetCell(r, c, cell)

		}
	}
}

// updateTableRowColor is the function that update table-cell color.
// FileRow as argument is a new selected item.
// You can see a old selected item position in the fileTable.selectedItem.Row
func (fileTable *FileTable) updateTableRowColor(fileRow *model.FileRow) {

	fileTable.selectedFileRow = fileRow
	oldRow := fileTable.beforeSelectedRow

	if oldRow != constant.ItemNotSelected {
		cell := fileTable.Table.GetCell(oldRow, 0)
		cell.SetBackgroundColor(tcell.NewRGBColor(0, 0, 0))
	}

	cell := fileTable.Table.GetCell(fileRow.Row, 0)
	cell.SetBackgroundColor(tcell.NewRGBColor(100, 120, 200))

	log.Log.Println("updateTableCellColor", fileRow.Row)
	fileTable.beforeSelectedRow = fileRow.Row
}
