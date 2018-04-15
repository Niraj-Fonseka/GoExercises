package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Running")

	go sleepFunction(3, 3*time.Second)
	go sleepFunction(10, 10*time.Second)
	go sleepFunction(5, 5*time.Second)

	fmt.Scanln() // wait for Enter Key

}

func sleepFunction(secs int, t time.Duration) {
	fmt.Printf("Sleeping for : %d secs \n", secs)
	time.Sleep(t)
	fmt.Printf("Done : %d \n", secs)
}
