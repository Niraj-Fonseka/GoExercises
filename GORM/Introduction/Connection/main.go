package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {

	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dbase := db.DB()
	defer db.Close()

	err = dbase.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Database connection established")
}
