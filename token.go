package main

import "fmt"

type Object interface {
	toString() string
}

type Token struct {
	ttype  TokenType
	lexeme string
	line   uint32
}

func (token *Token) ToString() string {
	return fmt.Sprintf("%s %s %s", string(token.ttype), token.lexeme)
}

func (token *Token) ToLiteral() interface{} {
	if true {
		return "nice"
	}
	return 1
}
