package main

import (
	_ "embed"
	"fmt"
	"github.com/gravataLonga/ninja/interpreter"

	// "github.com/gravataLonga/ninja/evaluator"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
	"github.com/gravataLonga/ninja/repl"
	flag "github.com/spf13/pflag"
	"io"
	"os"
	"strings"
)

//go:embed version.txt
var version string

var exec = flag.StringP("exec", "e", "", "Runs the given code.")
var astS = flag.BoolP("ast", "a", false, "Return AST structure")

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Version: %s. \nUsage: ninja [flags] [program file] [arguments]\n\nAvailable flags:\n", version)

		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	object.StandardInput = os.Stdin
	object.StandardOutput = os.Stdout
	object.Arguments = args
	object.ExitFunction = os.Exit

	if len(os.Args) == 1 {
		runRepl(os.Stdin, os.Stdout)
		return
	}

	if len(*exec) > 0 {
		execCode(*exec, os.Stdout)
		return
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%v", err)
		if err != nil {
			os.Exit(1)
			return
		}
		os.Exit(1)
		return
	}

	execCode(string(file), os.Stdout)
}

func runRepl(in io.Reader, out io.Writer) {
	replProgram := repl.NewRepel(out, in)
	replProgram.Version(version)

	replProgram.Start()
}

func execCode(input string, writer io.Writer) {
	env := object.NewEnvironment()
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	// s := semantic.New()

	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		printParserErrors(p.Errors(), writer)
		return
	}

	/*program = s.Analysis(program)
	if len(s.Errors()) != 0 {
		printSemanticErrorsErrors(s.Errors(), writer)
		return
	}*/
	i := interpreter.New(os.Stdout, env)
	result := i.Interpreter(program)
	if result != nil {
		fmt.Fprintf(writer, result.Inspect())
	}
}

func printParserErrors(errors []string, writer io.Writer) {
	fmt.Fprintf(writer, "🔥 Fire at core! parser errors:")
	for _, msg := range errors {
		fmt.Fprintf(writer, "\t %s\n", msg)
	}
}

func printSemanticErrorsErrors(errors []string, writer io.Writer) {
	fmt.Fprintf(writer, "🔥 Fire at core! semantic errors:")
	for _, msg := range errors {
		fmt.Fprintf(writer, "\t %s\n", msg)
	}
}
