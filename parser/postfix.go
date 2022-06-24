package parser

import "ninja/ast"

func (p *Parser) parsePostfixExpression() ast.Expression {
	return &ast.PostfixExpression{
		Token:    p.prevToken,
		Operator: string(p.curToken.Literal),
	}
}
