package parser

import (
	"monkeylang/ast"
	"monkeylang/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatment {
	stmt := &ast.ReturnStatment{Token: p.curToken}

	p.nextToken()

	// TODO: parse expression

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
