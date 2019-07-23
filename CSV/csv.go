package main 

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)


func main(){
	file , err := os.Open("input.csv")
	if err != nil {
		log.Println("Cant Open File")
	}
	
	r := csv.NewReader(file)
	for {
		record , err := r.Read()		
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}