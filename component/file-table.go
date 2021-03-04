package component

import (
	"filemanager/data"
	"filemanager/log"
	"os"

	"github.com/rivo/tview"
)

// FileRow はファイル一覧テーブルの一行分のファイル情報を表します
type fileRow struct {
	Name string
	dir  bool
	tag  string
}

// FileTable はファイル一覧を表示するテーブル情報を保持します。
type FileTable struct {
	Table *tview.Table
	files []os.FileInfo
}

// CreateTable はファイル一覧テーブルを作成します。
func CreateTable() *FileTable {
	table := tview.NewTable()
	table.SetTitle("file list")
	table.SetBorder(true)

	data := data.ReadAllFile()

	fileTable := &FileTable{table, data}

	fileTable.ShowFiles("")
	return fileTable
}

// EventReceiver はイベントを取得します。
func (fileTable *FileTable) EventReceiver(app *tview.Application, searchWordChannel chan string) {
	for {
		select {
		case word := <-searchWordChannel:
			fileTable.Table.Clear()

			app.QueueUpdateDraw(func() {
				fileTable.ShowFiles(word)
			})
		}
	}
}

// ShowFiles はテーブルに一覧を表示します。
func (fileTable *FileTable) ShowFiles(word string) {

	filteredList := data.FilterFile(fileTable.files, word)

	for _, file := range filteredList {
		log.Log.Println("  ", file.Name())
	}

	cols, rows := 1, len(filteredList)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			fileName := filteredList[r].Name()
			cell := tview.NewTableCell(fileName)

			fileTable.Table.SetCell(r, c, cell)

		}
	}
}
