package main 

import (
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