package main

import (
	"fmt"
)

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	// msg := Message{
	// 	To:      []string{"test@email.com"},
	// 	From:    "from@email.com",
	// 	Content: "Content stuff",
	// }

	// failedMessage := FailedMessage{
	// 	ErrorMessage:    "Error Message",
	// 	OriginalMessage: Message{},
	// }

	// //msgCh <- msg

	// errCh <- failedMessage
	//select block is overloaded to handle multiple tasks
	select {
	case recievedMsg := <-msgCh:
		fmt.Println(recievedMsg)
	case recievedError := <-errCh:
		fmt.Println(recievedError)
	default:
		fmt.Println("No Messages")
	}

	fmt.Println(<-msgCh)
	fmt.Println(<-errCh)
}
