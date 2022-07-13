package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"strconv"
)

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(string(p.curToken.Literal), 64)
	if err != nil {
		p.newError("could not parse %q as float", p.curToken.Literal)
		return nil
	}

	lit.Value = value
	return lit
}
