package main

import (
	animals "GoExercises/InterfaceStuff/animals"
	registration "GoExercises/InterfaceStuff/registration"
	"fmt"
)

func main() {

	dog1 := animals.NewDog("dog1", "female", 5)
	registration.RegisterAnimal(*dog1)

	cat := animals.NewCat("cat1", "male", 2, "yellow")
	registration.RegisterAnimal(cat)

	fmt.Println("---- Printing animals ----")
	registration.PrintAnimals()

	// you cant run this code because the animal needs to satisfy the Animals interface

	// alien := animals.NewAlien("alien1", 200)
	// registration.RegisterAnimal(alien)

}
