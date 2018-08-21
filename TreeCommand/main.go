package main

import (
	"os"
	"fmt"
	"log"
	"path/filepath"
)




func main(){

	//reading the args 
	//If no parameters need to read the directories in the current directory
	args := []string{"."}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}



	for _ , arg := range args {
		err := tree(arg)
		if err != nil {
			log.Printf("Tree of the %s : %v \n", arg , err)
		}
	}
}

//We want to print the current path and their children
func tree (root string ) error {
	err := filepath.Walk(root , func(path string , fi os.FileInfo , err error) error{

		if err != nil {
			return err
		}


		//Skipping all the dotfiles and its directories
		if fi.Name()[0] == '.' {
			return filepath.SkipDir
		}
		fmt.Println(path)
		return nil
	})
	return err
}


