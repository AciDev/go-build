package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/acidev/go-build/check"
)

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go awaitingInterupt(sigs)

	check.Check()
}

func awaitingInterupt(sigs chan os.Signal) {
	sig := <-sigs
	fmt.Println()
	fmt.Println(sig)
	os.Exit(0)
}
