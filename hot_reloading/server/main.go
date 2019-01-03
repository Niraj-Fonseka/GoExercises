package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello Changed"
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/hello", sayHello)

	fmt.Println("Running app")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
