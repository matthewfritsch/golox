package main

import (
	"fmt"
)

// type LoxError struct {
// 	line uint32
// 	code uint32
// 	msg  string
// }

// func (e *LoxError) Error() string {
// 	return fmt.Sprintf("Error %d - %s", e.code, e.msg)
// }

// func build_error(line uint32, message string) LoxError {
// 	msg := fmt.Sprintf("%s\n\n%s", string(debug.Stack()), message)
// 	return LoxError{line, 1, msg}
// }

// func throw_lox_error(le LoxError) {
// 	// TODO: throw particular error code based on error type
// 	fmt.Printf("Ran into error %s; exiting...\n", le.Error())
// 	os.Exit(int(le.code))
// }

// func throw_error(line uint32, msg string) {
// 	throw_lox_error(build_error(line, msg))
// }

// func throw_internal_error(e error) {
// 	fmt.Printf("Error: %s\n", e.Error())
// }

func error(line uint32, message string) {
	report(line, "", message)
}

func report(line uint32, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
	hadError = true
}
