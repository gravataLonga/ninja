package parser

import (
	"ninja/ast"
	"ninja/token"
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
		p.newError("expected next token to be %s, got %s instead. %s", token.IDENT, p.curToken.Type, p.l.FormatLineCharacter())
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
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseInfixAssignExpression(left ast.Expression) ast.Expression {
	stmt := &ast.AssignStatement{Token: p.curToken}
	if n, ok := left.(*ast.Identifier); ok {
		stmt.Name = n
	} else if n, ok := left.(*ast.IndexExpression); ok {
		stmt.Name = n
	} else {
		p.nextToken()
		stmt.Value = p.parseExpression(LOWEST)

		p.newError("illegal \"%s\" assignment to \"%s\"", stmt.Value.TokenLiteral(), left.TokenLiteral())
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}
