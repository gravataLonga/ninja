package interpreter_test

import (
	"fmt"
	"testing"

	"github.com/gravataLonga/ninja/object"
)

// TestPostfixErrorMessageCapitalization documents that postfixExpression() in
// postfix.go produces "Postfix operation not allowed on …" with a capitalized
// first word and no source location.
// Per the uniform style (Rule 2, Rule 5, Category 5), operator errors should be
// lowercase and append location via v.Token.HumanLocation() at the outermost call site.
// Token: v.Token is the ++ / -- operator; its Offset is the position of the second char.
func TestPostfixErrorMessageCapitalization(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// "hello" = 7 chars; ++ second char at offset 9
		{
			`"hello"++`,
			`Postfix operation not allowed on STRING at [Line: 1, Offset: 9]`,
		},
		// true = 4 chars; ++ second char at offset 6
		{
			`true++`,
			`Postfix operation not allowed on BOOLEAN at [Line: 1, Offset: 6]`,
		},
		// true = 4 chars; -- second char at offset 6
		{
			`true--`,
			`Postfix operation not allowed on BOOLEAN at [Line: 1, Offset: 6]`,
		},
		// "var s = \"ninja\"; " = 18 chars; -- second char at offset 20
		{
			`var s = "ninja"; s--`,
			`Postfix operation not allowed on STRING at [Line: 1, Offset: 20]`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestPostfixErrorMessageCapitalization[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}
			if errObj.Message != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errObj.Message)
			}
		})
	}
}

// TestDeleteStatementErrorMessagePrefix documents that VisitDelete() in
// delete_statement.go prefixes all messages with "DeleteStatement." — an
// internal Go type name that should not appear in user-facing messages.
// Per the uniform style (Rule 1, Rule 2, Rule 5), prefixes should be dropped,
// messages should use sentence-case without trailing periods, and location
// should be appended via v.Token.HumanLocation().
// Token: v.Token is the `delete` keyword; its Offset is the position of the
// first character after the keyword (whitespace), as the lexer records position
// of the character read after completing an identifier.
func TestDeleteStatementErrorMessagePrefix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// "delete " = 7 chars; `delete` keyword offset 7
		{
			`delete b[0]`,
			`DeleteStatement.left b identifier not found. at [Line: 1, Offset: 7]`,
		},
		// "var a = []; " = 12 chars, "delete" at 13-18, offset 19
		{
			`var a = []; delete a[{}]`,
			`DeleteStatement.index must be a Integer. Got: *object.Hash at [Line: 1, Offset: 19]`,
		},
		// "var a = [0, 1, 2]; " = 19 chars, "delete" at 20-25, offset 26
		{
			`var a = [0, 1, 2]; delete a[5]`,
			`DeleteStatement.index must be equal or less than the total of the items. Got: *object.Integer at [Line: 1, Offset: 26]`,
		},
		{
			`var a = [0, 1, 2]; delete a[-1]`,
			`DeleteStatement.index must be equal or greater than the 0. Got: *object.Integer at [Line: 1, Offset: 26]`,
		},
		// "var a = \"\"; " = 12 chars, "delete" at 13-18, offset 19
		{
			`var a = ""; delete a[0]`,
			`DeleteStatement.left only work with array or hash object. Got: *object.String at [Line: 1, Offset: 19]`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestDeleteStatementErrorMessagePrefix[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}
			if errObj.Message != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errObj.Message)
			}
		})
	}
}

// TestEnumErrorMessageGrammar documents two grammar issues in enum.go:
//  1. "identifier %s don't exists" should be "identifier %q doesn't exist".
//  2. The identifier-not-found message uses string concatenation instead of
//     a format call, producing a bare message with no source location.
//
// Per the uniform style (Rule 5, Rule 6, Rule 7), grammar must be correct, the
// same concept ("identifier not found") must read identically everywhere, and
// location must be appended via v.Token.HumanLocation().
// Token: v.Token is the `::` DOUBLE_COLON operator; its Offset is the position
// of the second colon (lexer records position of second char for 2-char tokens).
func TestEnumErrorMessageGrammar(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// "enum T { case A: 1 }; T" = 23 chars; "::" second colon at offset 25
		{
			`enum T { case A: 1 }; T::B`,
			`identifier B don't exists on enum object at [Line: 1, Offset: 25]`,
		},
		// "enum Status { case Active: 1; case Inactive: 0 }; Status" = 56 chars; "::" second colon at offset 58
		{
			`enum Status { case Active: 1; case Inactive: 0 }; Status::Unknown`,
			`identifier Unknown don't exists on enum object at [Line: 1, Offset: 58]`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestEnumErrorMessageGrammar[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}
			if errObj.Message != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errObj.Message)
			}
		})
	}
}

// TestIndexErrorMessageTypo documents that indexStringExpression() in index.go
// uses the misspelled word "isnt" instead of "isn't" (or a better phrasing).
// Per the uniform style (Rule 5, Rule 6), grammar errors should be corrected and
// location should be appended via v.Token.HumanLocation() in VisitIndexExpr.
// Token: v.Token is the `[` bracket; its Offset is the position of `[` itself
// (single-char tokens record the position of the character, not the char after).
func TestIndexErrorMessageTypo(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// "hello" = 7 chars; `[` at offset 8
		{
			`"hello"[1.5]`,
			`index isnt integer: FLOAT at [Line: 1, Offset: 8]`,
		},
		// "ninja" = 7 chars; `[` at offset 8
		{
			`"ninja"[true]`,
			`index isnt integer: BOOLEAN at [Line: 1, Offset: 8]`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestIndexErrorMessageTypo[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}
			if errObj.Message != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errObj.Message)
			}
		})
	}
}

// TestFunctionArityErrorMessageCapitalization documents that validateArguments()
// in function.go produces "Function expected …" with a capitalized first word
// and embeds the source location inside the format string via v.Token.String()
// (which includes the token TYPE "(" before the position) instead of using
// v.Token.HumanLocation() (location only).
// Per the uniform style (Rule 2, Rule 5), messages should be lowercase and
// location should be appended via HumanLocation(), not include the token type.
func TestFunctionArityErrorMessageCapitalization(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// `(` call token at offset 16; desired: HumanLocation only (no "( " prefix)
		{
			`function (x) {}();`,
			`Function expected 1 parameters, got 0 at [Line: 1, Offset: 16]`,
		},
		// `(` call token at offset 15
		{
			`function () {}(0);`,
			`Function expected 0 parameters, got 1 at [Line: 1, Offset: 15]`,
		},
		// `(` call token at offset 20
		{
			`function (x, y) { }(1, 2, 3);`,
			`Function expected 2 parameters, got 3 at [Line: 1, Offset: 20]`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestFunctionArityErrorMessageCapitalization[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}
			if errObj.Message != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errObj.Message)
			}
		})
	}
}

// TestImportIOErrorPrefix documents that VisitImportExpr() in import.go
// uses the "IO Error:" category prefix and wraps the file path in single quotes
// instead of using %q (double-quote style).
// It also uses v.Token.String() (includes token type "IMPORT") instead of
// v.Token.HumanLocation() (location only), making the message verbose and inconsistent.
// Per the uniform style (Rule 1, Rule 4, Rule 5), prefixes should be dropped,
// values should use %q, and location should come from HumanLocation() (no token type).
// Token: v.Token is the `import` keyword; offset = position of char after the keyword.
func TestImportIOErrorPrefix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// "import" = 6 chars; offset 7 = position of space after keyword.
		// Desired: HumanLocation only — drop "IMPORT " that Token.String() prepends to "at [...]"
		{
			`import "nonexistent_xyz_file.ninja"`,
			`IO Error: error reading file 'nonexistent_xyz_file.ninja': open nonexistent_xyz_file.ninja: no such file or directory at [Line: 1, Offset: 7]`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestImportIOErrorPrefix[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}
			if errObj.Message != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errObj.Message)
			}
		})
	}
}

// TestIdentifierNotFoundInconsistency documents that "identifier not found" is
// produced in two different formats depending on where it originates:
//
//   - VisitIdentExpr (assign_statement.go): uses v.Token.String() which embeds
//     the token TYPE "IDENT" before the position — e.g.
//     "identifier not found: add IDENT at [Line: 1, Offset: 25]"
//
//   - VisitScopeOperatorExpression (enum.go): uses bare string concatenation,
//     producing no location at all — e.g. "identifier not found: undeclaredEnum"
//
// Per the uniform style (Rule 5, Rule 7), the same concept must read identically
// everywhere using HumanLocation() (no token type in the location string):
//
//	"identifier not found: <name> at [Line: X, Offset: Y]"
//
// Token for VisitIdentExpr: v.Token is the IDENT; offset = position of char
// after the identifier (lexer records position of char read after the identifier loop).
// Token for VisitScopeOperatorExpression: v.Token is the `::` DOUBLE_COLON;
// offset = position of second colon.
func TestIdentifierNotFoundInconsistency(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Via VisitIdentExpr — should use HumanLocation() (drop "IDENT " token type)
		// "add" ends at offset 24; char after = `(` at offset 25
		{
			`function () { return add(); }();`,
			`identifier not found: add at [Line: 1, Offset: 25]`,
		},
		// Via VisitScopeOperatorExpression — currently has no location at all
		// "undeclaredEnum" = 14 chars; "::" second colon at offset 16
		{
			`undeclaredEnum::A`,
			`identifier not found: undeclaredEnum at [Line: 1, Offset: 16]`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestIdentifierNotFoundInconsistency[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}
			if errObj.Message != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errObj.Message)
			}
		})
	}
}

// TestAssignErrorMessageQuoteStyle documents that assign_statement.go uses %q
// for property names (producing Go double-quote style) but leaves the type name
// unquoted, while the audit (Rule 4) calls for backtick-quoting type names.
// It also documents the grammatically broken index-out-of-range message.
// Per the uniform style (Rule 5), all errors should append location via
// v.Token.HumanLocation(). The parser must store the `=` token in AssignStatement
// (not the first RHS token) so that the reported position is the assignment operator.
// Token offset for `=`: position of the `=` character itself (single-char token).
func TestAssignErrorMessageQuoteStyle(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// "var x = 1; x.name " = 18 chars; `=` at offset 19
		{
			`var x = 1; x.name = "test"`,
			`cannot assign property "name" on INTEGER at [Line: 1, Offset: 19]`,
		},
		{
			`var a = [1, 2, 3]; a[-1] = 0`,
			`index out of range, got -1 not positive index at [Line: 1, Offset: 26]`,
		},
		{
			`var a = [1]; a[5] = 0`,
			`index out of range, got 5 but array has only 1 elements at [Line: 1, Offset: 19]`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestAssignErrorMessageQuoteStyle[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}
			if errObj.Message != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errObj.Message)
			}
		})
	}
}
