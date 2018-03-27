package main

import (
	"fmt"
	db "go_exercises/ledis-integration/ledisDB"
	setValues "go_exercises/ledis-integration/packageone"
)

func main() {

	key := []byte("ThisKey")
	value := []byte("Value")

	db.DBLEDIS.Set(key, value)

	out, _ := db.DBLEDIS.Get(key)
	fmt.Println(string(out))

	setValues.SetSomeMoreVals()

	key2 := []byte("ThisKeyTwo")

	out2, _ := db.DBLEDIS.Get(key2)

	fmt.Println(string(out2))
}
