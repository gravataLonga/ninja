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
		if p.peekTokenIs(token.LBRACKET) {
			return p.parseInfixAssignExpression()
		}
		return p.parseExpressionStatement()
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
