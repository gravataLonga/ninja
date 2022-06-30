package lexer

import (
	"fmt"
	"io"
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

func New(in io.Reader) *Lexer {
	r, err := io.ReadAll(in)
	if err != nil {
		return nil
	}
	l := &Lexer{input: string(r), lineNumber: 0}
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
			tok = newToken(token.EQ, []byte{ch, l.ch})
		} else {
			tok = newToken(token.ASSIGN, []byte{l.ch})
		}

	case ';':
		tok = newToken(token.SEMICOLON, []byte{l.ch})
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '*':
		tok = newToken(token.ASTERISK, []byte{l.ch})
	case '/':
		if l.peekChar() == '/' {
			l.skipSingleLineComment()
			return l.NextToken()
		}

		if l.peekChar() == '*' {
			l.skipMultiLineComment()
			return l.NextToken()
		}
		tok = newToken(token.SLASH, []byte{l.ch})
	case '&':
		if l.peekChar() != '&' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.ILLEGAL, []byte{ch, l.ch})
		} else {
			ch := l.ch
			l.readChar()
			tok = newToken(token.AND, []byte{ch, l.ch})
		}

	case '|':
		if l.peekChar() != '|' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.ILLEGAL, []byte{ch, l.ch})
		} else {
			ch := l.ch
			l.readChar()
			tok = newToken(token.OR, []byte{ch, l.ch})
		}

	case '-':
		if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.DECRE, []byte{ch, l.ch})
		} else {
			tok = newToken(token.MINUS, []byte{l.ch})
		}
	case '+':
		if l.peekChar() == '+' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.INCRE, []byte{ch, l.ch})
		} else {
			tok = newToken(token.PLUS, []byte{l.ch})
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.GTE, []byte{ch, l.ch})
		} else {
			tok = newToken(token.GT, []byte{l.ch})
		}

	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.LTE, []byte{ch, l.ch})
		} else {
			tok = newToken(token.LT, []byte{l.ch})
		}

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.NEQ, []byte{ch, l.ch})
		} else {
			tok = newToken(token.BANG, []byte{l.ch})
		}
	case ':':
		tok = newToken(token.COLON, []byte{l.ch})
	case '(':
		tok = newToken(token.LPAREN, []byte{l.ch})
	case ')':
		tok = newToken(token.RPAREN, []byte{l.ch})
	case '{':
		tok = newToken(token.LBRACE, []byte{l.ch})
	case '}':
		tok = newToken(token.RBRACE, []byte{l.ch})
	case '[':
		tok = newToken(token.LBRACKET, []byte{l.ch})
	case ']':
		tok = newToken(token.RBRACKET, []byte{l.ch})
	case ',':
		tok = newToken(token.COMMA, []byte{l.ch})
	case '.':
		tok = newToken(token.DOT, []byte{l.ch})
	case 0:
		tok.Literal = []byte{0}
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readDigit()
			tok.Type = token.DigitType(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, []byte{l.ch})
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) FormatLineCharacter() string {
	return fmt.Sprintf("[line: %d, character: %d]", l.lineNumber+1, l.characterPositionInLine)
}

func newToken(tokenType token.TokenType, ch []byte) token.Token {
	return token.Token{Type: tokenType, Literal: ch}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() []byte {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return []byte(l.input[position:l.position])
}

// readDigit read integer and floats
func (l *Lexer) readDigit() []byte {
	position := l.position
	for isDigit(l.ch) || (l.ch == '.' && isDigit(l.peekChar())) {
		l.readChar()
	}
	return []byte(l.input[position:l.position])
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

func (l *Lexer) readString() []byte {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return []byte(l.input[position:l.position])
}
