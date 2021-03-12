package component

import (
	"github.com/rivo/tview"
)

func CreateDescription() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetBorder(true)
	textView.SetTitle("description")
	textView.SetText("(left or right): move dir   (up or down): change file")
	return textView
}
