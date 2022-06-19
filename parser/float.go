package parser

import (
	"ninja/ast"
	"strconv"
)

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		p.newError("could not parse %q as float", p.curToken.Literal)
		return nil
	}

	lit.Value = ast.FloatSmall(value, 10)
	return lit
}
