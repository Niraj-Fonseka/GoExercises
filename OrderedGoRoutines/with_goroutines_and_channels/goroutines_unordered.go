package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	firstchan := make(chan string)
	secondchan := make(chan string)
	thirdchan := make(chan string)
	fourthchan := make(chan string)

	go FourthFunc(fourthchan)
	go SecondFunc(secondchan)
	go FirstFunc(firstchan)
	go ThirdFunc(thirdchan)

	fmt.Printf("Returned : %s \n", <-firstchan)
	fmt.Printf("Returned : %s \n", <-secondchan)
	fmt.Printf("Returned : %s \n", <-thirdchan)
	fmt.Printf("Returned : %s \n", <-fourthchan)

	fmt.Println(time.Since(start))

}

func FirstFunc(ch chan<- string) {
	fmt.Println("Executing first function")
	time.Sleep(7 * time.Second)
	ch <- "first"
}

func SecondFunc(ch chan<- string) {
	fmt.Println("Executing second function")
	time.Sleep(5 * time.Second)
	ch <- "second"
}

func ThirdFunc(ch chan<- string) {
	fmt.Println("Executing third function")
	time.Sleep(2 * time.Second)

	ch <- "third"
}

func FourthFunc(ch chan<- string) {
	fmt.Println("Executing fourth function")
	time.Sleep(10 * time.Second)

	ch <- "fourth"
}
