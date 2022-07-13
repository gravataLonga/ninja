package semantic

import (
	"fmt"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/parser"
	"strings"
	"testing"
)

func TestScopeVariable(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			`function () { var a = 1; }`,
		},
		{
			`function () { var a = true; }`,
		},
		{
			`function () { var a = true; var b = a; }`,
		},
		{
			`var a = 1`,
		},
		{
			`var a = true`,
		},
		{
			`var a = []`,
		},
		{
			`var a = {}`,
		},
		{
			`function (x) { return x; }`,
		},
		{
			`function () { function (y) { var a = y + 1; } }`,
		},
		{
			`function () { var a = -1; }`,
		},
		{
			`function () { var a = !true; }`,
		},
		{
			`var a = 1; function () { var b = a; }`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestScopeVariable[%d]", i), func(t *testing.T) {
			s := testSemantic(tt.input, t)

			if len(s.Errors()) > 0 {
				t.Fatalf("Semantic Analysis got errors. Got: %v", s.Errors())
			}
		})
	}
}

func TestScopeVariablesWrong(t *testing.T) {
	tests := []struct {
		input       string
		erroMessage interface{}
	}{
		{
			`function () { var a = a; }`,
			"Can't read local variable \"a\" in its own initializer IDENT at [Line: 1, Offset: 24]",
		},
		{
			`function () { var a = a + 1 }`,
			"Can't read local variable \"a\" in its own initializer IDENT at [Line: 1, Offset: 24]",
		},
		{
			`function () { if (true) {var b = b;} else {} }`,
			"Can't read local variable \"b\" in its own initializer IDENT at [Line: 1, Offset: 35]",
		},
		{
			`function () { var a = "local"; function () { var a = a; } }`,
			"Can't read local variable \"a\" in its own initializer IDENT at [Line: 1, Offset: 55]",
		},
		{
			`function () { var a = "local"; function () { function () { var a = a; } } }`,
			"Can't read local variable \"a\" in its own initializer IDENT at [Line: 1, Offset: 69]",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestScopeVariablesWrong[%d]", i), func(t *testing.T) {
			s := testSemantic(tt.input, t)

			if len(s.Errors()) <= 0 {
				t.Fatalf("Semantic Analyse didn't get any error")
			}

			if s.Errors()[0] != tt.erroMessage {
				t.Errorf("Semantic Analyse expected %s. Got: %s", tt.erroMessage, s.Errors()[0])
			}
		})
	}
}

func TestMarkVariableAsGlobal(t *testing.T) {
	s := testSemantic("var a = 0", t)

	if len(s.globalVariable) <= 0 {
		t.Fatalf("Variable a isn't mark as global")
	}

}

// checkParserErrors check if there are parser errors
func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

// checkParserErrors check if there are parser errors
func checkSemanticErrors(t *testing.T, s *Semantic) {
	errors := s.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("semantic analyzer has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("semantic analyzer : %q", msg)
	}
	t.FailNow()
}

// testSemantic execute input code and check if there are parser error
// and return result object.Object
func testSemantic(input string, t *testing.T) *Semantic {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	s := New()
	s.Analysis(p.ParseProgram())

	checkParserErrors(t, p)
	return s
}
