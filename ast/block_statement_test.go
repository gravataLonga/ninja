package ast

import (
	"testing"
)

func TestBlockStatement_String(t *testing.T) {

	tests := []struct {
		stmts    []Statement
		expected string
	}{
		{
			[]Statement{
				createReturnStatement(createIntegerLiteral(1)),
			},
			"return 1;",
		},
		{
			[]Statement{
				createVarStatement("me", createIntegerLiteral(1)),
				createReturnStatement(createIdentifier("me")),
			},
			"var me = 1;return me;",
		},
	}

	for _, tt := range tests {
		block := createBlockStatements(tt.stmts)

		if block.String() != tt.expected {
			t.Errorf("ReturnStatment isnt equal to %s. Got: %s", tt.expected, block.String())
		}

		if block.TokenLiteral() != "{" {
			t.Errorf("ReturnStatment.TokenLiteral() isnt equal to %s. Got: %s", "{", block.TokenLiteral())
		}
	}

}
