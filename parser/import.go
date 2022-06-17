package parser

import (
	"ninja/ast"
	"ninja/token"
)

func (p *Parser) parseImport() ast.Expression {
	expression := &ast.Import{Token: p.curToken}

	if !p.expectPeek(token.STRING) {
		return nil
	}

	expression.Filename = p.parseExpression(LOWEST)

	return expression
}
