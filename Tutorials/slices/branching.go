package main

import "./greeting"

func main() {
	//	var s = greeting.Salutation{"Bob", "Hello"}
	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What's up ?"},
	}

	//deleting in a slice
	slice = append(slice[:1], slice[2:]...)
	slice = append(slice, greeting.Salutation{"Frank", "Hi"})
	greeting.Greet(slice, greeting.CreatePrintFunction("??"), true, 5)
}
