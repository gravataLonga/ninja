package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.IDENT:
		if p.peekTokenIs(token.ASSIGN) {
			return p.parseAssignStatement()
		}
		expr := p.parseExpressionStatement()
		if !p.peekTokenIs(token.ASSIGN) {
			return expr
		}

		p.nextToken()
		p.nextToken()
		assign := &ast.AssignStatement{Token: p.curToken, Left: expr, Right: p.parseExpression(LOWEST)}
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
		return assign
	case token.DELETE:
		return p.parseDeleteStatement()
	case token.VAR:
		return p.parseVarStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.BREAK:
		return p.parseBreakStatement()
	case token.ENUM:
		return p.parseEnum()
	default:
		return p.parseExpressionStatement()
	}
}
