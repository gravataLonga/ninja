package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gravataLonga/ninja/lexer"
)

func parseAndGetErrors(input string) []string {
	l := lexer.New(strings.NewReader(input))
	p := New(l)
	p.ParseProgram()
	return p.Errors()
}

// TestParserPeekErrorTrailingPeriod documents that peekError() in parser.go
// appends a trailing period to "expected … got … instead." messages.
// Per the uniform style (Rule 2), error messages should not end with a period.
// Note: inputs with `var` require a trailing semicolon to avoid the infinite-loop
// recovery in parseVarStatement that scans until it finds a semicolon.
func TestParserPeekErrorTrailingPeriod(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`var = 1;`,
			`expected next token to be IDENT, got = at [Line: 1, Offset: 5] instead.`,
		},
		{
			`enum { case A: 1 }`,
			`expected next token to be IDENT, got { at [Line: 1, Offset: 6] instead.`,
		},
		{
			`for (var i = 0; i <= parts-1; i = i + 1)`,
			`expected next token to be {, got EOF at [Line: 1, Offset: 41] instead.`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestParserPeekErrorTrailingPeriod[%d]", i), func(t *testing.T) {
			errs := parseAndGetErrors(tt.input)
			if len(errs) == 0 {
				t.Fatalf("expected at least 1 error, got 0")
			}
			if errs[0] != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errs[0])
			}
		})
	}
}

// TestParserReturnErrorCapitalizationAndPeriod documents that parseReturnStatement()
// in return.go produces a message with a capitalized first word and a trailing period.
// Per the uniform style (Rule 2), messages should start lowercase and have no trailing period.
func TestParserReturnErrorCapitalizationAndPeriod(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`return >`,
			`Next token expected to be nil or expression. Got: > at [Line: 1, Offset: 8].`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestParserReturnErrorCapitalizationAndPeriod[%d]", i), func(t *testing.T) {
			errs := parseAndGetErrors(tt.input)
			if len(errs) == 0 {
				t.Fatalf("expected at least 1 error, got 0")
			}
			if errs[0] != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errs[0])
			}
		})
	}
}

// TestParserEnumFatalErrorPrefix documents that parseEnum() in enum.go
// uses the "Fatal error:" category prefix on duplicate-identifier errors.
// Per the uniform style (Rule 1), category prefixes should be dropped.
func TestParserEnumFatalErrorPrefix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`enum T { case A: 1; case A: 2 }`,
			`Fatal error: Cannot redefine identifier A`,
		},
		{
			`enum Colors { case Red: 1; case Blue: 2; case Red: 3 }`,
			`Fatal error: Cannot redefine identifier Red`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestParserEnumFatalErrorPrefix[%d]", i), func(t *testing.T) {
			errs := parseAndGetErrors(tt.input)
			if len(errs) == 0 {
				t.Fatalf("expected at least 1 error, got 0")
			}
			if errs[0] != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errs[0])
			}
		})
	}
}

// TestParserArgumentsGrammarError documents that parseParameterWithOptional()
// in arguments.go emits an unclear and grammatically incorrect error message
// when a required parameter follows a default parameter.
// Per the uniform style, the message should be clear and grammatically correct.
func TestParserArgumentsGrammarError(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`var f = function(a = 1, b) {};`,
			`require arguments must be on declare first`,
		},
		{
			`function f(x = 0, y) {}`,
			`require arguments must be on declare first`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestParserArgumentsGrammarError[%d]", i), func(t *testing.T) {
			errs := parseAndGetErrors(tt.input)
			if len(errs) == 0 {
				t.Fatalf("expected at least 1 error, got 0")
			}
			if errs[0] != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errs[0])
			}
		})
	}
}

// TestParserNoPrefixFnInternalJargon documents that noPrefixParseFnError() in
// parser.go exposes internal implementation detail ("prefix parse function")
// in the user-facing error message.
// Per the uniform style (Category 1), messages should not name internal parser state.
func TestParserNoPrefixFnInternalJargon(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`var x = >;`,
			`no prefix parse function for > found`,
		},
		{
			`var x = >=;`,
			`no prefix parse function for >= found`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestParserNoPrefixFnInternalJargon[%d]", i), func(t *testing.T) {
			errs := parseAndGetErrors(tt.input)
			if len(errs) == 0 {
				t.Fatalf("expected at least 1 error, got 0")
			}
			if errs[0] != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errs[0])
			}
		})
	}
}

// TestParserIllegalAssignmentQuoteStyle documents that parseExpressionOrAssignStatement()
// in statement.go uses %q (Go double-quote style) to quote values in the error message.
// Per the uniform style (Rule 4), user-supplied values should use %q and type names
// should use backticks — but the inconsistency here is the mixed quoting approach
// across different error sites in the codebase.
func TestParserIllegalAssignmentQuoteStyle(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`1 = 2;`,
			`illegal "2" assignment to "1"`,
		},
		{
			`true = false;`,
			`illegal "false" assignment to "true"`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestParserIllegalAssignmentQuoteStyle[%d]", i), func(t *testing.T) {
			errs := parseAndGetErrors(tt.input)
			if len(errs) == 0 {
				t.Fatalf("expected at least 1 error, got 0")
			}
			if errs[0] != tt.expected {
				t.Errorf("wrong error message.\n\texpected=%q\n\tgot=%q", tt.expected, errs[0])
			}
		})
	}
}
