package main

import (
	"fmt"

	"./greeting"
)

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {
	//	var s = greeting.Salutation{"Bob", "Hello"}
	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What's up ?"},
	}

	fmt.Fprintf(&salutations[0], "The count is %d ", 10)
	salutations.Greet(greeting.CreatePrintFunction("??"), true, 5)
}
