package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
)

func main() {

	vlan := 68

	user := "a++++++++++a"
	pass := "6++++++++++9"
	targethost := "xx.xx.xx.xx:2525"
	//targethost := "ios-xe-mgmt.cisco.com:8181"

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", targethost, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	sess, err := conn.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}

	stdin, _ := sess.StdinPipe()
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	sess.Shell()

	cmds := []string{
		"conf t",
		"int g1/0/7",
		"switchport access vlan " + strconv.Itoa(vlan),
		"shut",
		"no shut",
		"exit",
	}


	for _, cmd := range cmds {
		fmt.Fprintf(stdin, "%s\n", cmd)
	}
	fmt.Fprintf(stdin, "exit\n")
	fmt.Fprintf(stdin, "exit\n")
	sess.Wait()
	sess.Close()
}
