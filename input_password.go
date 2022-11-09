package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "syscall"

    "golang.org/x/term"
)

func main() {
    user, pswd, _ := userinfo() // Calling the func userinfo()
  
    //rendering the result
    fmt.Printf("Username: %s, Password: %s\n", user, pswd)
    fmt.Println("========================================")
    fmt.Println(user)
    fmt.Println(pswd)
  
}

func userinfo() (string, string, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Input User: ")
    user, err := reader.ReadString('\n')
    if err != nil {
        return "", "", err
    }

    fmt.Print("Input Password: ")
    bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        return "", "", err
    }

    pswd := string(bytePassword)
  
    return strings.TrimSpace(user), strings.TrimSpace(pswd), nil
}
