package parser

import (
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 % 5;", 5, "%", 5},
		{"5 & 5;", 5, "&", 5},
		{"5 | 5;", 5, "|", 5},
		{"5 ^ 5;", 5, "^", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 >> 5;", 5, ">>", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 << 5;", 5, "<<", 5},
		{"5 >= 5;", 5, ">=", 5},
		{"5 <= 5;", 5, "<=", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"5 && 5;", 5, "&&", 5},
		{"5 || 5;", 5, "||", 5},
		{"foobar + barfoo", "foobar", "+", "barfoo"},
		{"foobar - barfoo", "foobar", "-", "barfoo"},
		{"foobar * barfoo", "foobar", "*", "barfoo"},
		{"foobar / barfoo", "foobar", "/", "barfoo"},
		{"foobar % barfoo", "foobar", "%", "barfoo"},
		{"foobar & barfoo", "foobar", "&", "barfoo"},
		{"foobar | barfoo", "foobar", "|", "barfoo"},
		{"foobar ^ barfoo", "foobar", "^", "barfoo"},
		{"foobar > barfoo", "foobar", ">", "barfoo"},
		{"foobar >> barfoo", "foobar", ">>", "barfoo"},
		{"foobar < barfoo", "foobar", "<", "barfoo"},
		{"foobar << barfoo", "foobar", "<<", "barfoo"},
		{"foobar >= barfoo", "foobar", ">=", "barfoo"},
		{"foobar <= barfoo", "foobar", "<=", "barfoo"},
		{"foobar == barfoo", "foobar", "==", "barfoo"},
		{"foobar != barfoo", "foobar", "!=", "barfoo"},
		{"foobar && barfoo", "foobar", "&&", "barfoo"},
		{"foobar || barfoo", "foobar", "||", "barfoo"},
		{"true == false", true, "==", false},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
		{"false || false", false, "||", false},
		{"false && false", false, "&&", false},
	}

	for i, tt := range infixTests {
		t.Run(fmt.Sprintf("TestParsingInfixExpressions[%d]", i), func(t *testing.T) {
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

			if !testInfixExpression(t, stmt.Expression, tt.leftValue,
				tt.operator, tt.rightValue) {
				return
			}
		})

	}
}
