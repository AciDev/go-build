package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("./1")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	}

	var status bool = false

	go func() {
		err := cmd.Wait()
		if err != nil {

		}
		status = true
	}()

	for status == false {
		var i string
		if status == false {
			fmt.Scanf("%s\n", &i)
			io.WriteString(stdin, i+"\n")
			time.Sleep(150 * time.Millisecond)
		}
	}
}
