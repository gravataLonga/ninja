package parser

import (
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/token"
)

const (
	_ int = iota
	LOWEST
	ASSIGN
	LOGICAL       // || and &&
	EQUALS        // ==
	LESS_GREATER  // > or <
	SHIFT_BITWISE // >> or <<
	SUM           //+
	BITWISE       // ~, |, &, ^
	PRODUCT       //*
	POW           // **
	POSTFIX       // -- or ++ but in postfix position
	PREFIX        // -X or !X
	CALL          // myFunction(X)
	INDEX
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type fnInfixPrecedence struct {
	fn         infixParseFn
	precedence int
}

type fnPrefixPrecedence struct {
	fn         prefixParseFn
	precedence int
}

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn

	prefixParsePrecedence map[token.TokenType]int
	infixParsePrecedence  map[token.TokenType]int
}

// associativity if 1 then is right, 0, mean left.
var associativity = map[token.TokenType]int{
	token.EXPONENCIAL: 1,
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.prefixParsePrecedence = make(map[token.TokenType]int)
	p.registerPrefix(token.IDENT, p.parseIdentifier, LOWEST)
	p.registerPrefix(token.INT, p.parseIntegerLiteral, LOWEST)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral, LOWEST)
	p.registerPrefix(token.TRUE, p.parseBoolean, LOWEST)
	p.registerPrefix(token.FALSE, p.parseBoolean, LOWEST)
	p.registerPrefix(token.STRING, p.parseString, LOWEST)
	p.registerPrefix(token.BANG, p.parsePrefixExpression, PREFIX)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression, PREFIX)
	p.registerPrefix(token.DECRE, p.parsePrefixExpression, PREFIX)
	p.registerPrefix(token.INCRE, p.parsePrefixExpression, PREFIX)
	p.registerPrefix(token.IF, p.parseIfExpression, LOWEST)
	p.registerPrefix(token.FUNCTION, p.parseFunction, LOWEST)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression, LOWEST)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral, LOWEST)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral, LOWEST)
	p.registerPrefix(token.FOR, p.parseLoopLiteral, LOWEST)
	p.registerPrefix(token.IMPORT, p.parseImport, LOWEST)
	// p.registerPrefix(token.BIT_NOT, p.parsePrefixExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.infixParsePrecedence = make(map[token.TokenType]int)
	p.registerInfix(token.ASSIGN, p.parseInfixAssignExpression, ASSIGN)
	p.registerInfix(token.PLUS, p.parseInfixExpression, SUM)
	p.registerInfix(token.MINUS, p.parseInfixExpression, SUM)
	p.registerInfix(token.SLASH, p.parseInfixExpression, PRODUCT)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression, PRODUCT)
	p.registerInfix(token.EXPONENCIAL, p.parseInfixExpression, POW)
	p.registerInfix(token.MOD, p.parseInfixExpression, PRODUCT)
	p.registerInfix(token.EQ, p.parseInfixExpression, EQUALS)
	p.registerInfix(token.NEQ, p.parseInfixExpression, EQUALS)
	p.registerInfix(token.LT, p.parseInfixExpression, LESS_GREATER)
	p.registerInfix(token.GT, p.parseInfixExpression, LESS_GREATER)
	p.registerInfix(token.GTE, p.parseInfixExpression, LESS_GREATER)
	p.registerInfix(token.LTE, p.parseInfixExpression, LESS_GREATER)
	p.registerInfix(token.AND, p.parseInfixExpression, LOGICAL)
	p.registerInfix(token.OR, p.parseInfixExpression, LOGICAL)
	p.registerInfix(token.BIT_AND, p.parseInfixExpression, BITWISE)
	p.registerInfix(token.BIT_OR, p.parseInfixExpression, BITWISE)
	p.registerInfix(token.BIT_XOR, p.parseInfixExpression, BITWISE)
	p.registerInfix(token.SHIFT_RIGHT, p.parseInfixExpression, SHIFT_BITWISE)
	p.registerInfix(token.SHIFT_LEFT, p.parseInfixExpression, SHIFT_BITWISE)
	p.registerInfix(token.LPAREN, p.parseCallExpression, CALL)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression, INDEX)
	p.registerInfix(token.DOT, p.parseObjectCallExpression, CALL)
	p.registerInfix(token.DOUBLE_COLON, p.parseEnumAccessorExpression, CALL)

	// Postfix but we only change associativity to right
	p.registerInfix(token.INCRE, p.parsePostfixExpression, POSTFIX)
	p.registerInfix(token.DECRE, p.parsePostfixExpression, POSTFIX)

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
	p.newError("no prefix parse function for %s found", t)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn, precedence int) {
	p.prefixParseFns[tokenType] = fn
	p.prefixParsePrecedence[tokenType] = precedence
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn, precedence int) {
	p.infixParseFns[tokenType] = fn
	p.infixParsePrecedence[tokenType] = precedence
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
		p.newError("expected next token to be %s, got %s instead.", t[0], p.peekToken)
		return
	}

	listTokens := ""
	for _, i := range t {
		listTokens += listTokens + " " + fmt.Sprintf("%s", i)
	}

	p.newError("expected next token to be %s, got %s instead.", listTokens, p.peekToken)
}

func (p *Parser) peekPrecedence() int {
	precedence, ok := p.infixParsePrecedence[p.peekToken.Type]
	if !ok {
		return LOWEST
	}

	return precedence
}

// curAssociativity 0 means left associativity
// > 0 means right associativity
func (p *Parser) curAssociativity() int {
	if p, ok := associativity[p.curToken.Type]; ok {
		return p
	}
	return 0
}

func (p *Parser) curPrecedence() int {
	precedence, ok := p.infixParsePrecedence[p.curToken.Type]
	if !ok {
		return LOWEST
	}

	return precedence
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
