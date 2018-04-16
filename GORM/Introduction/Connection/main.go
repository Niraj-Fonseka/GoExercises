// package main

// import (
// 	"fmt"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/lib/pq"
// )

// func main() {

// 	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	dbase := db.DB()
// 	defer db.Close()

// 	err = dbase.Ping()

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Println("Database connection established")
// }

package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
	ID        uint   `json:”id”`
	FirstName string `json:”firstname”`
	LastName  string `json:”lastname”`
}

func main() {
	db, _ := gorm.Open("sqlite3", "./gorm.db")
	defer db.Close()
	p1 := Person{FirstName: "John", LastName: "Doe"}
	p2 := Person{FirstName: "Jane", LastName: "Smith"}
	fmt.Println(p1.FirstName)
	fmt.Println(p2.LastName)
}
