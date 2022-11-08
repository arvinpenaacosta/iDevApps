package main

import (
	"os"
	"fmt"
	//"strings"
)


func main() {
	
	args := os.Args
fmt.Println(len(args))
fmt.Println(len(args[1:]))
fmt.Println(args[1:])
fmt.Println("===================")
	
	for i, line := range args{
		fmt.Println(i, line)
	}



fmt.Println("===================")

	cmds := []string{
		"conf t",
		"int g1/0/7",
		"switchport access vlan 68",
		"shut",
		"no shut",
		"exit",

	}	
 
	for _, cmd := range cmds{
		//fmt.Fprintf(stdin, "%s\n",cmd)
		fmt.Println(cmd)
	}










/*
fmt.Println("===================")
	var s, sep string
	for i := 1; i < len(args[1:]); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	cmds := strings.Split(s, ",")
	fmt.Println(cmds)
*/






 //   str := "strawberry, blueberry, raspberry"
 //   str := args[1:]
//    fmt.Printf("strings.SplitAfter(): %#v\n", strings.SplitAfter(str, ", "))

 //   str := "strawberry, blueberry, raspberry"
 //   fmt.Printf("strings.Split(): %#v\n", strings.Split(str, ", "))	
}