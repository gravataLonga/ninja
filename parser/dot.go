package parser

import (
	"github.com/gravataLonga/ninja/ast"
)

func (p *Parser) parseDotExpression(left ast.Expression) ast.Expression {
	dotExpression := &ast.Dot{Token: p.curToken}

	p.nextToken()
	fn := p.parseIdentifier()
	p.nextToken()

	dotExpression.Right = p.parseCallExpression(fn)

	dotExpression.Object = left
	return dotExpression
}
