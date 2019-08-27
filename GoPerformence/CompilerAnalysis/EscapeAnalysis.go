package main

import "fmt"

/*
 go build -gcflags="-m -m"
 */

 type T struct{
 }

func main() {
	makeMap()
}


func inc(x *int) {
	*x++
}

func makeMap() {
	t := new(T)
	fmt.Println(t)
	//m := make(map[string]int)
	//fmt.Println(m)
}

