package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello"
	w.Write([]byte(message))
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uplading file")
	//r.ParseMultipartForm(50 << 20)

	file, handler, err := r.FormFile("myfile")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("FileName : ", handler.Filename)

	// tempFile, err := ioutil.TempFile("temp-csv", "upload-csv-*.csv")

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// defer tempFile.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bytes))

}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/upload", handleFileUpload)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
