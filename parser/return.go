package parser

import (
	"ninja/ast"
	"ninja/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
		return stmt
	}

	// @todo probably i'm doing something wrong.
	if p.peekTokenAny(token.VAR, token.RETURN, token.DECRE, token.INCRE, token.NEQ, token.PLUS, token.MINUS, token.LTE, token.LT, token.GT, token.GTE) {
		p.newError("Next token expected to be nil or expression. Got: %s. %s", p.peekToken.Type, p.l.FormatLineCharacter())
		p.nextToken()
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
		return nil
	}

	p.nextToken()
	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
