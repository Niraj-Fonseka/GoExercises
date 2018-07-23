package main

import (
	"fmt"
	"log"

	"github.com/shomali11/slacker"
)

func main() {
	bot := slacker.NewClient("Token Here")

	fmt.Println(bot)

	bot.Help(func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("Your own help function...")
	})

	bot.Command("ping", "Ping!", func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("pong")
	})

	err := bot.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
