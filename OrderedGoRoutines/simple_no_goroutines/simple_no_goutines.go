package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	FirstFunc()

	SecondFunc()

	ThirdFunc()

	FourthFunc()

	fmt.Printf("Total time to finish : %s \n", time.Since(start).String())

}

func FirstFunc() {
	fmt.Println("-- Executing first function --")
	time.Sleep(7 * time.Second)
	fmt.Println("-- First Function finished --")
}

func SecondFunc() {
	fmt.Println("-- Executing second function --")
	time.Sleep(5 * time.Second)
	fmt.Println("-- Second Function finished --")
}

func ThirdFunc() {
	fmt.Println("-- Executing third function --")
	time.Sleep(2 * time.Second)
	fmt.Println("-- Third Function finished --")
}

func FourthFunc() {
	fmt.Println("-- Executing fourth function --")
	time.Sleep(10 * time.Second)
	fmt.Println("-- Fourth Function finished --")
}
