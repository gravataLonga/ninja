package main

import (
	"fmt"
	"io/ioutil"
	"ninja/evaluator"
	"ninja/lexer"
	"ninja/object"
	"ninja/parser"
	"os"
	"testing"
)

/*
func Fib(n int) int {
        if n < 2 {
                return n
        }
        return Fib(n-1) + Fib(n-2)
}
*/

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
				l := lexer.New(code + " fib(" + fmt.Sprint(v.fib) + "); ")
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
	originalStdOut := os.Stdout
	temporaryStdOut, err := os.CreateTemp("", "TestMain_execCode")
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCode", err)
	}
	defer os.Remove(temporaryStdOut.Name())
	os.Stdout = temporaryStdOut

	execCode(`var a = 2 + 1; a;`, temporaryStdOut)

	resultOut, err := os.ReadFile(temporaryStdOut.Name())
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCode", err)
	}

	if string(resultOut) != "3" {
		os.Stdout = originalStdOut
		fmt.Printf("--- stdout ---\n%s--- expected ---\n%s", resultOut, "3")
		t.Errorf("%s: stdout does not match expected", "TestMain_execCode")
	}
}

func TestMain_execCodeAssertions(t *testing.T) {
	originalStdOut := os.Stdout
	temporaryStdOut, err := os.CreateTemp("", "TestMain_execCodeAssertions")
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCodeAssertions", err)
	}
	defer os.Remove(temporaryStdOut.Name())
	os.Stdout = temporaryStdOut

	code := readFile(t, "./tests/assertions.nj")
	expected := readFile(t, "./tests/expected.txt")
	execCode(code, temporaryStdOut)

	resultOut, err := os.ReadFile(temporaryStdOut.Name())
	if err != nil {
		t.Fatalf("%s: %s", "TestMain_execCodeAssertions", err)
	}

	if string(resultOut) != expected {
		os.Stdout = originalStdOut
		fmt.Printf("--- stdout ---\n%s--- expected ---\n%s", resultOut, expected)
		t.Errorf("%s: stdout does not match expected", "TestMain_execCodeAssertions")
	}
}

func readFile(t *testing.T, filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("unable to open file %s", filename)
	}

	return string(file)
}
