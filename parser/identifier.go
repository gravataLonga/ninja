package parser

import (
	"ninja/ast"
	"ninja/token"
)

func (p *Parser) parseIdentifier() ast.Expression {
	if p.peekTokenIs(token.ASSIGN) {
		p.peekError(token.IDENT)
		return &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}
	}

	return &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}
}
