package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Log はログの本体です
var Log = logrus.New()

// SetLog はログの設定を行います。
func SetLog() {
	fileName := "debug.log"
	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}

	errorLogFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
	Log.Out = errorLogFile
}
