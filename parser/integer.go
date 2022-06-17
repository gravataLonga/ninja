package parser

import (
	"ninja/ast"
	"strconv"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.newError("could not parse %q as integer", p.curToken.Literal)
		return nil
	}

	lit.Value = value
	return lit
}
