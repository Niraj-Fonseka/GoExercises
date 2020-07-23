package main

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Username string `env:"USERNAME,required"`
	Password string `env:"PASSWORD,required"`
}

func main() {

	ctx := context.Background()

	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("Username : %s \n", c.Username)
	log.Printf("Password : %s \n", c.Password)
}
