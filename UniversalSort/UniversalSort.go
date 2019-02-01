package main

import (
	"fmt"
	"math/rand"
	"reflect"
	// "reflect"
)

type Vertex struct {
	X int
	Y int
	S string
}

type Circle struct {
	Radius int
	Color  string
}

func main() {
	fmt.Println("Hello World")
	// a := []int{1,3,5,6,1,2,4,56,123,2}
	// fmt.Println(SortByField(a))
	var vs []Vertex

	vs = append(vs, Vertex{5, 2, "Zebra"})
	vs = append(vs, Vertex{1, 3, "Wne"})
	vs = append(vs, Vertex{1, 5, "Aree"})
	vs = append(vs, Vertex{2, 3, "Two"})

	fmt.Println(SortUni(vs, "S"))

}

func switchThis(vs interface{}) {
	switch vs.(type) {
	case []Vertex:
		casted := vs.([]Vertex)
		fmt.Println("Vertex", casted)
	case []Circle:
		fmt.Println("Circle")
	default:
		fmt.Println("Not a type")
	}
}

func compare(field1 interface{}, field2 interface{}) bool {

	switch field1.(type) {
	case int:
		if field1.(int) < field2.(int) {
			return true
		} else {
			return false
		}
	case string:
		if field1.(string) < field2.(string) {
			return true
		} else {
			return false
		}
	default:
		return false
	}

}

func SortUni(a interface{}, field string) interface{} {

	switch a.(type) {
	case []Vertex:
		k := a.([]Vertex)
		if len(k) < 2 {
			return a
		}

		left, right := 0, len(k)-1

		pivot := rand.Int() % len(k)

		k[pivot], k[right] = k[right], k[pivot]

		for i, _ := range k {

			if compare(getFieldValue(k[i], field), getFieldValue(k[right], field)) {
				k[left], k[i] = k[i], k[left]
				left++
			}
		}

		k[left], k[right] = k[right], k[left]

		SortUni(k[:left], field)
		SortUni(k[left+1:], field)

		return a
	default:
		return false
	}

}

func getFieldValue(a interface{}, field string) interface{} {
	r := reflect.ValueOf(a)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface()
}
