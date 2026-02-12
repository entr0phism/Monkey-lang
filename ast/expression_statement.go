package ast

import (
	"bytes"

	"monkeylang/token"
)

/*
Expression can also be a statement.
It's not a distinct statement, but a wrapper that consists of 1 expression
This is needed because we can write code like this
let x = 5;
x + 10; // this is not and expression, but a statement and will not return a value
*/

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {
}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

type PrefixExpression struct {
	Token    token.Token // The prefix token ex: !, -
	Operator string
	Right    Expression
}

func (p *PrefixExpression) expressionNode()      {}
func (p *PrefixExpression) TokenLiteral() string { return p.Token.Literal }
func (p *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")

	return out.String()
}
