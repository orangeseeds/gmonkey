package parser

import (
	"fmt"

	"github.com/orangeseeds/gmonkey/ast"
	"github.com/orangeseeds/gmonkey/lexer"
	"github.com/orangeseeds/gmonkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so that curToken and peekToken are both set
	/*
		    when a parser is created both curToken and peekToken are null,
		    on first p.nextToken()
			curToken = peekToken i.e. null
			peekToken = first token in the input

		    on second p.nextToken()
			curToken = peekToken i.e first token in the input
			peekToken = second token in the input
	*/
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		// if statement is found, curToken will already be at the end of the statement by the time it reaches here
		// else we will go to the next immediate token
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	// in expectPeek if true then position shifts to next token else remain in same place
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	//For now we are skippin the expressions util we encounter a semicolon
	// i.e. in a statement like let a = 5 ; -> we are skipping the expression 5.

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {

	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {

	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
