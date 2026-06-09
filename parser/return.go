package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
		return stmt
	}

	if _, ok := p.prefixParseFns[p.peekToken.Type]; !ok && !p.peekTokenIs(token.SEMICOLON) && !p.peekTokenIs(token.EOF) {
		p.newError("Next token expected to be nil or expression. Got: %s.", p.peekToken)
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
