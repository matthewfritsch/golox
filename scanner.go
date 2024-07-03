package main

type Scanner struct {
	source  string
	tokens  []Token
	start   uint32
	current uint32
	line    uint32
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
		scanner.String()
		break
	default:
		error(scanner.line, "Unexpected character.")
		break
	}
}

func (scanner *Scanner) String() {
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

	value := scanner.source[scanner.start+1 : scanner.current-1]
	scanner.AddTokenByLiteral(STRING, value)
}

// Return current character if not at EOF
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
	scanner.AddTokenByLiteral(ttype, "")
}

// Adds a Token using ttype and literal
func (scanner *Scanner) AddTokenByLiteral(ttype TokenType, literal string) {
	var text string = scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, Token{ttype, text, literal, scanner.line})
}
