package main

//This will need to be built with the whole package ( go build)

//void AsmMovl();
import "C"

func main() {
	C.AsmMovl()
}
