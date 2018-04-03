package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are thr times that try mne's souls. \n"

	words := strings.Split(phrase, " ")

	//To make it a buffered channel we need to specify the length of the argument
	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	//if the channel is empty the loop will terminate
	for msg := range ch {
		fmt.Print(msg + " ")
	}

}
