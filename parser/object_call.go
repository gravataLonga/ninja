package parser

import (
	"ninja/ast"
)

func (p *Parser) parseObjectCallExpression(left ast.Expression) ast.Expression {
	objectCall := &ast.ObjectCall{Token: p.curToken}

	p.nextToken()
	fn := p.parseIdentifier()
	p.nextToken()

	objectCall.Call = p.parseCallExpression(fn)

	objectCall.Object = left
	return objectCall
}
