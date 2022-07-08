package semantic

import (
	"fmt"
	"ninja/lexer"
	"ninja/parser"
	"strings"
	"testing"
)

func TestResolveVar(t *testing.T) {
	tests := []struct {
		input       string
		erroMessage interface{}
	}{
		{
			`var a = a;`,
			"Can't read local variable a in its own initializer",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestResolveVar[%d]", i), func(t *testing.T) {
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

// testEval execute input code and check if there are parser error
// and return result object.Object
func testSemantic(input string, t *testing.T) *Semantic {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	s := New(p)
	s.Analyze()

	checkParserErrors(t, p)
	return s
}
