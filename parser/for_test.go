package parser

import (
	"ninja/ast"
	"ninja/lexer"
	"testing"
)

func TestForStatement(t *testing.T) {
	tests := []struct {
		input               string
		initialStatement    string
		condition           string
		iterationExpression string
		block               string
	}{
		{input: "for (;;) {}", initialStatement: "", condition: "", iterationExpression: "", block: ""},
		{input: "for (var i = 0;;) {}", initialStatement: "var i = 0;", condition: "", iterationExpression: "", block: ""},
		{input: "for (var i = 0; i <= 3;) {}", initialStatement: "var i = 0;", condition: "(i <= 3)", iterationExpression: "", block: ""},
		{input: "for (var i = 0; i <= 3; var i = i + 1) {}", initialStatement: "var i = 0;", condition: "(i <= 3)", iterationExpression: "var i = (i + 1);", block: ""},
		{input: "for (var i = 0; i <= 3; var i = i + 1) { puts(i); }", initialStatement: "var i = 0;", condition: "(i <= 3)", iterationExpression: "var i = (i + 1);", block: "puts(i)"},
		{input: "for(;i <= 3;) {}", initialStatement: "", condition: "(i <= 3)", iterationExpression: "", block: ""},
		{input: "for(;;i = i + 1) {}", initialStatement: "", condition: "", iterationExpression: "i = (i + 1);", block: ""},
		{input: "for(;i <= 10 && a <= 10;) {}", initialStatement: "", condition: "((i <= 10) && (a <= 10))", iterationExpression: "", block: ""},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		fr := stmt.Expression.(*ast.ForStatement)

		if fr.InitialCondition != nil && fr.InitialCondition.String() != tt.initialStatement {
			t.Errorf("For.InitialCondition.String() isn't %s. Got: %s", tt.initialStatement, fr.InitialCondition.String())
		}

		if fr.Condition != nil && fr.Condition.String() != tt.condition {
			t.Errorf("For.Condition.String() isn't %s. Got: %s", tt.condition, fr.Condition.String())
		}

		if fr.Iteration != nil && fr.Iteration.String() != tt.iterationExpression {
			t.Errorf("For.Iteration.String() isn't %s. Got: %s", tt.iterationExpression, fr.Iteration.String())
		}

		if fr.Body != nil && fr.Body.String() != tt.block {
			t.Errorf("For.Body.String() isn't %s. Got: %s", tt.block, fr.Body.String())
		}
	}
}
