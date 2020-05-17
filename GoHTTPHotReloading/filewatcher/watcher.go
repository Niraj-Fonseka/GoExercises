package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

var (
	appPath = "/home/hungryotter/go/src/GoExercises/GoHTTPHotReloading/app"
)

func main() {
	fw, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	// pwd, err := os.Getwd()

	// if err != nil {
	// 	panic(err)
	// }

	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		log.Println(err)
		os.Mkdir(appPath, os.ModeDir)
	}

	Watch(fw, appPath, build)
}

func build() {
	fmt.Println("Building")

	cmd := exec.Command("go", "build", "-o", "/home/hungryotter/go/src/GoExercises/GoHTTPHotReloading/app/app", "/home/hungryotter/go/src/GoExercises/GoHTTPHotReloading/app/*")

	bytes, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bytes))
}

func Watch(fw *fsnotify.Watcher, volume string, callback func()) {
	log.Println("Starting the watch .. ")
	fmt.Println("watching : ", volume)
	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-fw.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("created file:", event.Name)

					callback() // make this get executed in a go routine with a channel

				}
			case err, ok := <-fw.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err := fw.Add(volume)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
