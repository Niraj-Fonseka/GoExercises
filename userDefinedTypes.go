package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {

	var s = Salutation{"Niraj", "Hello"}
	fmt.Println(&s.name)
	fmt.Println(&s.greeting)
}
