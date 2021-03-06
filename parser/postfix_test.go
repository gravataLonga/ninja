package parser

import (
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestParsingPosfixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"i++", "++", "i"},
		{"i--", "--", "i"},
		{"1--", "--", 1},
		{"1++", "++", 1},
	}

	for i, tt := range prefixTests {
		t.Run(fmt.Sprintf("TestParsingPosfixExpressions[%d]", i), func(t *testing.T) {
			l := lexer.New(strings.NewReader(tt.input))
			p := New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if len(program.Statements) != 1 {
				t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
			}

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
			}

			exp, ok := stmt.Expression.(*ast.PostfixExpression)
			if !ok {
				t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
			}
			if exp.Operator != tt.operator {
				t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
			}

			testLiteralExpression(t, exp.Left, tt.value)
		})
	}
}
