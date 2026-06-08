package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseFunctionParameters() []*ast.Parameter {

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return []*ast.Parameter{}
	}

	p.nextToken()

	identifiers := p.parseParameterWithOptional()

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseParameterWithOptional() []*ast.Parameter {
	var identifiers []*ast.Parameter
	isOnRequiredParameters := true

	if p.peekTokenIs(token.ASSIGN) {
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		p.nextToken()
		p.nextToken()
		option := p.parseExpression(LOWEST)
		parameter := &ast.Parameter{Token: p.curToken, Name: ident, Default: option}
		identifiers = append(identifiers, parameter)
		isOnRequiredParameters = false
	} else {
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		parameter := &ast.Parameter{Token: p.curToken, Name: ident}
		identifiers = append(identifiers, parameter)
	}

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		if p.peekTokenIs(token.ASSIGN) {
			ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			p.nextToken()
			p.nextToken()
			option := p.parseExpression(LOWEST)
			parameter := &ast.Parameter{Token: p.curToken, Name: ident, Default: option}
			identifiers = append(identifiers, parameter)
			isOnRequiredParameters = false
			continue
		}

		if !isOnRequiredParameters {
			p.newError("require arguments must be on declare first")
			return nil
		}

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		parameter := &ast.Parameter{Token: p.curToken, Name: ident}
		identifiers = append(identifiers, parameter)
	}

	return identifiers
}
