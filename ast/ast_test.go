package ast

import (
	"ninja/token"
	"strconv"
)

func createIdentifier(name string) *Identifier {
	return &Identifier{
		Token: token.Token{Type: token.IDENT, Literal: []byte(name)},
		Value: name,
	}
}

func createVarStatement(name string, value Expression) *VarStatement {
	return &VarStatement{
		Token: token.Token{Type: token.VAR, Literal: []byte("var")},
		Name:  createIdentifier(name),
		Value: value,
	}
}

func createIntegerLiteral(value int64) Expression {
	return &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: []byte(strconv.FormatInt(value, 10))}, Value: value}
}

func createFloatLiteral(value float64) Expression {
	return &FloatLiteral{Token: token.Token{Type: token.FLOAT, Literal: []byte(strconv.FormatFloat(value, 'f', -1, 64))}, Value: value}
}

func createBoolean(value bool) Expression {
	literal := "true"
	var tokenName token.TokenType = token.TRUE
	if !value {
		literal = "false"
		tokenName = token.FALSE
	}

	return &Boolean{Token: token.Token{Type: tokenName, Literal: []byte(literal)}, Value: value}
}

func createReturnStatement(returned Expression) Statement {
	return &ReturnStatement{
		Token:       token.Token{Type: token.RETURN, Literal: []byte("return")},
		ReturnValue: returned,
	}
}

func createBlockStatements(stmts []Statement) Statement {
	return &BlockStatement{Token: token.Token{Type: token.LBRACE, Literal: []byte("{")}, Statements: stmts}
}
