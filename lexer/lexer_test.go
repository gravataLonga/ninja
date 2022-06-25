package lexer

import (
	"ninja/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
@
|=
&=
!true
!=
>
<
>=
<=
++5;
5--;
var five = 5;
var conta_0_L3Tr4as = "Other Variable";
var pi = 3.1415;
var booleano = true;
var booleanoFalso = false;
var conta = five * pi;
var contaMetada = (conta / 2) - 10;

function saySomething(name) {
    return "Said: " + name;
}

// Single Comment
saySomething("Hello " + name);

/*
Multiple Comment
*/
if (five > 10) {
    saySomething("Ups");
} elseif (five == 10) {
    saySomething("Its same");
} else {
    saySomething("All fine");
}

var statusCode = [200, 300, 400];

if (statusCode[0] != 200) {
    saySomething("200 status code");
}

var hashesResponse = {"nope":"an error happend", "ok":"everthing is ok"}

for (var i = 0; i <= len(statusCode); var i = i + 1) {
	puts(statusCode[i])
	break;
}

5 && 10
10 || 20

import "testing.mo"

delete x[0];
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ILLEGAL, "@"},
		{token.ILLEGAL, "|="},
		{token.ILLEGAL, "&="},
		{token.BANG, "!"},
		{token.TRUE, "true"},
		{token.NEQ, "!="},
		{token.GT, ">"},
		{token.LT, "<"},
		{token.GTE, ">="},
		{token.LTE, "<="},
		{token.INCRE, "++"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.DECRE, "--"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "conta_0_L3Tr4as"},
		{token.ASSIGN, "="},
		{token.STRING, "Other Variable"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "pi"},
		{token.ASSIGN, "="},
		{token.FLOAT, "3.1415"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "booleano"},
		{token.ASSIGN, "="},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "booleanoFalso"},
		{token.ASSIGN, "="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "conta"},
		{token.ASSIGN, "="},
		{token.IDENT, "five"},
		{token.ASTERISK, "*"},
		{token.IDENT, "pi"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "contaMetada"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.IDENT, "conta"},
		{token.SLASH, "/"},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.MINUS, "-"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "function"},
		{token.IDENT, "saySomething"},
		{token.LPAREN, "("},
		{token.IDENT, "name"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.STRING, "Said: "},
		{token.PLUS, "+"},
		{token.IDENT, "name"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IDENT, "saySomething"},
		{token.LPAREN, "("},
		{token.STRING, "Hello "},
		{token.PLUS, "+"},
		{token.IDENT, "name"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.GT, ">"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "saySomething"},
		{token.LPAREN, "("},
		{token.STRING, "Ups"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSEIF, "elseif"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "saySomething"},
		{token.LPAREN, "("},
		{token.STRING, "Its same"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.IDENT, "saySomething"},
		{token.LPAREN, "("},
		{token.STRING, "All fine"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.VAR, "var"},
		{token.IDENT, "statusCode"},
		{token.ASSIGN, "="},
		{token.LBRACKET, "["},
		{token.INT, "200"},
		{token.COMMA, ","},
		{token.INT, "300"},
		{token.COMMA, ","},
		{token.INT, "400"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "statusCode"},
		{token.LBRACKET, "["},
		{token.INT, "0"},
		{token.RBRACKET, "]"},
		{token.NEQ, "!="},
		{token.INT, "200"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "saySomething"},
		{token.LPAREN, "("},
		{token.STRING, "200 status code"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.VAR, "var"},
		{token.IDENT, "hashesResponse"},
		{token.ASSIGN, "="},
		{token.LBRACE, "{"},
		{token.STRING, "nope"},
		{token.COLON, ":"},
		{token.STRING, "an error happend"},
		{token.COMMA, ","},
		{token.STRING, "ok"},
		{token.COLON, ":"},
		{token.STRING, "everthing is ok"},
		{token.RBRACE, "}"},
		{token.FOR, "for"},
		{token.LPAREN, "("},
		{token.VAR, "var"},
		{token.IDENT, "i"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.LTE, "<="},
		{token.IDENT, "len"},
		{token.LPAREN, "("},
		{token.IDENT, "statusCode"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "i"},
		{token.ASSIGN, "="},
		{token.IDENT, "i"},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "puts"},
		{token.LPAREN, "("},
		{token.IDENT, "statusCode"},
		{token.LBRACKET, "["},
		{token.IDENT, "i"},
		{token.RBRACKET, "]"},
		{token.RPAREN, ")"},
		{token.BREAK, "break"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "5"},
		{token.AND, "&&"},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.OR, "||"},
		{token.INT, "20"},
		{token.IMPORT, "import"},
		{token.STRING, "testing.mo"},
		{token.DELETE, "delete"},
		{token.IDENT, "x"},
		{token.LBRACKET, "["},
		{token.INT, "0"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},

		{token.EOF, "\x00"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if string(tok.Literal) != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func BenchmarkLexer_NextToken(b *testing.B) {
	input := `var a = 0; "ola" != true; if (a > 0) { return 1; }; import "testing";`
	for i := 0; i < b.N; i++ {
		l := New(input)

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
		l := New(input)

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

	l := New(input)

	l.readChar()

	if l.ch != 0 {
		t.Fatalf("lexer.readChar expected 0. Got: %v", l.ch)
	}
}

func TestLexer_KeepTrackPosition(t *testing.T) {
	// not we are looking by identifier "a"
	tests := []struct {
		input        string
		linePosition int
		charPosition int
		expected     string
	}{
		{
			`var a = 0;`,
			0,
			4,
			"[line: 1, character: 4]",
		},
		{
			`
var b = 0;
var a = 0;`,
			2,
			4,
			"[line: 3, character: 4]",
		}, {
			`
var b = 0;
/*
Hello World
*/
var a = 0;`,
			5,
			4,
			"[line: 6, character: 4]",
		},
	}

	for _, tt := range tests {
		l := New(tt.input)
		for {
			curToken := l.NextToken()
			if l.peekChar() == 'a' {
				break
			}

			if curToken.Type == token.EOF {
				t.Fatalf("Unable to find token a. Got: EOF")
				break
			}
		}

		if tt.linePosition != l.lineNumber {
			t.Errorf("Wrong line, expected %d. Got: %d", tt.linePosition, l.lineNumber)
		}

		if tt.charPosition != l.characterPositionInLine {
			t.Errorf("Wrong line, expected %d. Got: %d", tt.charPosition, l.characterPositionInLine)
		}

		if tt.expected != l.FormatLineCharacter() {
			t.Errorf("l.formatLineCharacter Expected %s. Got: %s", tt.expected, l.FormatLineCharacter())
		}

	}
}
