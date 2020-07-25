package main

import (
	"log"
	"net/http"
)

var (
	REDIRECT_URL = "http://localhost:8080"
)

func HandleUploadRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Redirecting")

	http.Redirect(w, r, REDIRECT_URL+"/upload", 301)
}

func main() {
	log.Println("Starting server ..")

	http.HandleFunc("/upload", HandleUploadRequest)

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
