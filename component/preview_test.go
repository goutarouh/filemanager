package component

//
////no-use
//
//import (
//	"filemanager/mock"
//	"filemanager/model"
//	"github.com/rivo/tview"
//	"testing"
//	"time"
//)
//
//var app *tview.Application
//var filePreview *FilePreview
//var eventChannel model.EventChannel
//var fileRowDirectory *model.FileRow
//var fileRowBigFile *model.FileRow
//var fileRow1NormalFIle *model.FileRow
//
//
////init needed variable
//func init() {
//	app = tview.NewApplication()
//	filePreview = CreatePreview()
//	go app.SetRoot(filePreview.TextView, false).Run()
//
//	updateTableContents := make(chan *model.TableContentsInfo)
//	updatePreview := make(chan *model.FileRow)
//	updateCellBackground := make(chan *model.FileRow)
//
//	eventChannel = model.EventChannel{
//		UpdateCellBackground: updateCellBackground,
//		UpdatePreview: updatePreview,
//		UpdateTableContents: updateTableContents,
//	}
//
//	fileRowDirectory = &model.FileRow{
//		FileInfo: mock.FileMock{FileName: "Dir", IsDirectory: true},
//		Row: 1,
//	}
//
//	fileRowBigFile = &model.FileRow{
//		FileInfo: mock.FileMock{FileName: "File1", FileSize: 300000, IsDirectory: false},
//		Row: 2,
//	}
//
//	fileRow1NormalFIle = &model.FileRow{
//		FileInfo: mock.FileMock{FileName: "File1", FileSize: 1000, IsDirectory: false},
//		Row: 3,
//	}
//}
//
//func TestEventReceiver(t *testing.T) {
//
//	//if err := app.SetRoot(filePreview.TextView, true).EnableMouse(true).Run(); err != nil {
//	//	panic(err)
//	//}
//	app.SetRoot(filePreview.TextView, true).
//
//	go filePreview.EventReceiver(app, eventChannel)
//	time.Sleep(1000)
//
//	eventChannel.UpdatePreview <- fileRowDirectory
//	time.Sleep(1000)
//
//	t.Fatal("test: ", filePreview.TextView.GetText(true), ": test")
//
//
//	//eventChannel.UpdatePreview <- fileRowBigFile
//	//eventChannel.UpdatePreview <- fileRow1NormalFIle
//}
