package main

import (
	"fmt"
	
	"strconv"

)

func main() {
	intf := "g1/0/9"
	vlan := 68

	//args := os.Args

	cmds := []string{
		"conf t",
		"int "+ intf,
		"switchport access vlan " + strconv.Itoa(vlan),
		"shut",
		"no shut",
		"exit",
	}


	for _, cmd := range cmds {
		fmt.Println(cmd)
	}

}
