package main

import (
	"fmt"
)

/* This wont work because the main thread executes and finishes before the channel sends the message */
func main() {

	ch := make(chan string)

	ch <- "Hello"
	// <- recieve operator
	fmt.Println(<-ch)

}
