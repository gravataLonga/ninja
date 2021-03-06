package parser

import (
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestImportStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`import "testing.mod";`, "import \"testing.mod\""},
		{`import "testing" + "other.txt";`, `import "(testing + other.txt)"`},
	}

	for _, tt := range tests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}
