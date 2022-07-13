package parser

import "github.com/gravataLonga/ninja/ast"

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: string(p.curToken.Literal),
		Left:     left,
	}

	associativity := p.curAssociativity()

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence - associativity)
	return expression
}
