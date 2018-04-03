package greeting

import "fmt"

//Salutation struct
type Salutation struct {
	Name     string
	Greeting string
}

//Printer type
type Printer func(string)

//Greet function
func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {
	for _, s := range salutation {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

//TypeSwitchTest function
func TypeSwitchTest(x interface{}) {

	switch t := x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println(t)
		fmt.Println("unknown")
	}
}

func GetPrefix(name string) (prefix string) {

	switch {
	case name == "Bob":
		prefix = "Mr "
		fallthrough
	case name == "Joe", name == "Amy", len(name) == 10:
		prefix = "Dr "
	case name == "Mary":
		prefix = "Mrs "
	default:
		prefix = "Dude "
	}

	return
}

//CreateMessage function this is a variadic func
func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey! " + name
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
