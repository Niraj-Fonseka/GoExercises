package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:9000/home")

	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Response %d!", 500)
		return
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bytes))
	fmt.Fprintf(w, "Response %d!", resp.StatusCode)
}

func main() {
	http.HandleFunc("/client", handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
