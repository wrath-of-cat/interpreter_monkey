package main

import (
	"fmt"
	"monkey_interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome %s, to the Monkey Programming Language interpreter!\n", user.Username)
	fmt.Printf("Writing any amount of code you want!\n")
	repl.Start(os.Stdin, os.Stdout)
}
