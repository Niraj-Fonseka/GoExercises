package main

import "fmt"

/*
function that takes a single (numeric) argument and which returns another function that is an accumulator.
The returned accumulator function in turn also takes a single numeric argument,
and returns the sum of all the numeric values passed in so far to that accumulator (including the initial value passed when the accumulator was created)
*/

/*
 Input : Empty Interface
 Return : function that takes in an empty interface that returns an interaface
*/
func accumulator(sum interface{}) func(interface{}) interface{} {
	return func(nv interface{}) interface{} {
		switch s := sum.(type) {
		case int:
			switch n := nv.(type) {
			case int:
				sum = s + n
			case float64:
				sum = float64(s) + n
			}
		case float64:
			switch n := nv.(type) {
			case int:
				sum = s + float64(n)
			case float64:
				sum = s + n
			}
		default:
			sum = nv
		}
		return sum
	}
}

func main() {
	x := accumulator(2)
	fmt.Println(x(5))
	accumulator(3)
	fmt.Println(x(2.3))
}
