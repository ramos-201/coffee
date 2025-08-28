package token

type TokenType string

const (
	USE   TokenType = "USE"
	IDENT TokenType = "IDENT"

	PRINT TokenType = "PRINT"

	LPARENT TokenType = "LPARENT" // (
	RPARENT TokenType = "RPARENT" // )
	DOT     TokenType = "DOT"     // .

	STRING TokenType = "STRING" // 'text'

	EOF TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}
