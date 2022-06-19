package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // be a string don't have same performance as using int or byte
}

func (t *TokenType) String() string {
	return string(*t)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	INCRE    = "++"
	DECRE    = "--"

	LT  = "<"
	GT  = ">"
	LTE = "<="
	GTE = ">="
	EQ  = "=="
	NEQ = "!="
	AND = "&&"
	OR  = "||"

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	FUNCTION         = "FUNCTION"
	FUNCTION_LITERAL = "FUNCTION_LITERAL"
	VAR              = "VARIABLE"
	TRUE             = "TRUE"
	FALSE            = "FALSE"
	IF               = "IF"
	ELSEIF           = "ELSEIF"
	ELSE             = "ELSE"
	RETURN           = "RETURN"
	IMPORT           = "IMPORT"
	FOR              = "FOR"
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
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func GuessTypeOfDigit(digit string) TokenType {
	hasDot := strings.IndexByte(digit, '.')
	if hasDot >= 0 {
		return FLOAT
	}
	return INT
}
