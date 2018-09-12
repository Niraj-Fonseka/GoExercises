package main

import (
	"fmt"
	"golang.org/x/net/context"
	
	"time"

	"cloud.google.com/go/pubsub"
)

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cl, err := pubsub.NewClient(ctx, "marwan-test")
	if err != nil {
                panic(err)
	}

	// _, err = cl.CreateTopic(ctx, "email")
	// if err != nil {
	// 	fmt.Println(err) // if this program is running inside a docker image, it will panic, otherwise -- create topic succeeds.
	// }

	const top = "email"
	// Create a topic to subscribe to.

	fmt.Println("Creating Topic")
	topic := cl.Topic(top)
	ok, err := topic.Exists(ctx)
	if err != nil {
		fmt.Println(err)
	}
	if ok {
		fmt.Println("Exists")
	}
}