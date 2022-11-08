package main

import (
	"os"
	"fmt"
	//"strings"
)


func main() { 
	
	args := os.Args
	fmt.Println("Whole Length :", len(args))
	fmt.Println("arg Length :",len(args[1:]))
	fmt.Println("arg List :", args[1:])

	

fmt.Println("\n===================\n")
	
	for i, line := range args{
		fmt.Println(i, line)
	}


	fmt.Println("arg List :", args[1])
	fmt.Println("arg List :", args[2])

}