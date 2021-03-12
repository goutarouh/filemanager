package model

import (
	"os"
)

// EventChannel is the struct that manage all event happened in
// the screen in this app.
type EventChannel struct {

	// UpdateTableContents is an event that occurred when search-
	// input-text is changed and then pressed Enter.
	// It is send to the table to filter file list.
	UpdateTableContents chan *TableContentsInfo

	// UpdatePreview is an event that occurred when item is selected
	// or selected item is changed.
	// It is send to the Preview to show contents.
	UpdatePreview chan *UpdatePreview

	// SelectedTableRow is an event that occurred when item is selected
	// or selected item is changed.
	// It is send to the Table to update its cell background color.
	UpdateCellBackground chan *FileRow
}

type TableUpdateType int

const (
	SearchInputText  TableUpdateType = 0
	LeftRightPressed TableUpdateType = 1
	Other            TableUpdateType = 2
)

// TableContentsInfo is a struct that is used for updating
// table contents and carried on channel.
type TableContentsInfo struct {
	Word            string
	Path            string
	TableUpdateType TableUpdateType
}

// UpdatePreview is a struct that is used for updating preview
type UpdatePreview struct {
	FileRow *FileRow
	AbsPath string
}

// CellInfo is a struct that is used to get the cell
// information when clicked.
type CellInfo struct {
	FileRow   FileRow
	TableRows TableRows
}

// FileRow is a struct that mean a data of a table-row.
type FileRow struct {
	FileInfo os.FileInfo
	Row      int
}

// TableRows is a struct that is used for updating
// table-cell color with channel.
// BeforeRow is a number which is old row position and
// it is used to update default background color(black).
// AfterRow is a number which is new row position and
// it is used to update selected status color.
type TableRows struct {
	BeforeRow int
	AfterRow  int
}
