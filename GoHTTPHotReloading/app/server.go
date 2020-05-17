package main

import (
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello"
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/hello", sayHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
