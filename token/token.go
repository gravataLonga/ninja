package token

import (
	"bytes"
)

type TokenType int8

type Token struct {
	Type    TokenType
	Literal []byte // be a string don't have same performance as using int or byte
}

func (t TokenType) String() string {
	return []string{
		"ILLEGAL",
		"EOF",
		"IDENT",
		"INT",
		"FLOAT",
		"STRING",
		"=",
		"+",
		"-",
		"!",
		"*",
		"/",
		"++",
		"--",
		"<",
		">",
		"<=",
		">=",
		"==",
		"!=",
		"&&",
		"||",
		",",
		";",
		":",
		"(",
		")",
		"{",
		"}",
		"[",
		"]",
		"FUNCTION",
		"FUNCTION_LITERAL",
		"VARIABLE",
		"TRUE",
		"FALSE",
		"IF",
		"ELSEIF",
		"ELSE",
		"RETURN",
		"IMPORT",
		"FOR",
		"DELETE",
	}[t]
}

const (
	ILLEGAL TokenType = iota //  "ILLEGAL"
	EOF                      // "EOF"

	IDENT  // "IDENT"
	INT    // "INT"
	FLOAT  // "FLOAT"
	STRING // "STRING"

	ASSIGN   // "="
	PLUS     // "+"
	MINUS    // "-"
	BANG     // "!"
	ASTERISK // "*"
	SLASH    // "/"
	INCRE    // "++"
	DECRE    // "--"

	LT  // "<"
	GT  // ">"
	LTE // "<="
	GTE // ">="
	EQ  // "=="
	NEQ // "!="
	AND // "&&"
	OR  // "||"

	COMMA     // ","
	SEMICOLON // ";"
	COLON     // ":"
	LPAREN    // "("
	RPAREN    // ")"
	LBRACE    // "{"
	RBRACE    // "}"
	LBRACKET  // "["
	RBRACKET  // "]"

	FUNCTION         // "FUNCTION"
	FUNCTION_LITERAL // "FUNCTION_LITERAL"
	VAR              // "VARIABLE"
	TRUE             // "TRUE"
	FALSE            // "FALSE"
	IF               // "IF"
	ELSEIF           // "ELSEIF"
	ELSE             // "ELSE"
	RETURN           // "RETURN"
	IMPORT           // "IMPORT"
	FOR              // "FOR"
	DELETE           // "DELETE"
)

var keywords = map[string]TokenType{
	"var":      VAR,
	"true":     TRUE,
	"false":    FALSE,
	"function": FUNCTION,
	"return":   RETURN,
	"if":       IF,
	"elseif":   ELSEIF,
	"else":     ELSE,
	"for":      FOR,
	"import":   IMPORT,
	"delete":   DELETE,
}

func LookupIdentifier(ident []byte) TokenType {
	if tok, ok := keywords[string(ident)]; ok {
		return tok
	}
	return IDENT
}

func DigitType(digit []byte) TokenType {
	hasDot := bytes.IndexByte(digit, '.')
	if hasDot >= 0 {
		return FLOAT
	}
	return INT
}
