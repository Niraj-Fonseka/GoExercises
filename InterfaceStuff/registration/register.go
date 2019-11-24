package registration

import (
	"fmt"
)

type Animal interface {
	Name() string
	Age() int
	Gender() string
}

var (
	AnimalsStore map[string]Animal
)

func init() {
	AnimalsStore = make(map[string]Animal)
}

func RegisterAnimal(animal Animal) {
	AnimalsStore[animal.Name()] = animal
}

func PrintAnimals() {
	for _, v := range AnimalsStore {
		fmt.Printf("Name : %s , Age : %d , Gender : %s \n", v.Name(), v.Age(), v.Gender())
	}
}
