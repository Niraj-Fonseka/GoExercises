package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	wg.Add(4)

	go FirstFunc()

	go SecondFunc()

	go ThirdFunc()

	go FourthFunc()
	wg.Wait()

	fmt.Printf("Total time to finish : %s \n", time.Since(start).String())

}

func FirstFunc() {
	fmt.Println("-- Executing first function --")
	time.Sleep(7 * time.Second)
	fmt.Println("-- First Function finished --")
	defer wg.Done()
}

func SecondFunc() {
	fmt.Println("-- Executing second function --")
	time.Sleep(5 * time.Second)
	fmt.Println("-- Second Function finished --")
	defer wg.Done()

}

func ThirdFunc() {
	fmt.Println("-- Executing third function --")
	time.Sleep(2 * time.Second)
	fmt.Println("-- Third Function finished --")
	defer wg.Done()

}

func FourthFunc() {
	fmt.Println("-- Executing fourth function --")
	time.Sleep(10 * time.Second)
	fmt.Println("-- Fourth Function finished --")
	defer wg.Done()

}
