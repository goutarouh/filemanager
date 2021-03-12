package data

import (
	"io/ioutil"
	"os"
	"strings"
)

// ReadAllFileName get all dirs and files at the path you provide to this as the arg.
func ReadAllFileName(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		files = make([]os.FileInfo, 0)
	}

	return files
}

// FilterFile filter "fileList" by "word" arg.
func FilterFile(fileList []os.FileInfo, word string) []os.FileInfo {
	filteredList := make([]os.FileInfo, 0, len(fileList))
	for _, file := range fileList {
		if strings.Contains(file.Name(), word) {
			filteredList = append(filteredList, file)
		}
	}
	return filteredList
}

// ReadFile reads contents of file.
// if err, it returns default string.
func ReadFile(fileName string) string {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "sorry, I can't show you the file contents."
	}
	return string(contents)
}
