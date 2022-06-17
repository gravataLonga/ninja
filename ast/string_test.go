package ast

import (
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			createVarStatement("myVar", createIdentifier("anotherVar")),
		},
	}

	if program.String() != "var myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
