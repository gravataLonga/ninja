package semantic

import (
	"fmt"
	"ninja/lexer"
	"ninja/parser"
	"strings"
	"testing"
)

func TestScopeVariables(t *testing.T) {
	tests := []struct {
		input       string
		erroMessage interface{}
	}{
		{
			`var a = a;`,
			"Can't read local variable \"a\" in its own initializer IDENT at [Line: 1, Offset: 10]",
		},
		{
			`var a = [1, b];`,
			"Variable \"b\" not declare yet IDENT at [Line: 1, Offset: 14]",
		},
		{
			`var a = {"t": b}`,
			"Variable \"b\" not declare yet IDENT at [Line: 1, Offset: 16]",
		},
		{
			`var a = {b: 1}`,
			"Variable \"b\" not declare yet IDENT at [Line: 1, Offset: 11]",
		},
		{
			`if (a) {} else {}`,
			"Variable \"a\" not declare yet IDENT at [Line: 1, Offset: 6]",
		},
		{
			`if (true) {a} else {}`,
			"Variable \"a\" not declare yet IDENT at [Line: 1, Offset: 13]",
		},
		{
			`if (true) {} else {a}`,
			"Variable \"a\" not declare yet IDENT at [Line: 1, Offset: 21]",
		},
		{
			`if (true) {var b = b;} else {}`,
			"Can't read local variable \"b\" in its own initializer IDENT at [Line: 1, Offset: 21]",
		},
		{
			`var a = "local"; function () { var a = a; }`,
			"Can't read local variable \"a\" in its own initializer IDENT at [Line: 1, Offset: 6]",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestScopeVariables[%d]", i), func(t *testing.T) {
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
	s.Analyze(p.ParseProgram())

	checkParserErrors(t, p)
	return s
}
