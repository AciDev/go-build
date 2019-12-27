package main

import "fmt"

func main() {
	fmt.Println("Hello, What's your favorite number?")
	var i string
	fmt.Scanf("%s\n", &i)
	fmt.Println("Ah I like ", i, " too.")
	fmt.Println("Hello, What's your favorite number?")
	fmt.Scanf("%s\n", &i)
	fmt.Println("Ah I like ", i, " too.")
}
