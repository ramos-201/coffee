package lexer

import "unicode"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	var tok Token
	switch l.ch {
	case 0:
		tok = Token{Type: EOF, Literal: ""}
	case '(':
		tok = Token{Type: LPARENT, Literal: string(l.ch)}
	case ')':
		tok = Token{Type: RPARENT, Literal: string(l.ch)}
	case '.':
		tok = Token{Type: DOT, Literal: string(l.ch)}
	case '\'':
		tok = l.readString()
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tok = lookupIdent(ident)
			return tok
		} else {
			tok = Token{Type: EOF, Literal: ""}
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readString() Token {
	l.readChar()
	start := l.position

	for l.ch != '\'' && l.ch != 0 {
		l.readChar()
	}

	str := l.input[start:l.position]

	return Token{Type: STRING, Literal: str}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func lookupIdent(ident string) Token {
	switch ident {
	case "use":
		return Token{Type: USE, Literal: ident}
	case "print":
		return Token{Type: PRINT, Literal: ident}
	default:
		return Token{Type: IDENT, Literal: ident}
	}
}
