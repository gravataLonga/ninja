package parser

import (
	"ninja/ast"
	"ninja/token"
)

func (p *Parser) parseIdentifier() ast.Expression {
	postfix := p.postfixParseFns[p.peekToken.Type]
	if postfix != nil {
		p.nextToken()
		return postfix()
	}

	if p.peekTokenIs(token.ASSIGN) {
		p.peekError(token.IDENT)
		return &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}
	}

	return &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}
}
