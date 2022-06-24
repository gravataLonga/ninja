package parser

import "ninja/ast"

func (p *Parser) parseIdentifier() ast.Expression {
	postfix := p.postfixParseFns[p.peekToken.Type]
	if postfix != nil {
		p.nextToken()
		return postfix()
	}

	return &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}
}
