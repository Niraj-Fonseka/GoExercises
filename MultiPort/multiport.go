package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	errChan := make(chan error)
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello")
	})

	fmt.Println("Staritng server on port 9090")
	go RunServerOneInBackground(errChan)

	http.HandleFunc("/yow", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello yow")
	})

	fmt.Println("Staritng server on port 9091")

	go RunServerTwoInBackground(errChan)

	for err := range errChan {
		log.Fatal(err)
	}
}

func RunServerOneInBackground(c chan error) {
	c <- http.ListenAndServe(":9090", nil)
}

func RunServerTwoInBackground(c chan error) {
	c <- http.ListenAndServe(":9091", nil)
}
