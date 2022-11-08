package main

import (

	"os"
	"os/exec"
)

func main() {

  	// Show Product Key
  	dos_cmd := "wmic path SoftwareLicensingService get OA3xOriginalProductKey"

	// Show Serial Number
  	//dos_cmd := "wmic bios get serialnumber"
    
  	cmdr := exec.Command("cmd", "/c", dos_cmd)  //copy read file to myfile.txt 
	cmdr.Stdout = os.Stdout
  	cmdr.Run()
}

/*
choose which dos command you wish to run @ Line 12 & Line 15

*/
