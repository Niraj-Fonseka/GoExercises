package main

//godoc -v go_exercises/GoDoc/packageDoc > docs.md
//Had to run that
import (
	"fmt"
	docs "go_exercises/GoDoc/packageDoc"
)

func main() {
	fmt.Println("Initializing go doc test")
	docs.Example()
}
