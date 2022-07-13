package main

import (
	"fmt"
	"io/ioutil"
	"github.com/gravataLonga/ninja/evaluator"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
	"os"
	"strings"
	"testing"
)

var code = `function fib(n) { if (n < 2) { return n; } return fib(n-1) + fib(n-2); };`

var table = []struct {
	input int
	fib   int
}{
	{input: 100, fib: 5},
	{input: 1000, fib: 5},
	{input: 74382, fib: 5},
	{input: 382399, fib: 5},

	{input: 100, fib: 10},
	{input: 1000, fib: 10},
	{input: 74382, fib: 10},
	{input: 382399, fib: 10},

	{input: 100, fib: 20},
	{input: 1000, fib: 20},
	{input: 74382, fib: 20},
	{input: 382399, fib: 20},
}

func BenchmarkExecCode(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			// run the Fib function b.N times
			for n := 0; n < b.N; n++ {

				env := object.NewEnvironment()
				l := lexer.New(strings.NewReader(code + " fib(" + fmt.Sprint(v.fib) + "); "))
				p := parser.New(l)

				program := p.ParseProgram()
				if len(p.Errors()) > 0 {
					continue
				}
				evaluator.Eval(program, env)
			}
		})
	}
}

func TestMain_execCode(t *testing.T) {

	temporaryStdOut, fn, err := createStdInOut("TestMain_execCode")
	defer fn()
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCode", err)
	}

	execCode(`var a = 2 + 1; a;`, temporaryStdOut)

	resultOut, err := os.ReadFile(temporaryStdOut.Name())
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCode", err)
	}

	if string(resultOut) != "3" {
		fmt.Printf("--- stdout ---\n%s--- expected ---\n%s", resultOut, "3")
		t.Errorf("%s: stdout does not match expected", "TestMain_execCode")
	}
}

func TestMain_execCodeSpecialCharacter(t *testing.T) {
	temporaryStdOut, fn, err := createStdInOut("TestMain_execCode")
	defer fn()
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCode", err)
	}

	execCode("import \"./testdata/multiple_lines.ninja\"; input.split(\"\n\")", temporaryStdOut)

	resultOut, err := os.ReadFile(temporaryStdOut.Name())
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCode", err)
	}

	if string(resultOut) != "[29x13x26, 11x11x14, 27x2x5, 6x10x13, 15x19x10, 26x29x15, 8x23x6, 17x8x26, 20x28x3, 14x3x5, 10x9x8]" {
		fmt.Printf("--- stdout ---\n%s--- expected ---\n%s", resultOut, "3")
		t.Errorf("%s: stdout does not match expected", "TestMain_execCode")
	}
}

func TestMain_execCodeAssertions(t *testing.T) {
	temporaryStdOut, fn, err := createStdInOut("TestMain_execCodeAssertions")
	defer fn()
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCode", err)
	}

	code := readFile(t, "./testdata/assertions.ninja")
	expected := readFile(t, "./testdata/expected.txt")
	execCode(code, temporaryStdOut)

	resultOut, err := os.ReadFile(temporaryStdOut.Name())
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCodeAssertions", err)
	}

	if string(resultOut) != expected {
		t.Errorf("%s: stdout does not match expected. Output: %s", "TestMain_execCodeAssertions", resultOut)
	}
}

func readFile(t *testing.T, filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("unable to open file %s", filename)
	}

	return string(file)
}

func createStdInOut(name string) (*os.File, func(), error) {
	originalStdOut := os.Stdout
	temporaryStdOut, err := os.CreateTemp("", name)
	if err != nil {
		return nil, nil, err
	}

	os.Stdout = temporaryStdOut

	return temporaryStdOut, func() {
		defer os.Remove(temporaryStdOut.Name())
		os.Stdout = originalStdOut
	}, nil
}
