package ast

import (
	"ninja/token"
	"testing"
)

func TestBlockStatement_String(t *testing.T) {
	stmts := []Statement{
		&ReturnStatement{Token: token.Token{Type: token.RETURN, Literal: "return"}, ReturnValue: &IntegerLiteral{
			Token: token.Token{Type: token.INT, Literal: "1"}, Value: 1,
		}},
	}
	block := &BlockStatement{Token: token.Token{Type: token.LBRACE, Literal: "{"}, Statements: stmts}

	if block.String() != "return 1;" {
		t.Errorf("ReturnStatment isnt equal to %s. Got: %s", "return 1;", block.String())
	}

	if block.TokenLiteral() != "{" {
		t.Errorf("ReturnStatment.TokenLiteral() isnt equal to %s. Got: %s", "{", block.TokenLiteral())
	}

}
