package ast

import "github.com/orangeseeds/gmonkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

/*
    Statement -> Doesn't produce a value
    example:
	let x = 5
	return 5

    Expression -> Produces a value
    example:
	5
	add(5,5)
*/

/*
   An example let statement to be implemented:
   let a = 5

   @field Name -> holds the identifier of the binding
   @field Value -> holds the expression that produces the value
*/
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
