package interpreter

import (
	"bytes"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
	"strings"
	"testing"
)

func runCapture(t *testing.T, input string) string {
	t.Helper()
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		t.Fatalf("parse errors: %v", p.Errors())
	}
	var out bytes.Buffer
	object.StandardOutput = &out
	i := New(&out, object.NewEnvironment())
	NewResolver(i).Resolve(program)
	i.Interpreter(program)
	return strings.TrimSpace(out.String())
}

// TestResolverClosureScoping verifies that an inner function reads the variable
// from its defining scope, not from a later shadowing binding.
func TestResolverClosureScoping(t *testing.T) {
	input := `
var x = "global";
function outer() {
    var x = "outer";
    function inner() { puts(x); }
    inner();
}
outer();
`
	got := runCapture(t, input)
	if got != "outer" {
		t.Fatalf("expected closure to read outer x, got %q", got)
	}
}

// TestResolverMutateOuterVariable verifies that assigning to a variable from
// an outer scope mutates the correct binding (not a new local shadow).
func TestResolverMutateOuterVariable(t *testing.T) {
	input := `
function makeCounter() {
    var count = 0;
    return function() { count = count + 1; puts(count); };
}
var counter = makeCounter();
counter();
counter();
`
	got := runCapture(t, input)
	lines := strings.Split(got, "\n")
	if len(lines) != 2 || lines[0] != "1" || lines[1] != "2" {
		t.Fatalf("expected counter to print 1 then 2, got %v", lines)
	}
}
