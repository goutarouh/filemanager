package component

import (
	"filemanager/data"
	"os"
	"strings"

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
func (fileTable *FileTable) EventReceiver(searchInputTextChannel chan string) {
	for {
		select {
		case word := <-searchInputTextChannel:
			fileTable.Table.Clear()
			fileTable.ShowFiles(word)
		}
	}
}

// ShowFiles はテーブルに一覧を表示します。
func (fileTable *FileTable) ShowFiles(word string) {
	cols, rows := 1, len(fileTable.files)-1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			fileName := fileTable.files[r].Name()

			if !strings.Contains(fileName, word) {
				continue
			}

			fileTable.Table.SetCell(r, c,
				tview.NewTableCell(fileName).
					SetAlign(tview.AlignCenter))
		}
	}
}
