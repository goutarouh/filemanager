package main

import (
	"filemanager/app"
	"filemanager/log"
)

func init() {
	log.SetLog()
	log.Log.Println("filemanager start")
}

func main() {
	defer log.Log.Println("filemanager end")

	app.ExecuteApp()
}
