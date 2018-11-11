package main

import "fmt"

type User struct {
	FirstName string 
  LastName string
  UserName string
}

var names = []User { 
    User {
        FirstName: "Hee", 
        LastName: "Littrell",
        UserName: "None", 
    },
    User {
        FirstName: "Laurice", 
        LastName: "Doig",
        UserName: "None",  
    },
    User {
        FirstName: "Adelina", 
        LastName: "Chavera",
        UserName: "None",  
    },
    User {
        FirstName: "Chester", 
        LastName: "Quinlan",
        UserName: "None",  
    },
}

func main() {
 
 
 for index, _ := range names {
			indexAddr := &names[index]

			userLogBuilder := fmt.Sprintf("%s_%s", indexAddr.FirstName , indexAddr.LastName)

			indexAddr.UserName = userLogBuilder

	}

fmt.Println(names)
  
}

