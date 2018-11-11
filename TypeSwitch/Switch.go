package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type TestOrders struct {
	Id        int    `gorm:"primary_key"`
	Year      int    ` json:"year"`
	Timestamp string `gorm:"size:50" json:"timestamp"`
	Boss      int    ` json:"boss"`
}

func printType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer:", x)
	case bool:
		fmt.Println("Bool:", x)
	case string:
		fmt.Println("String:", x)
	case Person:
		fmt.Println("Person :", x.(Person).Age)

	case TestOrders:
		fmt.Println("Person :", x.(TestOrders).Boss)
	default:

		fmt.Println("Who cares?", x)
	}
}

func main() {
	printType(1)
	printType("Hello")
	printType(true)

	//Checking custom type
	p := Person{
		Name: "Niraj",
		Age:  10,
	}
	printType(p)

	orders := TestOrders{
		Year: 2010,
		Boss: 29,
	}

	printType(orders)
}
