package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	MysleepAndTalk(ctx, 5*time.Second, "hello")
}

func MysleepAndTalk(ctx context.Context, d time.Duration, s string) {
	time.Sleep(d)
	fmt.Println(s)
}
