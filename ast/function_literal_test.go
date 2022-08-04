package ast

import (
	"github.com/gravataLonga/ninja/token"
	"strconv"
	"testing"
)

func TestFunctionLiteral_String(t *testing.T) {
	tests := []struct {
		parameters []int64
		body       int64
		expected   string
	}{
		{
			[]int64{}, 1, "function() {return 1;}",
		},
		{
			[]int64{1}, 2, "function(1) {return 2;}",
		},
		{
			[]int64{1, 2}, 4, "function(1, 2) {return 4;}",
		},
	}

	for _, tt := range tests {

		stmts := []Statement{
			createReturnStatement(createIntegerLiteral(tt.body)),
		}

		blockStatement := &BlockStatement{Token: token.Token{Type: token.LBRACE, Literal: "{"}, Statements: stmts}
		var argumentsIdentifier []Expression
		for _, arg := range tt.parameters {
			integerLiteral := &Identifier{Token: token.Token{Type: token.INT, Literal: strconv.FormatInt(arg, 10)}, Value: strconv.FormatInt(arg, 10)}
			argumentsIdentifier = append(argumentsIdentifier, integerLiteral)
		}
		fn := &FunctionLiteral{
			Token:      token.Token{Type: token.FUNCTION, Literal: "function"},
			Parameters: argumentsIdentifier,
			Body:       blockStatement,
		}

		if fn.String() != tt.expected {
			t.Fatalf("Function.String() isn't expected %s. Got %s", tt.expected, fn.String())
		}
	}
}

func TestFunctionLiteralWithIdentifier_String(t *testing.T) {
	tests := []struct {
		parameters []int64
		body       int64
		expected   string
	}{
		{
			[]int64{}, 1, "function add() {return 1;}",
		},
		{
			[]int64{1}, 2, "function add(1) {return 2;}",
		},
		{
			[]int64{1, 2}, 4, "function add(1, 2) {return 4;}",
		},
	}

	for _, tt := range tests {

		stmts := []Statement{
			createReturnStatement(createIntegerLiteral(tt.body)),
		}

		blockStatement := &BlockStatement{Token: token.Token{Type: token.LBRACE, Literal: "{"}, Statements: stmts}
		var argumentsIdentifier []Expression
		for _, arg := range tt.parameters {
			integerLiteral := &Identifier{Token: token.Token{Type: token.INT, Literal: strconv.FormatInt(arg, 10)}, Value: strconv.FormatInt(arg, 10)}
			argumentsIdentifier = append(argumentsIdentifier, integerLiteral)
		}
		fn := &FunctionLiteral{
			Token:      token.Token{Type: token.FUNCTION, Literal: "function"},
			Parameters: argumentsIdentifier,
			Body:       blockStatement,
			Name:       createIdentifier("add"),
		}

		if fn.String() != tt.expected {
			t.Fatalf("Function.String() isn't expected %s. Got %s", tt.expected, fn.String())
		}
	}
}
