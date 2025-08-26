package lexer

type TokenType string

const (
	USE   = "USE"
	IDENT = "IDENT"
	PRINT = "PRINT"

	LPARENT = "LPARENT"
	RPARENT = "RPARENT"
	DOT     = "DOT"

	STRING = "STRING"

	EOF = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}
