package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseFunctionParameters() []ast.Expression {

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return []ast.Expression{}
	}

	p.nextToken()

	identifiers := p.parseParameterWithOptional()

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseParameterWithOptional() []ast.Expression {
	var identifiers []ast.Expression
	isOnRequiredParameters := true

	if p.peekTokenIs(token.ASSIGN) {
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		p.nextToken()
		infix := p.parseInfixExpression(ident)
		identifiers = append(identifiers, infix)
		isOnRequiredParameters = false
	} else {
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		if p.peekTokenIs(token.ASSIGN) {
			ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			p.nextToken()
			infix := p.parseInfixExpression(ident)
			identifiers = append(identifiers, infix)
			isOnRequiredParameters = false
			continue
		}

		if !isOnRequiredParameters {
			p.newError("require arguments must be on declare first")
			return nil
		}

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	return identifiers
}
