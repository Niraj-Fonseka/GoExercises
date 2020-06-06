package main

import "fmt"

type Object struct {
	ObjectName  string
	ObjectValue int
}

func NewObject() *Object {
	return &Object{
		ObjectName:  "Name1",
		ObjectValue: 100,
	}
}

func (o *Object) String() string {
	return fmt.Sprintf("Name : %s , Values : %d \n", o.ObjectName, o.ObjectValue)
}
