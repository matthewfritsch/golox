package main

import (
	"bufio"
	"fmt"
	"os"
)

var hadError bool = false

func runFile(s string) {
	bytes, err := os.ReadFile(s)
	if err == nil {
		fmt.Printf("Failed to read file '%s'. Exiting...", s)
		os.Exit(1)
	}
	run(string(bytes))
	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if hadError {
			print("# ")
			hadError = false
		} else {
			print("> ")
		}
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Printf("Prompted with '%s'(%d chars)\n", line, len(line))
		run(line)
		if line == "error" {
			hadError = true
		}
	}
}

func run(source string) {
	scanner := Scanner{source}
	var tokens []Token = scanner.ScanTokens()

	for _, token := range tokens {
		println(token)
	}
}
