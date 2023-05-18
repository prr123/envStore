// envMain
// program to ass a n environmental variable
// author: prr azul software
// date 18 May 2023
//

package main

import (
	"os"
	"fmt"
)

func main() {

	usrStr:="addEnv {env=dir}"

	numArg := len(os.Args)

	if numArg < 2 || numArg>3 {
		fmt.Printf("usage: %s\n", usrStr)
		os.Exit(-1)
	}

	fmt.Printf("env variable: %s\n", os.Args[1])

	tbyt := []byte(os.Args[1])

	pos := -1
	for i:=0; i<len(tbyt); i++ {
		if tbyt[i] == '=' {
			pos = i
			break
		}
	}
	if pos == -1 {
		fmt.Printf("no = char!\n")
		os.Exit(-1)
	}

	envVar := string(tbyt[:pos])
	dirVar := string(tbyt[pos+1:])

	fmt.Printf("envVar: %s dir: %s\n", envVar, dirVar)
}
