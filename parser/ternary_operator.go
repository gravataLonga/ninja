package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseTernaryOperator(left ast.Expression) ast.Expression {
	ternary := &ast.TernaryOperatorExpression{Token: p.curToken}
	ternary.Condition = left
	p.nextToken()
	ternary.Consequence = p.parseExpression(LOWEST)
	if !p.expectPeek(token.COLON) {
		return nil
	}
	p.nextToken()
	ternary.Alternative = p.parseExpression(LOWEST)
	return ternary
}

func (p *Parser) parseElvisOperator(left ast.Expression) ast.Expression {
	elvisOperator := &ast.ElvisOperatorExpression{Token: p.curToken}
	elvisOperator.Left = left
	p.nextToken()
	elvisOperator.Right = p.parseExpression(LOWEST)
	return elvisOperator
}
