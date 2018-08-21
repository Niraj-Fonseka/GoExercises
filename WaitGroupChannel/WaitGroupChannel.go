package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepFun(sec time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(sec * time.Second)
	fmt.Println("goroutine exit")
}

func main() {

	go WaitGroupRoutine()
	go WaitGroupRoutine()
	go WaitGroupRoutine()
	go WaitGroupRoutine()

	fmt.Println("Main goroutine exit")
	fmt.Scanln()
}

func WaitGroupRoutine() {
	var wg sync.WaitGroup

	wg.Add(3)
	go sleepFun(1, &wg)
	go sleepFun(3, &wg)
	go sleepFun(2, &wg)
	wg.Wait()
}
