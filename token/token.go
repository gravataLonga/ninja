package token

import (
	"fmt"
)

type TokenType int8
type DigitType int8

type Location struct {
	Line   int
	Offset int
}

type Token struct {
	Type    TokenType
	Literal string
	Location
}

// String() is used to transform TokenType int8 in it is string format, for better
// human debugging.
func (t TokenType) String() string {
	list := [...]string{
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
		"**",
		"%",
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
		"&",
		"|",
		"^",
		"~",
		"<<",
		">>",
		"?",
		"?:",
		".",
		",",
		";",
		":",
		"::",
		"(",
		")",
		"{",
		"}",
		"[",
		"]",
		"FUNCTION",
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
		"BREAK",
		"ENUM",
		"CASE",
	}

	if len(list)-1 < int(t) {
		return ""
	}

	return list[t]
}

const (
	DIGIT_TYPE_DECIMAL DigitType = iota
	DIGIT_TYPE_FLOAT
	DIGIT_TYPE_SCIENTIFIC_NOTATION
	DIGIT_TYPE_HEXADECIMAL
)

const (
	ILLEGAL TokenType = iota //  "ILLEGAL"
	EOF                      // "EOF"

	IDENT  // "IDENT"
	INT    // "INT"
	FLOAT  // "FLOAT"
	STRING // "STRING"

	ASSIGN      // "="
	PLUS        // "+"
	MINUS       // "-"
	BANG        // "!"
	ASTERISK    // "*"
	EXPONENCIAL // "**"
	MOD         // "%"
	SLASH       // "/"
	INCRE       // "++"
	DECRE       // "--"

	LT  // "<"
	GT  // ">"
	LTE // "<="
	GTE // ">="
	EQ  // "=="
	NEQ // "!="
	AND // "&&"
	OR  // "||"

	BIT_AND     // "&"
	BIT_OR      // "|"
	BIT_XOR     // "^"
	BIT_NOT     // "~"
	SHIFT_LEFT  // "<<"
	SHIFT_RIGHT // ">>"

	QUESTION_MARK  // "?"
	ELVIS_OPERATOR // "?:"

	DOT          // "."
	COMMA        // ","
	SEMICOLON    // ";"
	COLON        // ":"
	DOUBLE_COLON // "::"
	LPAREN       // "("
	RPAREN       // ")"
	LBRACE       // "{"
	RBRACE       // "}"
	LBRACKET     // "["
	RBRACKET     // "]"

	FUNCTION // "FUNCTION"
	VAR      // "VARIABLE"
	TRUE     // "TRUE"
	FALSE    // "FALSE"
	IF       // "IF"
	ELSEIF   // "ELSEIF"
	ELSE     // "ELSE"
	RETURN   // "RETURN"
	IMPORT   // "IMPORT"
	FOR      // "FOR"
	DELETE   // "DELETE"
	BREAK    // "BREAK"
	ENUM     // "ENUM"
	CASE     // "CASE"

	ENDTOKEN // Special token, only for testing purposes
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
	"break":    BREAK,
	"enum":     ENUM,
	"case":     CASE,
}

// LookupIdentifier it will search from []byte() it's keyword token
func LookupIdentifier(ident []byte) TokenType {
	if tok, ok := keywords[string(ident)]; ok {
		return tok
	}
	return IDENT
}

func (t Token) String() string {
	return fmt.Sprintf("%s at [Line: %d, Offset: %d]", t.Type, t.Line, t.Offset)
}

func (d DigitType) IsEqual(digitType DigitType) bool {
	return d == digitType
}

func (d DigitType) TokenType() TokenType {
	switch d {
	case DIGIT_TYPE_FLOAT:
		return FLOAT
	case DIGIT_TYPE_DECIMAL:
		return INT
	case DIGIT_TYPE_SCIENTIFIC_NOTATION:
		return FLOAT
	case DIGIT_TYPE_HEXADECIMAL:
		return INT
	}
	return ILLEGAL
}
