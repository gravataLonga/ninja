package parser

import (
	"ninja/ast"
	"ninja/token"
)

func (p *Parser) parseDeleteStatement() *ast.DeleteStatement {

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt := &ast.DeleteStatement{Token: p.curToken}

	if !p.curTokenIs(token.IDENT) {
		p.newError("expected current token to be %s, got %s instead.", token.IDENT, p.curToken.Type)
		return nil
	}

	stmt.Left = p.parseIdentifier()

	if !p.expectPeek(token.LBRACKET) {
		return nil
	}

	p.nextToken()
	stmt.Index = p.parseExpression(LOWEST)
	p.nextToken()

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
