package parser

import (
	"coffee/cmd/coffee/internal/ast"
	"coffee/cmd/coffee/internal/lexer"
	"coffee/cmd/coffee/internal/token"
)

type Parser struct {
	l       *lexer.Lexer
	curTok  token.Token
	peekTok token.Token
	errors  []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curTok.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Type {
	case token.USE:
		return p.parseUseStatement()
	case token.IDENT:
		return p.parseExpressionStatement()
	default:
		p.errors = append(p.errors, "Unexpected token: "+p.curTok.Literal)
		return nil
	}
}

func (p *Parser) parseUseStatement() ast.Statement {
	stmt := &ast.UseStatement{Token: p.curTok}
	p.nextToken()
	if p.curTok.Type == token.IDENT {
		stmt.Module = p.curTok.Literal
	} else {
		p.errors = append(p.errors, "Expected module name after 'use'")
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{Token: p.curTok}
	stmt.Expression = p.parseCallExpression()
	return stmt
}

// Solo soporta io.print('...')
func (p *Parser) parseCallExpression() ast.Expression {
	ident := &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	if p.peekTok.Type != token.DOT {
		p.errors = append(p.errors, "Expected '.' after module name")
		return ident
	}

	p.nextToken() // DOT
	p.nextToken() // print
	if p.curTok.Type != token.PRINT {
		p.errors = append(p.errors, "Expected 'print' function")
		return ident
	}

	if p.peekTok.Type != token.LPARENT {
		p.errors = append(p.errors, "Expected '(' after print")
		return ident
	}

	p.nextToken() // LPAREN
	p.nextToken() // argumento
	arg := &ast.StringLiteral{Token: p.curTok, Value: p.curTok.Literal}

	if p.peekTok.Type != token.RPARENT {
		p.errors = append(p.errors, "Expected ')' after argument")
		return arg
	}

	p.nextToken() // RPARENT
	return &ast.CallExpression{
		Token:     p.curTok,
		Function:  ident,
		Arguments: []ast.Expression{arg},
	}
}
