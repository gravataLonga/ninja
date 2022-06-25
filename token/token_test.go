package token

import (
	"testing"
)

func TestLookupIdentifier(t *testing.T) {
	tests := []struct {
		char          []byte
		expectedToken TokenType
	}{
		{[]byte("var"), VAR},
		{[]byte("true"), TRUE},
		{[]byte("false"), FALSE},
		{[]byte("function"), FUNCTION},
		{[]byte("return"), RETURN},
		{[]byte("if"), IF},
		{[]byte("elseif"), ELSEIF},
		{[]byte("else"), ELSE},
		{[]byte("for"), FOR},
		{[]byte("import"), IMPORT},
		{[]byte("delete"), DELETE},
		{[]byte("break"), BREAK},
		{[]byte("testing_var"), IDENT},
	}

	for _, tt := range tests {
		tok := LookupIdentifier(tt.char)
		if tok != tt.expectedToken {
			t.Errorf("Expected %s. Got: %s", tt.expectedToken, tok)
		}
	}
}

func TestTokenType_String(t *testing.T) {
	toks := getAllTokens()

	for _, tt := range toks {
		if tt.String() == "" {
			t.Fatalf("TokenType %d don't have string format. %s", tt, tt)
		}
	}
}

func TestDigitType(t *testing.T) {
	tests := []struct {
		digit    []byte
		expected TokenType
	}{
		{
			[]byte("1"),
			INT,
		},
		{
			[]byte("100000"),
			INT,
		},
		{
			[]byte("1.44"),
			FLOAT,
		},
		{
			[]byte("0.343"),
			FLOAT,
		},
		{
			[]byte(".1232"),
			FLOAT,
		},
	}

	for _, tt := range tests {
		v := DigitType(tt.digit)
		if v != tt.expected {
			t.Errorf("Expected TokenType %s. Got: %s", tt.expected, v)
		}
	}
}

func getAllTokens() []TokenType {
	ts := make([]TokenType, ENDTOKEN)
	for i := 0; i < int(ENDTOKEN); i++ {
		ts[i] = TokenType(i)
	}
	return ts
}
