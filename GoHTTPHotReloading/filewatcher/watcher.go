package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fin-man/file-watcher/watcher"
	"github.com/fsnotify/fsnotify"
)

func main() {
	log.Println("Starting a new filewatcher ")
	fw := watcher.NewFileWatcher()

	pwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	appPath := pwd + "/app"

	fw.Watch(appPath, ProcessFile)
}

func Watch(volume string, callback func(data ...interface{}) error) {
	log.Println("Starting the watch .. ")
	fmt.Println("watching : ", volume)
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-f.Watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("created file:", event.Name)
					fileName := f.ExtractFileName(event.Name)
					filePath := f.ExtractFilePath(event.Name)
					err := callback(filePath, fileName) // make this get executed in a go routine with a channel
					if err != nil {
						log.Printf("ERROR : %v \n", err)
					}
				}
			case err, ok := <-f.Watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err := f.Watcher.Add(volume)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
