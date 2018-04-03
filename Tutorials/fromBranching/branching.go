package main

import "./greeting"

func main() {
	//	var s = greeting.Salutation{"Bob", "Hello"}
	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What's up ?"},
	}
	greeting.Greet(slice, greeting.CreatePrintFunction("??"), true, 5)
}
