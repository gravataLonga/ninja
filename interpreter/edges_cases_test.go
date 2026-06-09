package interpreter_test

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gravataLonga/ninja/interpreter"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
	"github.com/gravataLonga/ninja/resolver"
)

func runFixture(t *testing.T, name string) (result object.Object, panicked interface{}) {
	t.Helper()
	src, err := os.ReadFile(filepath.Join("..", "fixtures", name))
	if err != nil {
		t.Fatalf("reading fixture %q: %s", name, err)
	}
	l := lexer.New(strings.NewReader(string(src)))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		t.Fatalf("parsing fixture %q produced errors: %v", name, p.Errors())
	}
	i := interpreter.New(io.Discard, object.NewEnvironment())
	resolver.NewResolver(i).Resolve(program)
	defer func() { panicked = recover() }()
	result = i.Interpreter(program)
	return
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		fixture        string
		wantError      bool   // result must be *object.Error
		wantInspect    string // result.Inspect() must equal this (only when !wantError)
		errContains    string // substring required in error message
		errNotContains string // substring forbidden in error message
		errOpOnce      string // operator that must appear exactly once in error message
		anyOK          bool   // accept any non-panicking result
	}{
		// --- case-01: block scope depth mismatch ---
		{
			name:        "block scope depth — if branch",
			fixture:     "case-01-block-scope-depth.nj",
			wantInspect: "5",
		},
		{
			name:        "block scope depth — else branch",
			fixture:     "case-01b-block-scope-else.nj",
			wantInspect: "5",
		},
		{
			name:        "block scope depth — nested if",
			fixture:     "case-01c-block-scope-nested-if.nj",
			wantInspect: "5",
		},

		// --- case-02: modulo by zero ---
		{
			name:      "modulo by zero — positive dividend",
			fixture:   "case-02-modulo-by-zero.nj",
			wantError: true,
		},
		{
			name:      "modulo by zero — negative dividend",
			fixture:   "case-02b-modulo-by-zero-negative.nj",
			wantError: true,
		},
		{
			name:      "modulo by zero — zero dividend",
			fixture:   "case-02c-modulo-by-zero-zero.nj",
			wantError: true,
		},

		// --- case-03: array slice out of range ---
		{
			name:    "array slice out of range — offset beyond length",
			fixture: "case-03-array-slice-out-of-range.nj",
			anyOK:   true,
		},
		{
			name:    "array slice out of range — empty array",
			fixture: "case-03b-array-slice-empty.nj",
			anyOK:   true,
		},
		{
			name:    "array slice out of range — start beyond end",
			fixture: "case-03c-array-slice-start-beyond-end.nj",
			anyOK:   true,
		},

		// --- case-04: prefix operator mutates operand ---
		{
			name:        "prefix mutates operand — minus on integer",
			fixture:     "case-04-prefix-mutates-operand.nj",
			wantInspect: "5",
		},
		{
			name:        "prefix mutates operand — bang on boolean",
			fixture:     "case-04b-prefix-bang-mutates-boolean.nj",
			wantInspect: "true",
		},
		{
			name:        "prefix mutates operand — bang on array",
			fixture:     "case-04c-prefix-bang-mutates-array.nj",
			wantInspect: "[1, 2]",
		},

		// --- case-05: logical operators do not short-circuit ---
		{
			name:        "logical short-circuit — false && side effect",
			fixture:     "case-05-logical-no-short-circuit.nj",
			wantInspect: "false",
		},
		{
			name:        "logical short-circuit — true || side effect",
			fixture:     "case-05b-logical-or-no-short-circuit.nj",
			wantInspect: "false",
		},

		// --- case-06: dot expr masks underlying error ---
		{
			name:           "dot error masking — plus type error",
			fixture:        "case-06-dot-error-masking.nj",
			wantError:      true,
			errContains:    "unknown operator",
			errNotContains: "on ERROR",
		},
		{
			name:           "dot error masking — minus type error",
			fixture:        "case-06b-dot-error-masking-minus.nj",
			wantError:      true,
			errContains:    "unknown operator",
			errNotContains: "on ERROR",
		},

		// --- case-07: index assign on non-collection ---
		{
			name:      "index assign on non-collection — integer",
			fixture:   "case-07-index-assign-silent-noop.nj",
			wantError: true,
		},
		{
			name:      "index assign on non-collection — string",
			fixture:   "case-07b-index-assign-on-string.nj",
			wantError: true,
		},
		{
			name:      "index assign on non-collection — boolean",
			fixture:   "case-07c-index-assign-on-boolean.nj",
			wantError: true,
		},

		// --- case-08: assign statement swallows RHS error ---
		{
			name:      "assign swallows error — string+int RHS",
			fixture:   "case-08-assign-swallows-error.nj",
			wantError: true,
		},
		{
			name:      "assign swallows error — int>string RHS",
			fixture:   "case-08b-assign-swallows-rhs-error.nj",
			wantError: true,
		},

		// --- case-09: default param evaluated in caller scope ---
		{
			name:        "default param caller scope — two params",
			fixture:     "case-09-default-param-caller-scope.nj",
			wantInspect: "6",
		},
		{
			name:        "default param caller scope — three params chained",
			fixture:     "case-09b-default-param-third-references-second.nj",
			wantInspect: "7",
		},

		// --- case-10: integer pow with negative exponent ---
		{
			name:        "integer pow negative exponent — 2**-1",
			fixture:     "case-10-integer-pow-negative-exponent.nj",
			wantInspect: "0.500000",
		},
		{
			name:        "integer pow negative exponent — 2**-2",
			fixture:     "case-10b-integer-pow-negative-exp-minus-two.nj",
			wantInspect: "0.250000",
		},
		{
			name:        "integer pow negative exponent — 3**-1",
			fixture:     "case-10c-integer-pow-negative-exp-base-three.nj",
			wantInspect: "0.333333",
		},

		// --- case-11: duplicated operator in error message ---
		{
			name:      "duplicated operator in error — <",
			fixture:   "case-11-duplicated-operator-in-error.nj",
			wantError: true,
			errOpOnce: "<",
		},
		{
			name:      "duplicated operator in error — >",
			fixture:   "case-11b-duplicated-operator-gt.nj",
			wantError: true,
			errOpOnce: ">",
		},
		{
			name:      "duplicated operator in error — -",
			fixture:   "case-11c-duplicated-operator-minus.nj",
			wantError: true,
			errOpOnce: "-",
		},
		{
			name:      "duplicated operator in error — *",
			fixture:   "case-11d-duplicated-operator-mul.nj",
			wantError: true,
			errOpOnce: "*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, panicked := runFixture(t, tt.fixture)

			if panicked != nil {
				t.Fatalf("unexpected panic: %v", panicked)
			}

			if tt.anyOK {
				return
			}

			if tt.wantError {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("want *object.Error, got %T: %s", result, inspectObj(result))
				}
				if tt.errContains != "" && !strings.Contains(errObj.Message, tt.errContains) {
					t.Fatalf("error message should contain %q, got: %q", tt.errContains, errObj.Message)
				}
				if tt.errNotContains != "" && strings.Contains(errObj.Message, tt.errNotContains) {
					t.Fatalf("error message should not contain %q, got: %q", tt.errNotContains, errObj.Message)
				}
				if tt.errOpOnce != "" && strings.Count(errObj.Message, tt.errOpOnce) != 1 {
					t.Fatalf("operator %q should appear exactly once in error message, got: %q", tt.errOpOnce, errObj.Message)
				}
				return
			}

			if result == nil {
				t.Fatalf("want inspect %q, got nil", tt.wantInspect)
			}
			if got := result.Inspect(); got != tt.wantInspect {
				t.Fatalf("want %q, got %q (%T)", tt.wantInspect, got, result)
			}
		})
	}
}

func inspectObj(o object.Object) string {
	if o == nil {
		return "<nil>"
	}
	return o.Inspect()
}
