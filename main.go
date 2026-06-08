package main

import (
	_ "embed"
	"fmt"

	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/interpreter"
	"github.com/gravataLonga/ninja/resolver"

	"io"
	"os"
	"strings"

	// "github.com/gravataLonga/ninja/evaluator"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
	"github.com/gravataLonga/ninja/repl"
	flag "github.com/spf13/pflag"
)

//go:embed version.txt
var version string

var inlineCode = flag.StringP("exec", "e", "", "Runs the given code.")
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
	object.Version = version

	var ninjaCode *ast.Program
	var errors []string

	if len(os.Args) == 1 {
		runRepl(os.Stdin, os.Stdout)
		return
	}

	if len(*inlineCode) > 0 {
		errors, ninjaCode = parseProgram(*inlineCode)
	} else {
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
		errors, ninjaCode = parseProgram(string(file))
	}

	if len(errors) > 0 {
		printParserErrors(errors, os.Stdout)
		os.Exit(1)
		return
	}

	if *astS {
		fmt.Printf("%s", ninjaCode.String())
		os.Exit(0)
		return
	}

	execCode(ninjaCode, os.Stdout)
}

func runRepl(in io.Reader, out io.Writer) {
	replProgram := repl.NewRepel(out, in)
	replProgram.Version(version)
	replProgram.Start()
}

func parseProgram(input string) (error []string, program *ast.Program) {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	program = p.ParseProgram()
	if len(p.Errors()) > 0 {
		return p.Errors(), nil
	}
	return nil, program
}

func execCode(ninjaCode *ast.Program, writer io.Writer) {
	env := object.NewEnvironment()

	i := interpreter.New(os.Stdout, env)
	resolver.NewResolver(i).Resolve(ninjaCode)
	result := i.Interpreter(ninjaCode)

	if result != nil {
		_, err := fmt.Fprintf(writer, result.Inspect())
		if err != nil {
			printParserErrors([]string{err.Error()}, writer)
			return
		}
	}
}

func printParserErrors(errors []string, writer io.Writer) {
	fmt.Fprintf(writer, "🔥 Fire at core! parser errors:")
	for _, msg := range errors {
		fmt.Fprintf(writer, "\t %s\n", msg)
	}
}
