package component

import (
	"filemanager/data"
	"filemanager/model"

	"github.com/rivo/tview"
)

// FilePreview is a struct that is used in the preview window.
type FilePreview struct {
	TextView *tview.TextView
}

// CreatePreview makes a tview.TextView of the preview window.
// In this function, I set the border and title of a preview window.
func CreatePreview() *FilePreview {
	textView := tview.NewTextView()
	textView.SetBorder(true)
	textView.SetTitle("preview")
	filePreview := &FilePreview{textView}
	return filePreview
}

// EventReceiver is a receiver of eventChannel and then update the preview window.
// This should be used with goroutine.
func (filePreview *FilePreview) EventReceiver(app *tview.Application, eventChannel model.EventChannel) {
	for {
		select {
		case updatePreview := <-eventChannel.UpdatePreview:
			app.QueueUpdateDraw(func() {
				if updatePreview.FileRow.FileInfo.IsDir() {
					filePreview.TextView.SetText("The selected item is a directory.. We can't show that..")
				} else {
					if updatePreview.FileRow.FileInfo.Size() < 5000 {
						contents := data.ReadFile(updatePreview.AbsPath)
						filePreview.TextView.SetText(contents)
					} else {
						filePreview.TextView.SetText("The selected item is too big.. We can't show that..")
					}
				}
			})
		}
	}
}
