package main

import "os"

func main() {
	var args []string = os.Args[1:]
	if len(args) > 1 {
		println("Usage: golox [script]")
		os.Exit(64)
	}
	if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}
