package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handle route using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server!")
	})

	http.HandleFunc("/upload", handler)

	// listen to port
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "http://localhost:9091/upload"

	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}
