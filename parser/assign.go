package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		p.nextToken()
		for !p.curTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	if p.curTokenIs(token.ASSIGN) {
		p.newError("expected next token to be %s, got %s instead.", token.IDENT, p.curToken)
		return nil
	}

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {

	stmt := &ast.AssignStatement{Token: p.curToken}
	stmt.Left = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Right = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseInfixAssignExpression(left ast.Expression) *ast.AssignStatement {
	expression := &ast.AssignStatement{
		Token: p.curToken,
		Left:  left,
	}

	p.nextToken()
	expression.Right = p.parseExpression(ASSIGN)
	return expression
}
