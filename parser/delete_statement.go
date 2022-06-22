package parser

import (
	"ninja/ast"
)

func (p *Parser) parseDeleteStatement() *ast.DeleteStatement {
	stmt := &ast.DeleteStatement{}

	stmt.Identifier = p.parseExpression(LOWEST)

	return stmt
}
