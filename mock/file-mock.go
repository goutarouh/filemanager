package mock

import (
	"os"
	"time"
)

type FileMock struct {
	FileName    string
	FileSize    int64
	IsDirectory bool
}

func (fileMock FileMock) Name() string {
	return fileMock.FileName
}

func (fileMock FileMock) Size() int64 {
	return fileMock.FileSize
}

func (fileMock FileMock) Mode() os.FileMode {
	return 0
}

func (fileMock FileMock) ModTime() time.Time {
	return time.Time{}
}

func (fileMock FileMock) IsDir() bool {
	return fileMock.IsDirectory
}

func (fileMock FileMock) Sys() interface{} {
	return "interface"
}
