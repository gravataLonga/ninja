package lexer

import (
	"fmt"
	"ninja/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte

	lineNumber              int
	characterPositionInLine int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, lineNumber: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1

	l.keepTrackLineAndCharPosition()

}

func (l *Lexer) keepTrackLineAndCharPosition() {
	if l.ch == '\n' {
		l.lineNumber += 1
		l.characterPositionInLine = 0
	} else {
		l.characterPositionInLine += 1
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newTokenFromString(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		if l.peekChar() == '/' {
			l.skipSingleLineComment()
			return l.NextToken()
		}

		if l.peekChar() == '*' {
			l.skipMultiLineComment()
			return l.NextToken()
		}
		tok = newToken(token.SLASH, l.ch)
	case '&':
		if l.peekChar() != '&' {
			tok = newToken(token.ILLEGAL, l.ch)
			return tok
		}
		ch := l.ch
		l.readChar()
		tok = newTokenFromString(token.AND, string(ch)+string(l.ch))
	case '|':
		if l.peekChar() != '|' {
			tok = newToken(token.ILLEGAL, l.ch)
			return tok
		}
		ch := l.ch
		l.readChar()
		tok = newTokenFromString(token.OR, string(ch)+string(l.ch))
	case '-':
		if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			tok = newTokenFromString(token.DECRE, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.MINUS, l.ch)
		}
	case '+':
		if l.peekChar() == '+' {
			ch := l.ch
			l.readChar()
			tok = newTokenFromString(token.INCRE, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.PLUS, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newTokenFromString(token.GTE, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.GT, l.ch)
		}

	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newTokenFromString(token.LTE, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.LT, l.ch)
		}

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newTokenFromString(token.NEQ, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readDigit()
			tok.Type = token.GuessTypeOfDigit(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) FormatLineCharacter() string {
	return fmt.Sprintf("[line: %d, character: %d]", l.lineNumber+1, l.characterPositionInLine)
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func newTokenFromString(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: ch}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readDigit read integer and floats
func (l *Lexer) readDigit() string {
	position := l.position
	for isDigit(l.ch) || (l.ch == '.' && isDigit(l.peekChar())) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipSingleLineComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}

	l.skipWhitespace()
}

func (l *Lexer) skipMultiLineComment() {

	for !(l.ch == '*' && l.peekChar() == '/') || l.ch == 0 {
		l.readChar()
	}
	l.readChar() // "*"
	l.readChar() // "/"

	l.skipWhitespace()
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return l.input[position:l.position]
}
