package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/token"
)

func (p *Parser) parseFunction() ast.Expression {

	if !p.peekTokenIs(token.LPAREN) && !p.peekTokenIs(token.IDENT) {
		p.peekError(token.LPAREN, token.IDENT)
		return nil
	}

	// Literal function, e.g.: var add = function() {};
	if !p.peekTokenIs(token.IDENT) {
		return p.parseFunctionLiteral()
	}

	// Normal Function, e.g.: function add() {};

	lit := &ast.FunctionLiteral{}
	p.nextToken()
	lit.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}
