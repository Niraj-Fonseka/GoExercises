package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	first := FirstFunc()
	fmt.Printf("Returned : %s \n", first)

	second := SecondFunc()
	fmt.Printf("Returned : %s \n", second)

	third := ThirdFunc()
	fmt.Printf("Returned : %s \n", third)

	fourth := FourthFunc()
	fmt.Printf("Returned : %s \n", fourth)

	fmt.Println(time.Since(start))

}

func FirstFunc() string {
	fmt.Println("Executing first function")
	time.Sleep(4 * time.Second)
	return "first"
}

func SecondFunc() string {
	fmt.Println("Executing second function")
	time.Sleep(3 * time.Second)
	return "second"
}

func ThirdFunc() string {
	fmt.Println("Executing third function")
	time.Sleep(2 * time.Second)

	return "third"
}

func FourthFunc() string {
	fmt.Println("Executing fourth function")
	time.Sleep(1 * time.Second)

	return "fourth"
}
