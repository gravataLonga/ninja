package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.IDENT:
		// var a = "ola";
		if p.peekTokenIs(token.ASSIGN) {
			return p.parseAssignStatement()
		}
		return p.parseExpressionOrAssignStatement()
	case token.DELETE:
		return p.parseDeleteStatement()
	case token.VAR:
		return p.parseVarStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.CONTINUE:
		return p.parseContinueStatement()
	case token.BREAK:
		return p.parseBreakStatement()
	case token.ENUM:
		return p.parseEnum()
	default:
		return p.parseExpressionOrAssignStatement()
	}
}

func (p *Parser) parseExpressionOrAssignStatement() ast.Statement {
	expr := p.parseExpressionStatement()
	if !p.peekTokenIs(token.ASSIGN) {
		return expr
	}

	switch expr.Expression.(type) {
	case *ast.Identifier, *ast.IndexExpression, *ast.Dot:
		p.nextToken()
		tokenAssign := p.curToken
		p.nextToken()
		assign := &ast.AssignStatement{Token: tokenAssign, Left: expr, Right: p.parseExpression(LOWEST)}
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
		return assign
	default:
		p.nextToken() // '='
		p.nextToken() // RHS start
		rhs := p.parseExpression(LOWEST)
		p.newError("illegal %q assignment to %q", rhs.TokenLiteral(), expr.Expression.TokenLiteral())
		return nil
	}
}
