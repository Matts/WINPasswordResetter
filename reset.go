package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	system32 = os.Getenv("systemroot") + "\\system32"
	oskbak   = system32 + "\\osk.bak"
	oskexe   = system32 + "\\osk.exe"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("An error has occurred", r)
		}
	}()
	fmt.Println("Hello World, Welcome to WINPasswordResetter.\nPlease specify the account of which the password should be removed")
	var input string
	fmt.Scanln(&input)

	// runs 'net user <user> ""'
	command := exec.Command("net", "USER", input, "")
	err := command.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("The password of", input, "has been removed")

	if _, err := os.Stat(oskbak); err == nil {
		// runs ping for delay, then deletes the current exe file. copies the backup eventually pausing
		comm := exec.Command("cmd.exe", "/C ping 127.0.0.1 && del "+oskexe+" && copy "+oskbak+" "+oskexe+" && pause")
		err = comm.Start()
		comm.Process.Release()
		if err != nil {
			fmt.Println(err)
		}
	}
}
