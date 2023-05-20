// envMain
// program to ass a n environmental variable
// author: prr azul software
// date 18 May 2023
//

package main

import (
	"os"
	"fmt"
	"strings"
	"os/exec"
)

func main() {

	usrStr:="addEnv {env=dir}"

	numArg := len(os.Args)

	if numArg < 2 || numArg>2 {
		fmt.Printf("numarg: %d\nusage: %s\n",numArg, usrStr)
		os.Exit(-1)
	}

	fmt.Printf("env variable: %s\n", os.Args[1])

	tbyt := []byte(os.Args[1])

	pos := -1
	for i:=0; i<len(tbyt); i++ {
		if tbyt[i] == '='{
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

	envVal := os.Getenv(envVar)
	if len(envVal) > 0 {
		fmt.Printf("key: %s has value: %s\n", envVar,envVal)
		os.Exit(-1)
	}

	fmt.Printf("envVar: %s dir: %s\n", envVar, dirVar)

	homeDir := os.Getenv("HOME")
	fmt.Printf("home dir: %s\n", homeDir)

/*
	filInfo, err := os.Stat(homeDir + "/.bashrc")
	if err != nil {
		fmt.Printf("bashrc info: %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("size: %d \n", filInfo.Size())
*/

	bashFil, err := os.OpenFile(homeDir +"/.bashrc", os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("bashrc open: %v\n", err)
		os.Exit(-1)
	}
	defer bashFil.Close()

//	size := filInfo.Size()
	_, err = bashFil.Seek(0, 2)
	if err != nil {
		fmt.Printf("bashrc seek: %v\n", err)
		os.Exit(-1)
	}
//	n, err := bashFil.WriteString("export util=\"" + homeDir + "/go/src/utilities\"\n")
	_, err = bashFil.WriteString("export " + envVar + "=\"" + dirVar + "\"\n")
	if err != nil {
		fmt.Printf("bashrc write: %v\n", err)
		os.Exit(-1)
	}

	Native()

	os.Exit(1)

	err = os.Setenv(envVar, dirVar)
	if err != nil {
		fmt.Printf("Setenv: %v\n", err)
		os.Exit(-1)
	}

   for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Printf("var: %s val: %s\n", pair[0], pair[1])
    }
//	fmt.Printf("written %d\n", n)
	fmt.Println("*** success ***")
}


func Native() {
    cmd, err := exec.Command("source", "$HOME/.bashrc").Output()
    if err != nil {
		fmt.Printf("exec Command: %v\n", err)
		return
    }
    fmt.Printf("bashrc exec: %s\n", string(cmd))
    return
}
