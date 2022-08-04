package token

import (
	"fmt"
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
		{[]byte("enum"), ENUM},
		{[]byte("case"), CASE},
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
		digit    DigitType
		expected TokenType
	}{
		{
			DIGIT_TYPE_DECIMAL,
			INT,
		},
		{
			DIGIT_TYPE_FLOAT,
			FLOAT,
		},
		{
			DIGIT_TYPE_SCIENTIFIC_NOTATION,
			FLOAT,
		},
		{
			DIGIT_TYPE_HEXADECIMAL,
			INT,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestDigitType[%d]", i), func(t *testing.T) {
			if tt.digit.TokenType() != tt.expected {
				t.Errorf("Expecetd %v. Got: %v", tt.digit.TokenType(), tt.expected)
			}
		})

	}
}

func getAllTokens() []TokenType {
	ts := make([]TokenType, ENDTOKEN)
	for i := 0; i < int(ENDTOKEN); i++ {
		ts[i] = TokenType(i)
	}
	return ts
}
