package log

import (
	"filemanager/constant"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Log はログの本体です
var Log = logrus.New()

// SetLog はログの設定を行います。
func SetLog() {
	if _, err := os.Stat(constant.LogFileName); err == nil {
		err = os.Remove(constant.LogFileName)
	}

	errorLogFile, err := os.OpenFile(constant.LogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
	Log.Out = errorLogFile
}
