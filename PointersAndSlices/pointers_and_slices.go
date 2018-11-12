package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Pointers and Slices in Go ---")

	arr := make([]int, 5)
	fmt.Println("Passing slice : regular")
	fmt.Printf("Address of the arr : %p \n", &arr)
	fmt.Println()

	Pass_Slice_Regular(arr)
	fmt.Println(arr)
	fmt.Println()
	fmt.Println("Passing by address")
	Pass_Slice_PassByAddress(&arr)
	fmt.Println(arr)
	fmt.Println()
}

func Pass_Slice_Regular(arr []int) {
	fmt.Println("Inside the pass_slice_regular function")
	fmt.Printf("Address of the arr inside the function : %p \n", &arr)
	fmt.Println("Appending to the slice")
	arr = append(arr, 10)
	fmt.Println("Testing assignement in the regular function")
	arr[1] = 22

}

func Pass_Slice_PassByAddress(arr *[]int) {
	fmt.Println("Inside the pass_slice_regular function")
	fmt.Printf("Address of the arr inside the function : %p \n", arr)
	fmt.Printf("Appending to the slice")
	*arr = append(*arr, 10)
	fmt.Println()

}
