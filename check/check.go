package check

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

// Check for checking changes
func Check() {
	res, status := argCheck(systemArgs())
	compile(res)
	if status == true {
		validate(res)
		modified := time.Now().Unix()
		for true {
			time.Sleep(100 * time.Millisecond)
			mod := lastModified(res)
			if modified < mod {
				fmt.Println("File has changed")
				compile(res)
				modified = mod
			}
		}
	} else {
		displayHelp()
		os.Exit(0)
	}
}

func systemArgs() []string {
	args := os.Args[1:]
	return args
}

func argCheck(args []string) (response string, status bool) {
	for i, el := range args {
		if el == "-f" || el == "--file" {
			return args[i+1], true
		} else if el == "-h" || el == "--help" {
			displayHelp()
			os.Exit(0)
		}
	}
	return "", false
}

func displayHelp() {
	fmt.Println(`Go-Build is a program to compile and view your program running as you save!

Please use one of our system args

-f, --file			Provide with a filename
-h, --help			Show the help screen
	`)
}

func validate(file string) bool {
	// Checking if user has go
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Go is not installed on this machine")
		os.Exit(0)
	}

	var outString string = string(out)[0:10]

	if outString != "go version" {
		fmt.Println("Go is not installed on this machine")
		os.Exit(0)
	}

	// Checking if the directory exists
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err, dir)
		fmt.Println("Could not correctly resolve the working directory")
		os.Exit(0)
	}

	// Checking if the file exists
	f, err := os.Open(file)
	if err != nil || f == nil {
		fmt.Println(err)
		fmt.Println("File does not exist")
		os.Exit(0)
	}

	return true
}

func lastModified(file string) int64 {
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		fmt.Println("File does not exist")
		os.Exit(0)
	}

	return info.ModTime().Unix()
}

func compile(file string) {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
