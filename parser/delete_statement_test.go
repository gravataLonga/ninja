package parser

import (
	"ninja/ast"
	"ninja/lexer"
	"testing"
)

func TestParsingDeleteIndexExpressions(t *testing.T) {
	input := "delete myArray[1]"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.DeleteStatement)
	if !ok {
		t.Fatalf("exp not *ast.DeleteStatement. got=%T", program.Statements[0])
	}

	index, ok := stmt.Identifier.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("exp not *ast.DeleteStatement.Identifier is not IndexExpression. got=%T", index)
	}

	if !testIdentifier(t, index.Left, "myArray") {
		return
	}

	if !testIntegerLiteral(t, index.Index, 1) {
		return
	}
}
