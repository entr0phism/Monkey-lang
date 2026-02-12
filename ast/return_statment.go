package ast

import (
	"fmt"

	"monkeylang/token"
)

type ReturnStatment struct {
	Token       token.Token
	ReturnValue Expression
}

func (r *ReturnStatment) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatment) statementNode() {
}

func (r *ReturnStatment) String() string {
	return fmt.Sprintf("%s %v;", r.Token.Literal, r.ReturnValue.String())
}
