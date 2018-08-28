package main

import (
	"syscall/js"
	"strconv"
)

func add(i []js.Value) {

	value1 := js.Global().Get("document").Call("getElementById" , i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById" , i[1].String()).Get("value").String()

	int1 , _ := strconv.Atoi(value1)
	int2 , _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value",int1+int2)
}

func substract(i []js.Value){
	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
	println(js.ValueOf(i[0].Int() - i[1].Int()).String())
}
func registerCallbacks() {
	js.Global().Set("add",js.NewCallback(add))
	js.Global().Set("substract",js.NewCallback(substract))

}

func main(){
	c := make(chan struct{},0)
	println("Go WebAssembly Initialized")
	registerCallbacks()

	<-c
}