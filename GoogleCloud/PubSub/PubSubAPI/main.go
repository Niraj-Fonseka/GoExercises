package main

import (
	"golang.org/x/net/context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

func main() {
	go initPubSub()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health is good")
	})

	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
		if proj == "" {
			fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
			os.Exit(1)
		}
		client, err := pubsub.NewClient(ctx, proj)

		const t = "example-topic"
		// Create a topic to subscribe to.

		fmt.Println("Creating Topic")
		topic := client.Topic(t)
		ok, err := topic.Exists(ctx)
		if err != nil {
			log.Fatal(err)
		}
		if ok {
			fmt.Println("Exists")
		}
		// Publish 10 messages on the topic.
		var results []*pubsub.PublishResult
		for i := 0; i < 2; i++ {
			res := topic.Publish(ctx, &pubsub.Message{
				Data: []byte(fmt.Sprintf("hello world #%d", i)),
			})
			results = append(results, res)
		}

		// Check that all messages were published.
		for _, r := range results {
			_, err := r.Get(ctx)
			if err != nil {
				fmt.Println(err)
			}
		}

		// [START pubsub_subscribe
	})

	http.ListenAndServe(":6060", nil)
}

func initPubSub() {
	ctx := context.Background()
	proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if proj == "" {
		fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}
	client, err := pubsub.NewClient(ctx, proj)

	const t = "example-topic"
	// Create a topic to subscribe to.

	fmt.Println("Creating Topic")
	topic := client.Topic(t)
	ok, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Println("Exists")
	}

	name := "example-subscription"

	// [START pubsub_create_pull_subscription]
	sub, err := client.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Created subscription: %v\n", sub)

	fmt.Println("Trying to get messages")
	// var mu sync.Mutex
	// received := 0
	creatingSub := client.Subscription(name)
	cctx, _ := context.WithCancel(ctx)
	errRecv := creatingSub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		// mu.Lock()
		// defer mu.Unlock()
		// received++
		// if received == 3 {
		// 	cancel()
		// }
	})
	if errRecv != nil {
		fmt.Println(errRecv)
	}
}
