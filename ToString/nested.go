package main

import "fmt"

type NestedObject struct {
	NestedObjectName string
	Object           *Object
}

func NewNestedObject(obj *Object) *NestedObject {
	return &NestedObject{
		NestedObjectName: "NestedName1",
		Object:           obj,
	}
}

func (n *NestedObject) String() string {
	return fmt.Sprintf("NestedObjectName : %s , Object : %v \n", n.NestedObjectName, n.Object)
}
