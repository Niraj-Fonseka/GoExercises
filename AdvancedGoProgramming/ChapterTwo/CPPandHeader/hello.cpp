// hello.cpp

#include <iostream>


//Extern is to link the symbol of the function follows the C language rules 
extern "C" {
    #include "hello.h"
}


void SayHello(const char* s){
    std::cout << s;
}