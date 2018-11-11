/*
/ Acts more like types
/
/
*/

package main

import "fmt"

//Salutation struct
type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

//Greet function
func Greet(salutation Salutation, do Printer) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	//fmt.Println(message, alternate)
	do(message)
	do(alternate)
}

//CreateMessage function this is a variadic func
func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	message = greeting[1] + " " + name
	alternate = "Hey !! " + name
	return
}

func do(s string) {
	fmt.Print(s)
}

//CreatePrintFunction  func
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func main() {
	var s = Salutation{}
	s.name = "Bob"
	s.greeting = "Hello"

	Greet(s, CreatePrintFunction("!!!"))
}
