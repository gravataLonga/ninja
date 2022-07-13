package parser

import "github.com/gravataLonga/ninja/ast"

func (p *Parser) parseString() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: string(p.curToken.Literal)}
}
