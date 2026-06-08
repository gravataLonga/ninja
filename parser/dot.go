package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseDotExpression(left ast.Expression) ast.Expression {
	dotExpression := &ast.Dot{Token: p.curToken}

	p.nextToken()
	if !p.curTokenIs(token.IDENT) {
		p.newError("expected property/method name after '.', got %q", p.curToken.Literal)
		return nil
	}

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	dotExpression.Right = ident
	dotExpression.Object = left

	return dotExpression
	// dotExpression := &ast.Dot{Token: p.curToken}
	// p.nextToken()
	// dotExpression.Right = p.parseExpression(DOT)
	// dotExpression.Object = left
	// return dotExpression
}
