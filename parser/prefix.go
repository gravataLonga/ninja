package parser

import "ninja/ast"

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: string(p.curToken.Literal),
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}
