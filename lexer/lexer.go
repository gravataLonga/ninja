package lexer

import (
	"io"
	"ninja/token"
	"unicode/utf8"
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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':

		tok = l.newTokenPeekOrDefault(token.ASSIGN, map[byte]token.TokenType{
			'=': token.EQ,
		})

	case ';':
		tok = l.newToken(token.SEMICOLON, []byte{l.ch})
	case '"':
		tok.Type = token.STRING
		tok.Literal = runesToUTF8(l.readString())
	case '*':
		tok = l.newToken(token.ASTERISK, []byte{l.ch})
	case '/':
		if l.peekChar() == '/' {
			l.skipSingleLineComment()
			return l.NextToken()
		}

		if l.peekChar() == '*' {
			l.skipMultiLineComment()
			return l.NextToken()
		}
		tok = l.newToken(token.SLASH, []byte{l.ch})
	case '&':
		if l.peekChar() != '&' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.ILLEGAL, []byte{ch, l.ch})
		} else {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.AND, []byte{ch, l.ch})
		}

	case '|':
		if l.peekChar() != '|' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.ILLEGAL, []byte{ch, l.ch})
		} else {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.OR, []byte{ch, l.ch})
		}

	case '-':

		tok = l.newTokenPeekOrDefault(token.MINUS, map[byte]token.TokenType{
			'-': token.DECRE,
		})

	case '+':

		tok = l.newTokenPeekOrDefault(token.PLUS, map[byte]token.TokenType{
			'+': token.INCRE,
		})
	case '>':

		tok = l.newTokenPeekOrDefault(token.GT, map[byte]token.TokenType{
			'=': token.GTE,
		})

	case '<':

		tok = l.newTokenPeekOrDefault(token.LT, map[byte]token.TokenType{
			'=': token.LTE,
		})

	case '!':

		tok = l.newTokenPeekOrDefault(token.BANG, map[byte]token.TokenType{
			'=': token.NEQ,
		})

	case ':':
		tok = l.newToken(token.COLON, []byte{l.ch})
	case '(':
		tok = l.newToken(token.LPAREN, []byte{l.ch})
	case ')':
		tok = l.newToken(token.RPAREN, []byte{l.ch})
	case '{':
		tok = l.newToken(token.LBRACE, []byte{l.ch})
	case '}':
		tok = l.newToken(token.RBRACE, []byte{l.ch})
	case '[':
		tok = l.newToken(token.LBRACKET, []byte{l.ch})
	case ']':
		tok = l.newToken(token.RBRACKET, []byte{l.ch})
	case ',':
		tok = l.newToken(token.COMMA, []byte{l.ch})
	case '.':
		tok = l.newToken(token.DOT, []byte{l.ch})
	case 0:
		tok = l.newToken(token.EOF, []byte{0})
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			return l.newToken(token.LookupIdentifier(literal), literal)
		} else if isDigit(l.ch) {
			literal := l.readDigit()
			return l.newToken(token.DigitType(literal), literal)
		} else {
			tok = l.newToken(token.ILLEGAL, []byte{l.ch})
		}
	}

	l.readChar()
	return tok
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

// keepTrackLineAndCharPosition is a method which keep tracking where position of
// pointer of lexer is point at.
func (l *Lexer) keepTrackLineAndCharPosition() {
	if l.ch == '\n' {
		l.lineNumber += 1
		l.characterPositionInLine = 0
	} else {
		l.characterPositionInLine += 1
	}
}

func (l *Lexer) newToken(tokenType token.TokenType, ch []byte) token.Token {
	location := token.Location{Line: l.lineNumber + 1, Offset: l.characterPositionInLine}
	return token.Token{Type: tokenType, Literal: ch, Location: location}
}

func (l *Lexer) newTokenPeekOrDefault(tokenType token.TokenType, expectedPeek map[byte]token.TokenType) token.Token {
	peekToken, ok := expectedPeek[l.peekChar()]
	if !ok {
		return l.newToken(tokenType, []byte{l.ch})
	}

	ch := l.ch
	l.readChar()
	return l.newToken(peekToken, []byte{ch, l.ch})
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

func (l *Lexer) readString() []rune {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return []rune(l.input[position:l.position])
}

func runesToUTF8(rs []rune) []byte {
	size := 0
	for _, r := range rs {
		size += utf8.RuneLen(r)
	}

	bs := make([]byte, size)

	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs
}
