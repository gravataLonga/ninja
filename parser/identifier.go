package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseIdentifier() ast.Expression {
	if p.peekTokenIs(token.ASSIGN) {
		p.peekError(token.IDENT)
		return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
