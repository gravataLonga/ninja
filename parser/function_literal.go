package parser

import (
	"ninja/ast"
	"ninja/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if !p.peekTokenIs(token.LPAREN) && !p.peekTokenIs(token.IDENT) {
		p.peekError(token.LPAREN, token.IDENT)
		return nil
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}
