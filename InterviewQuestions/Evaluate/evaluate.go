package main

import (
	"fmt"
)

func main() {

	fmt.Println(evaluate("6+9-12"))
}

func evaluate(str string) int {
	var stack []int

	for _, c := range str {
		stack = append(stack, int(c))
	}

	//isSecondNum := false
	final := string(stack[0])
	stack = stack[1 : len(stack)-1]
	num := ""

	var calculateStack []string

	calculateStack = append(calculateStack, final)
	//calculateStack[0] = final
	for len(stack) != 0 {
		//popping out from queue

		char := stack[0]
		charToString := string(char)
		//fmt.Println(string(char))

		if charToString != "+" || charToString != "-" {
			num += charToString
		} else if charToString == "+" || charToString == "-" {
			calculateStack = append(calculateStack, charToString)
			calculateStack = append(calculateStack, num)
		}

		fmt.Println(len(calculateStack))
		if len(calculateStack) == 3 { //we have three characters
			fmt.Println("this is the stack")
			fmt.Println(calculateStack)
		}
	}
	fmt.Println("Printing Stack")
	fmt.Println(stack)

	return 0
}
