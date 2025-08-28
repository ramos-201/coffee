package lexer

import (
	"coffee/cmd/coffee/internal/token"
	"unicode"
)

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

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	var tok token.Token
	switch l.ch {
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	case '(':
		tok = token.Token{Type: token.LPARENT, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.RPARENT, Literal: string(l.ch)}
	case '.':
		tok = token.Token{Type: token.DOT, Literal: string(l.ch)}
	case '\'':
		tok = l.readString()
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tok = lookupIdent(ident)
			return tok
		} else {
			tok = token.Token{Type: token.EOF, Literal: ""}
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

func (l *Lexer) readString() token.Token {
	l.readChar()
	start := l.position

	for l.ch != '\'' && l.ch != 0 {
		l.readChar()
	}

	str := l.input[start:l.position]

	return token.Token{Type: token.STRING, Literal: str}
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

func lookupIdent(ident string) token.Token {
	switch ident {
	case "use":
		return token.Token{Type: token.USE, Literal: ident}
	case "print":
		return token.Token{Type: token.PRINT, Literal: ident}
	default:
		return token.Token{Type: token.IDENT, Literal: ident}
	}
}
