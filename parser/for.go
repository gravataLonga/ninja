package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseLoopLiteral() ast.Expression {
	fr := &ast.ForStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// INITIAL CONDITION

	if p.peekTokenIs(token.VAR) {
		p.nextToken()
		fr.InitialCondition = p.parseVarStatement()
	} else if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	// CONDITION

	if !p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
		fr.Condition = p.parseExpression(LOWEST)
	}

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	// ITERATION

	if !p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		fr.Iteration = p.parseStatement()
		p.nextToken()
	} else {
		p.nextToken()
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	// BODY

	fr.Body = p.parseBlockStatement()

	return fr
}
