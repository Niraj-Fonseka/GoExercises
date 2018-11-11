package Structs

import "time"

//Credentials Struct for storing user auth data
type Credentials struct {
	Username string    `datastore:"description"`
	Password string    `datastore:"description"`
	Created  time.Time `datastore:"created"`
}
