package main

import "strconv"

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Scanner struct {
	source     string
	tokens     []Token
	start      uint32
	current    uint32
	line       uint32
	ml_comment uint16
}

func MakeScanner(source string) Scanner {
	var s Scanner
	s.source = source
	return s
}

func (scanner *Scanner) ScanTokens() []Token {
	for !scanner.IsAtEnd() {
		scanner.start = scanner.current
		scanner.ScanToken()
	}
	scanner.tokens = append(scanner.tokens, Token{EOF, "", "", scanner.line})
	return scanner.tokens
}

func (scanner *Scanner) IsAtEnd() bool {
	return scanner.current >= uint32(len(scanner.source))
}

func (scanner *Scanner) ScanToken() {
	var c byte = scanner.Advance()
	switch c {
	case '(':
		scanner.AddToken(LEFT_PAREN)
		break
	case ')':
		scanner.AddToken(RIGHT_PAREN)
		break
	case '{':
		scanner.AddToken(LEFT_BRACE)
		break
	case '}':
		scanner.AddToken(RIGHT_BRACE)
		break
	case ',':
		scanner.AddToken(COMMA)
		break
	case '.':
		scanner.AddToken(DOT)
		break
	case '-':
		scanner.AddToken(MINUS)
		break
	case '+':
		scanner.AddToken(PLUS)
		break
	case ';':
		scanner.AddToken(SEMICOLON)
		break
	case '*':
		scanner.AddToken(STAR)
		break
	case '!':
		ttype := BANG
		if scanner.Match('=') {
			ttype = BANG_EQUAL
		}
		scanner.AddToken(ttype)
		break
	case '=':
		ttype := EQUAL
		if scanner.Match('=') {
			ttype = EQUAL_EQUAL
		}
		scanner.AddToken(ttype)
		break
	case '<':
		ttype := LESS
		if scanner.Match('=') {
			ttype = LESS_EQUAL
		}
		scanner.AddToken(ttype)
		break
	case '>':
		ttype := GREATER
		if scanner.Match('=') {
			ttype = GREATER_EQUAL
		}
		scanner.AddToken(ttype)
		break
	case '/':
		if scanner.Match('/') {
			for scanner.Peek() != '\n' && !scanner.IsAtEnd() {
				scanner.Advance()
			}
		} else if scanner.Match('*') {
			scanner.ReadMultilineComment()
		} else {
			scanner.AddToken(SLASH)
		}
		break
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		scanner.line++
		break
	case '"':
		scanner.ReadString()
		break
	default:
		if IsDigit(c) {
			scanner.ReadNumber()
		} else if IsAlpha(c) {
			scanner.ReadIdentifier()
		} else {
			error(scanner.line, "Unexpected character.")
		}
		break
	}
}

func IsDigit(letter byte) bool {
	return letter >= '0' && letter <= '9'
}

func IsAlpha(letter byte) bool {
	lower := letter >= 'a' && letter <= 'z'
	upper := letter >= 'A' && letter <= 'Z'
	underscore := letter == '_'
	return lower || upper || underscore
}

func (scanner *Scanner) ReadMultilineComment() {
	scanner.ml_comment++
	for scanner.ml_comment > 0 {
		if scanner.IsAtEnd() {
			return
		} else if scanner.Peek() == '/' && scanner.PeekNext() == '*' {
			scanner.ml_comment++
			scanner.Advance()
		} else if scanner.Peek() == '*' && scanner.PeekNext() == '/' {
			scanner.ml_comment--
			scanner.Advance()
		}
		scanner.Advance()
	}
}

func (scanner *Scanner) ReadIdentifier() {
	for IsAlpha(scanner.Peek()) {
		scanner.Advance()
	}

	text := scanner.source[scanner.start:scanner.current]
	ttype, is_keyword := keywords[text]
	if !is_keyword {
		ttype = IDENTIFIER
	}
	scanner.AddToken(ttype)
}

func (scanner *Scanner) ReadNumber() {
	for IsDigit(scanner.Peek()) {
		scanner.Advance()
	}

	if scanner.Peek() == '.' && IsDigit(scanner.PeekNext()) {
		scanner.Advance()
		for IsDigit(scanner.Peek()) {
			scanner.Advance()
		}
	}

	as_float, _ := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 64)
	scanner.AddTokenByLiteral(NUMBER, as_float)
}

func (scanner *Scanner) ReadString() {
	for scanner.Peek() != '"' && !scanner.IsAtEnd() {
		if scanner.Peek() == '\n' {
			scanner.line++
		}
		scanner.Advance()
	}

	if scanner.IsAtEnd() {
		error(scanner.line, "Unterminated string.")
		return
	}

	scanner.Advance()

	as_string := scanner.source[scanner.start+1 : scanner.current-1]
	scanner.AddTokenByLiteral(STRING, as_string)
}

// Return current character if not at EOF
func (scanner *Scanner) PeekNext() byte {
	if scanner.current+1 >= uint32(len(scanner.source)) {
		return 0
	}
	return scanner.source[scanner.current+1]
}

func (scanner *Scanner) Peek() byte {
	if scanner.IsAtEnd() {
		return 0
	}
	return scanner.source[scanner.current]
}

// Checks if current character is what we expected. If so, advance.
func (scanner *Scanner) Match(expected byte) bool {
	if scanner.IsAtEnd() {
		return false
	}
	if scanner.source[scanner.current] != expected {
		return false
	}
	scanner.current++
	return true
}

// Returns current character, and steps forward. Like "character++"
func (scanner *Scanner) Advance() byte {
	character := scanner.source[scanner.current]
	scanner.current++
	return character
}

// For adding a token with a type that does not need a literal
func (scanner *Scanner) AddToken(ttype TokenType) {
	scanner.AddTokenByLiteral(ttype, nil)
	// TODO: get rid of interface{} as our literal type. can we just store as lexeme and `eval()` it?
}

// Adds a Token using ttype and literal
func (scanner *Scanner) AddTokenByLiteral(ttype TokenType, literal interface{}) {
	var text string = scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, Token{ttype, text, literal, scanner.line})
}
