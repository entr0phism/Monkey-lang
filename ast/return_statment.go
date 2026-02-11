package ast

import "monkeylang/token"

type ReturnStatment struct {
	Token       token.Token
	ReturnValue Expression
}

func (r *ReturnStatment) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatment) statementNode() {
}
