package ast

import "github.com/schweigerjonas/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // dummy method only included for better errors
}

type Expression interface {
	Node
	expressionNode() // dummy method only included for better errors
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

// Note: the LetStatement and Identifier types _implicitly_ implement the interfaces StatementNode
// and ExpressionNode respectively, by having all of the defined methods in the interface  in their
// method sets

type LetStatement struct { // let <identifier> = <expression>;
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// why is this an expression node?
// -> to keep the number of different node types smaller and since there can also be value-producing
// identifiers (e.g. global constants like PI)
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
