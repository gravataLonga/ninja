package semantic

import (
	"github.com/gravataLonga/ninja/ast"
	"testing"
)

func TestNew(t *testing.T) {
	program := &ast.Program{Statements: []ast.Statement{}}
	s := New(program)
	a := s.Analysis()

	if program.String() != a.String() {
		t.Fatalf("Analysis din't give same program string")
	}
}
