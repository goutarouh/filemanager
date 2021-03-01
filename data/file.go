package data

import (
	"io/ioutil"
	"log"
	"os"
)

// type File struct {
// 	Name string
// }

func ReadAllFile() []os.FileInfo {
	dir := "C:/Users/gouta/Documents/go-develop/github.com/goutarouh/filemanager"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}
