package lexer

import (
	"fmt"
	"github.com/gravataLonga/ninja/token"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
var true false if else import return
break for enum case delete () [] {} . ; : :: ,
!= == <= >= < > && || = & | ^ ~ << >> ?
+ - * ** / % 
// comment 
! 100 100.5 "hello" "\\"
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
		{token.IMPORT, "import"},
		{token.RETURN, "return"},
		{token.BREAK, "break"},
		{token.FOR, "for"},
		{token.ENUM, "enum"},
		{token.CASE, "case"},
		{token.DELETE, "delete"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.DOT, "."},
		{token.SEMICOLON, ";"},
		{token.COLON, ":"},
		{token.DOUBLE_COLON, "::"},
		{token.COMMA, ","},
		{token.NEQ, "!="},
		{token.EQ, "=="},
		{token.LTE, "<="},
		{token.GTE, ">="},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.AND, "&&"},
		{token.OR, "||"},
		{token.ASSIGN, "="},

		{token.BIT_AND, "&"},
		{token.BIT_OR, "|"},
		{token.BIT_XOR, "^"},
		{token.BIT_NOT, "~"},
		{token.SHIFT_LEFT, "<<"},
		{token.SHIFT_RIGHT, ">>"},

		{token.QUESTION_MARK, "?"},

		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.EXPONENCIAL, "**"},
		{token.SLASH, "/"},

		{token.MOD, "%"},
		{token.BANG, "!"},

		{token.INT, "100"},
		{token.FLOAT, "100.5"},
		{token.STRING, "hello"},
		{token.STRING, "\\"},
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
		t.Run(fmt.Sprintf("Test[%d]", i), func(t *testing.T) {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("testdata[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
			}
			if string(tok.Literal) != tt.expectedLiteral {
				t.Fatalf("testdata[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
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

func TestLexerReadString(t *testing.T) {
	input := `"\"foo\"";"\x00\x0a\x7f";"\r\n\t"`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STRING, "\"foo\""},
		{token.SEMICOLON, ";"},
		{token.STRING, "\x00\n\u007f"},
		{token.SEMICOLON, ";"},
		{token.STRING, "\r\n\t"},
	}
	lexer := New(strings.NewReader(input))

	for _, test := range tests {
		tok := lexer.NextToken()
		if tok.Type != test.expectedType {
			t.Fatalf("token type wrong. expected=%q, got=%q", test.expectedType, tok.Type)
		}

		if string(tok.Literal) != test.expectedLiteral {
			t.Fatalf("literal wrong. expected=%q, got=%q", test.expectedLiteral, tok.Literal)
		}
	}
}
