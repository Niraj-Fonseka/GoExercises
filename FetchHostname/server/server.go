package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := fmt.Sprintf("Request From : %s , Path : %s", r.Host, r.URL.Path)

	fmt.Fprintf(w, "client hostname %s!", client)
}

func main() {
	http.HandleFunc("/home", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
