package models

import (
	"fmt"
	"time"
)

//GetUsertByID gets User by id
//keep autoPreload true if needs to fetch all related data else keep false
func GetHealth() (health string, err error) {
	fmt.Println("Calling DB GetHealth")
	time.Sleep(10 * time.Second)
	return "Health is good !", nil
}
