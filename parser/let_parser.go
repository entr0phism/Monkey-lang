package parser

import (
	"monkeylang/ast"
	"monkeylang/token"
)

func (p *Parser) parseLetStatememt() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: parse expression

	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	// HACK: (;) handling of let statements
	p.nextToken()

	return stmt
}
