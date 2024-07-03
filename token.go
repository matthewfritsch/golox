package main

import (
	"fmt"
	"strconv"
)

type Object interface {
	toString() string
}

type Token struct {
	ttype   TokenType
	lexeme  string
	literal interface{}
	line    uint32
}

func (token *Token) ToString() string {
	literal_msg := "<no-literal>"
	if stringval, ok := token.literal.(string); ok {
		literal_msg = stringval
	} else if floatval, ok := token.literal.(float64); ok {
		literal_msg = strconv.FormatFloat(floatval, 'E', -1, 64)
	}
	return fmt.Sprintf("%v %s %s", token.ttype, string(token.lexeme), literal_msg)
}

// func (token *Token) ToLiteral() interface{} {
// 	if token.ttype == NUMBER {
// 		num, err := strconv.ParseInt(token.lexeme, 10, 64)
// 		if err == nil {
// 			error(0, fmt.Sprintf("Could not turn %s into integer", num))
// 		}
// 		return num
// 	}
// 	if token.ttype == STRING {
// 		return token.lexeme
// 	}
// 	return nil
// }
