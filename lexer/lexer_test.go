package lexer

import (
	"fmt"
	"ninja/token"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
var true false if else
break for () [] {} . ; : 
!= == <= >= < > && || = 
+ - * / 
// comment 
! 100 100.5 "hello" 
++5 --5 5++ 5-- count 
/* 
multiple comment 
*/ 
function delete @ 
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},

		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.BREAK, "break"},
		{token.FOR, "for"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.DOT, "."},
		{token.SEMICOLON, ";"},
		{token.COLON, ":"},
		{token.NEQ, "!="},
		{token.EQ, "=="},
		{token.LTE, "<="},
		{token.GTE, ">="},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.AND, "&&"},
		{token.OR, "||"},
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.BANG, "!"},

		{token.INT, "100"},
		{token.FLOAT, "100.5"},
		{token.STRING, "hello"},
		{token.INCRE, "++"},
		{token.INT, "5"},
		{token.DECRE, "--"},
		{token.INT, "5"},
		{token.INT, "5"},
		{token.INCRE, "++"},
		{token.INT, "5"},
		{token.DECRE, "--"},
		{token.IDENT, "count"},
		{token.FUNCTION, "function"},
		{token.DELETE, "delete"},
		{token.ILLEGAL, "@"},

		{token.EOF, "\x00"},
	}

	l := New(strings.NewReader(input))

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%s", tt.expectedType.String()), func(t *testing.T) {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
			}
			if string(tok.Literal) != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
			}
		})

	}
}

func BenchmarkLexer_NextToken(b *testing.B) {
	input := `var a = 0; "ola" != true; if (a > 0) { return 1; }; import "testing";`
	for i := 0; i < b.N; i++ {
		l := New(strings.NewReader(input))

		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break
			}
		}
	}
}

func BenchmarkLexer_NextTokenMedium(b *testing.B) {
	input := `
var a = 0; 
"ola" != true; 
if (a > 0) { return 1; }; 
import "testing";

/*
 With comments
*/
for (;;) {
	return 1+1*33 <= 45443.343;
}

function() {
	import "I";
}();
`
	for i := 0; i < b.N; i++ {
		l := New(strings.NewReader(input))

		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break
			}
		}
	}
}

func TestLexer_PastEOF(t *testing.T) {
	input := ``

	l := New(strings.NewReader(input))

	l.readChar()

	if l.ch != 0 {
		t.Fatalf("lexer.readChar expected 0. Got: %v", l.ch)
	}
}

func TestStringAcceptUtf8Character(t *testing.T) {
	input := `"नमस्ते दुनिया or Hello World or Olá Mundo"`
	l := New(strings.NewReader(input))

	tok := l.NextToken()

	if tok.Type != token.STRING {
		t.Fatalf("Expected string token, got: %s", tok.Type)
	}

	if string(tok.Literal) != `नमस्ते दुनिया or Hello World or Olá Mundo` {
		t.Fatalf("Expected string to be àãç, got: %s", tok.Literal)
	}

}
