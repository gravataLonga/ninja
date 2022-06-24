package parser

import (
	"fmt"
	"ninja/ast"
	"ninja/lexer"
	"ninja/token"
)

const (
	_ int = iota
	LOWEST
	ASSIGN
	LOGICAL      // || and &&
	EQUALS       // ==
	LESS_GREATER // > or <

	SUM     //+
	PRODUCT //*
	PREFIX  // -X or !X
	CALL    // myFunction(X)
	INDEX
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	prevToken token.Token
	curToken  token.Token
	peekToken token.Token

	prefixParseFns  map[token.TokenType]prefixParseFn
	infixParseFns   map[token.TokenType]infixParseFn
	postfixParseFns map[token.TokenType]postfixParseFn
}

type (
	prefixParseFn  func() ast.Expression
	infixParseFn   func(ast.Expression) ast.Expression
	postfixParseFn func() ast.Expression
)

var precedences = map[token.TokenType]int{
	token.ASSIGN:   ASSIGN,
	token.EQ:       EQUALS,
	token.NEQ:      EQUALS,
	token.LT:       LESS_GREATER,
	token.GT:       LESS_GREATER,
	token.GTE:      LESS_GREATER,
	token.LTE:      LESS_GREATER,
	token.OR:       LOGICAL,
	token.AND:      LOGICAL,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.DECRE:    SUM,
	token.INCRE:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.STRING, p.parseString)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.DECRE, p.parsePrefixExpression)
	p.registerPrefix(token.INCRE, p.parsePrefixExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunction)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)
	p.registerPrefix(token.FOR, p.parseLoopLiteral)
	p.registerPrefix(token.IMPORT, p.parseImport)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.ASSIGN, p.parseInfixAssignExpression)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NEQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.GTE, p.parseInfixExpression)
	p.registerInfix(token.LTE, p.parseInfixExpression)
	p.registerInfix(token.AND, p.parseInfixExpression)
	p.registerInfix(token.OR, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)

	p.postfixParseFns = make(map[token.TokenType]postfixParseFn)
	p.registerPostfix(token.INCRE, p.parsePostfixExpression)
	p.registerPostfix(token.DECRE, p.parsePostfixExpression)

	// we set curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) newError(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	p.errors = append(p.errors, s)
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	p.newError("no prefix parse function for %d found", t)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) registerPostfix(tokenType token.TokenType, fn postfixParseFn) {
	p.postfixParseFns[tokenType] = fn
}

func (p *Parser) curTokenIs(tok token.TokenType) bool {
	return p.curToken.Type == tok
}

func (p *Parser) peekTokenIs(tok token.TokenType) bool {
	return p.peekToken.Type == tok
}

func (p *Parser) peekTokenAny(toks ...token.TokenType) bool {
	for _, t := range toks {
		if p.peekToken.Type == t {
			return true
		}
	}
	return false
}

func (p *Parser) expectPeek(tok token.TokenType) bool {
	if p.peekTokenIs(tok) {
		p.nextToken()
		return true
	}
	p.peekError(tok)
	return false
}

func (p *Parser) peekError(t ...token.TokenType) {

	if len(t) == 1 {
		p.newError("expected next token to be %s, got %s instead. %s", t[0], p.peekToken.Type, p.l.FormatLineCharacter())
		return
	}

	listTokens := ""
	for _, i := range t {
		listTokens += listTokens + " " + fmt.Sprintf("%s", i)
	}

	p.newError("expected next token to be %s, got %s instead. %s", listTokens, p.peekToken.Type, p.l.FormatLineCharacter())
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) nextToken() {
	p.prevToken = p.curToken
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
