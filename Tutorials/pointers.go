package main

import "fmt"

func main() {

	message := "Hello World in Go ! "

	greeting := &message

	fmt.Println(message, greeting)
}
