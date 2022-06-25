package parser

import (
	"ninja/ast"
	"ninja/token"
)

func (p *Parser) parseBreakStatement() *ast.BreakStatement {
	breakStmt := &ast.BreakStatement{Token: p.curToken}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return breakStmt
}