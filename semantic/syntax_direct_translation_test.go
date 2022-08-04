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

/*
func TestDeclareIdentifierRegisterHops(t *testing.T) {
	input := `var a = 1`
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	program := p.ParseProgram()

	semantic := New(program)
	nodes := semantic.Analysis()

	astProgram, _ := nodes.(*ast.Program)
	astVarExpressionStatement, _ := astProgram.Statements[0].(*ast.VarStatement)
	astIdentifier := astVarExpressionStatement.Name

	if astIdentifier.Value != "a" {
		t.Fatalf("Identifier isn't a")
	}

	if astIdentifier.Stack == nil {
		t.Fatalf("Identifier didn't declare stack")
	}

	if astIdentifier.Stack.Size() != 1 {
		t.Errorf("Stack for identifier isn't equal 1")
	}
}
*/
