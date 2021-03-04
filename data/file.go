package data

import (
	"io/ioutil"
	"os"
	"strings"
)

// ReadAllFile は全てのファイルをリストで返します。
func ReadAllFile() []os.FileInfo {
	dir := "C:/Users/gouta/Documents/go-develop/github.com/goutarouh/filemanager"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		files = make([]os.FileInfo, 0)
	}

	return files
}

// FilterFile はファイル名の絞り込みを行います。
func FilterFile(fileList []os.FileInfo, word string) []os.FileInfo {
	filteredList := make([]os.FileInfo, 0, len(fileList))
	for _, file := range fileList {
		if strings.Contains(file.Name(), word) {
			filteredList = append(filteredList, file)
		}
	}
	return filteredList
}
