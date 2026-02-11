package ast

import "monkeylang/token"

/*
*** Represents the let statement ***
apart of the token itself we have 2 more fields
Name -> Identifier that implements Expression, as the Identifier (variable) could also produce the value .
Ex: let x = valueProducingIdentifier
Value -> could be a Literal value or an Expression (we implemented the Expression)
*/
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
