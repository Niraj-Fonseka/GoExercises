package main

import (
	"context"
	"log"

	"github.com/shomali11/slacker"
)

func main() {
	bot := slacker.NewClient("token-here")

	definition := &slacker.CommandDefinition{
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	}

	bot.Command("ping", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
