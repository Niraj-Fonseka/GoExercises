package main


//This will need to be built with the whole package ( go build)


//void SayHello(const char* s);
import "C"

func main() {
    C.SayHello(C.CString("Hello, World\n"))
}