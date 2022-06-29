package parser

import (
	"ninja/ast"
	"ninja/lexer"
	"strings"
	"testing"
)

func TestDeleteStatementArrayHashLiteral(t *testing.T) {
	input := "delete myArray[1]"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.DeleteStatement)
	if !ok {
		t.Fatalf("exp not *ast.DeleteStatement. got=%T", program.Statements[0])
	}

	index, ok := stmt.Index.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.DeleteStatement.Index is not IntegerLiteral. got=%T", stmt.Index)
	}

	left, ok := stmt.Left.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.DeleteStatement.Identifier is not IdentifierLiteral. got=%T", stmt.Left)
	}

	if !testIdentifier(t, left, "myArray") {
		return
	}

	if !testIntegerLiteral(t, index, 1) {
		return
	}
}
