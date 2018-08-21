package main 

import (
	"unsafe"
	"reflect"
	"fmt"
)


func AppendingToTheStart(){
	var a []int
	a = append(a, 1)               
	a = append(a, 1, 2, 3)        
	a = append(a, []int{1,2,3}...) 


	a = append([]int{0}, a...)        // 在开头添加1个元素
	a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片

	fmt.Println(a)
}

func main(){
	AppendingToTheStart()
}


func bytes2str(s [] byte) (p string) {
	data := make([]byte, len(s))

	for i , c := range s {
		data[i] = c 
	}

	hdr := (*reflect.StringHeader) (unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)

	return p
}