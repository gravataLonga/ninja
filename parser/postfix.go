package parser

import "github.com/gravataLonga/ninja/ast"

func (p *Parser) parsePostfixExpression(left ast.Expression) ast.Expression {

	postfix := &ast.PostfixExpression{Left: left}
	postfix.Token = p.curToken
	postfix.Operator = p.curToken.Literal

	return postfix
}
