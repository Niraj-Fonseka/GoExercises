package Functions

import (
	"fmt"
)

//FunctionTestOne
//Regular function that takes in no parameters
func FunctionTestOne() {
	fmt.Println("In FunctionTestOne")
}

//FunctionTesttwo
//Function that takes in one arguments and returns one
func FunctionTestTwo(a string) string {
	fmt.Println("In FunctionTestTwo")
	return a
}

//FunctionTestThree
//Function that takes in two arguments and returns two arguments
func FunctionTestThree(a string, b int) (string, int) {
	fmt.Println("In FunctionTestThree")
	return a, b
}
