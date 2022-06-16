package main

import (
	"fmt"
	"ninja/repl"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Tuga Programming Language\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
