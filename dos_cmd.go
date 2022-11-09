/*
 * Copyright (c) 2022 Arvin Acosta
 *
 * Permission to use, copy, modify, and distribute this software for any purpose
 * with or without fee is hereby granted, provided that the above copyright notice
 * and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
 * REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
 * FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT
 * OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, 
 * DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS 
 * ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

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
choose which dos command you wish to run @ Line 28 & Line 31

*/
