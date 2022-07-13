package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseEnum() ast.Statement {
	enum := &ast.EnumStatement{}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	enum.Identifier = p.parseIdentifier()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	branches := map[string]ast.Expression{}

	for !p.peekTokenIs(token.RBRACE) {
		if !p.expectPeek(token.CASE) {
			return nil
		}

		if !p.expectPeek(token.IDENT) {
			return nil
		}

		key := p.curToken.Literal

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}

		_, ok := branches[key]
		if ok {
			p.newError("Fatal error: Cannot redefine identifier %s", key)
			return nil
		}

		branches[key] = value
	}

	p.nextToken()
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	enum.Branches = branches
	return enum
}

func (p *Parser) parseEnumAccessorExpression(left ast.Expression) ast.Expression {
	scope := &ast.ScopeOperatorExpression{
		Token:            p.curToken,
		AccessIdentifier: left,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	ident := p.parseIdentifier()
	scope.PropertyIdentifier = ident
	return scope
}
