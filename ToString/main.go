package main

import "fmt"

func main() {

	obj := NewObject()

	fmt.Println(obj)

	nObj := NewNestedObject(obj)

	fmt.Println(nObj)
}
