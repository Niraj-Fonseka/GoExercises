package greeting

import "fmt"

//Salutation struct
type Salutation struct {
	Name     string
	Greeting string
}

//Renamable Interface
type Renamable interface {
	Rename(newName string)
}

//Rename function
func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

//Salutations type
type Salutations []Salutation

//Printer type
type Printer func(string)

//Greet function
func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {
	for _, s := range salutations {
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

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println("unknown")
	}
}

//GetPrefix function
func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Amy":  "Dr ",
		"Mary": "Mrs ",
	}

	prefixMap["Joe"] = "Jr "

	//	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}
	return "Dude "

	// /	return prefixMap[name]
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
