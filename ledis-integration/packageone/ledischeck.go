package packageone

import (
	db "go_exercises/ledis-integration/ledisDB"
)

func SetSomeMoreVals() {
	key := []byte("ThisKeyTwo")
	value := []byte("ValueTwo")

	db.DBLEDIS.Set(key, value)
}
