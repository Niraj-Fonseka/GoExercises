package main 

import (
	"fmt"
)

func main(){
	fmt.Println(recurse(5))
}




func recurse( n int ) int {
	
	if n == 1 {
		return 1
	}else{
		return recurse(n-1) + 1
	}

	
}