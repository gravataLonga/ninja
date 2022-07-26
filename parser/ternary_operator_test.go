package parser

import (
	"fmt"
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestTernaryOperatorProcedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`true ? 1 : 0`,
			`(true?1:0)`,
		},
		{
			`1 > 2 ? 30 * 50 : add()`,
			`((1 > 2)?(30 * 50):add())`,
		},
		{
			`add() ? 30 ** 40 << 1 + 2 : (true ? false : 0)`,
			`(add()?((30 ** 40) << (1 + 2)):(true?false:0))`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestTernaryOperatorProcedence[%d]", i), func(t *testing.T) {
			l := lexer.New(strings.NewReader(tt.input))
			p := New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if tt.expected != program.String() {
				t.Fatalf("Program didn't produce expected %s. Got: %s", tt.expected, program.String())
			}
		})
	}
}
