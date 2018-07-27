package main

import (
	"fmt"
	"os"
	)



func Init(){
	fmt.Println("Hello Welcome to Niraj CLI")
}

func main(){
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]	

    arg := os.Args[3]
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)	

}
