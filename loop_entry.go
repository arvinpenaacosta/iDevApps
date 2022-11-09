package main

import (
        "fmt"
        "os"
        "os/exec"


    "bufio"
    "strings"
    "syscall"

    "golang.org/x/term"
)

func main() {

    var MainLoop = true
    var SubLoop = true

    clscreen()
    logo2()

    user, pswd, _ := userinfo()

    //fmt.Printf("Username: %s, Password: %s\n", user, password)

    fmt.Println(user)
    fmt.Println(pswd)


    main := ""
    for MainLoop {
        clscreen()
        logo1()
        fmt.Println(user)
        fmt.Println(pswd)
        fmt.Println("======================")
        fmt.Print("Enter Port : ")
        fmt.Scan(&main)
        if main == "exit" {
            //MainLoop = false    
            //clscreen()
            break
        }
        MainLoop = true
        SubLoop = true 

        sub := ""
        for SubLoop {
            fmt.Print("Enter Interface : ")
            fmt.Scan(&sub)
            if sub == "exit" {
                clscreen()
                break
            }
        }
    }

   fmt.Println("program exit")

}

func userinfo() (string, string, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter Username: ")
    username, err := reader.ReadString('\n')
    if err != nil {
        return "", "", err
    }

    fmt.Print("Enter Password: ")
    bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        return "", "", err
    }

    password := string(bytePassword)
    return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func clscreen(){
    cmdr := exec.Command("cmd", "/c", "cls")  
    cmdr.Stdout = os.Stdout
    cmdr.Run()
}

func logo(){
    
fmt.Println(`  ___   _____  __ __  __  __ __ `)
fmt.Println(` /   \ |     \|  |  ||  ||  \  |`)
fmt.Println(`|  |  ||  |  ||  |  ||  ||     |`)
fmt.Println(`|     ||     /|  |  ||  ||     |`)
fmt.Println(`|__|__||__\__\ \___/ |__||__\__|`)
fmt.Println(`         microservices          `)
fmt.Println(`                                `)

}

func logoMe(){

fmt.Println(`   _____              .__        `)
fmt.Println(`  /  _  \__________  _|__| ____  `)
fmt.Println(` /  /_\  \_  __ \  \/ /  |/    \ `)
fmt.Println(`/    |    \  | \/\   /|  |   |  \`)
fmt.Println(`\____|__  /__|    \_/ |__|___|  /`)
fmt.Println(`        \/                    \/ `)
}


func logo1(){
    
//https://edukits.co/text-art/

fmt.Println(` _ ____              _                    `)
fmt.Println(`(_)  _ \  _____   __/ \   _ __  _ __  ___ `)
fmt.Println(`| | | | |/ _ \ \ / / _ \ | '_ \| '_ \/ __|`)
fmt.Println(`| | |_| |  __/\ V / ___ \| |_) | |_) \__ \`)
fmt.Println(`|_|____/ \___| \_/_/   \_\ .__/| .__/|___/`)
fmt.Println(`                         |_|   |_|        `)
fmt.Println(` Clear Port & Change VLAN - microservices `)
fmt.Println("==========================================")

}

func logo2(){
    
//https://edukits.co/text-art/

fmt.Println(`                                                   `)
fmt.Println(` _   _      _   __  __                     ___ ___ `)
fmt.Println(`| \ | | ___| |_|  \/  | ___   ___  _ __   |_ _|_ _|`)
fmt.Println(`|  \| |/ _ \ __| |\/| |/ _ \ / _ \| '_ \   | | | | `)
fmt.Println(`| |\  |  __/ |_| |  | | (_) | (_) | | | |  | | | | `)
fmt.Println(`|_| \_|\___|\__|_|  |_|\___/ \___/|_| |_| |___|___|`)
fmt.Println(`                                                   `)
fmt.Println("===================================================")
}

