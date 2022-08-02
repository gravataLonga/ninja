package ast

import (
	"github.com/gravataLonga/ninja/token"
	"strconv"
	"testing"
)

func TestFunction_String(t *testing.T) {
	tests := []struct {
		name       string
		parameters []int64
		body       int64
		expected   string
	}{
		{
			"add", []int64{}, 1, "function add() return 1;",
		},
		{
			"other", []int64{1}, 2, "function other(1) return 2;",
		},
		{
			"multiple", []int64{1, 2}, 4, "function multiple(1, 2) return 4;",
		},
	}

	for _, tt := range tests {

		stmts := []Statement{
			&ReturnStatement{Token: token.Token{Type: token.RETURN, Literal: "return"}, ReturnValue: &IntegerLiteral{
				Token: token.Token{Type: token.INT, Literal: strconv.FormatInt(tt.body, 10)}, Value: tt.body,
			}},
		}
		blockStatement := &BlockStatement{Token: token.Token{Type: token.LBRACE, Literal: "{"}, Statements: stmts}
		nameIdentifier := &Identifier{Token: token.Token{Type: token.IDENT, Literal: "var"}, Value: tt.name}
		argumentsIdentifier := []Expression{}
		for _, arg := range tt.parameters {
			integerLiteral := &Identifier{Token: token.Token{Type: token.INT, Literal: strconv.FormatInt(arg, 10)}, Value: strconv.FormatInt(arg, 10)}
			argumentsIdentifier = append(argumentsIdentifier, integerLiteral)
		}
		fn := &Function{
			Token:      token.Token{Type: token.FUNCTION, Literal: "function"},
			Name:       nameIdentifier,
			Parameters: argumentsIdentifier,
			Body:       blockStatement,
		}

		if fn.String() != tt.expected {
			t.Fatalf("Function.String() isn't expected %s. Got %s", tt.expected, fn.String())
		}
	}
}
