package main

import (
	_ "embed"
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"ninja/evaluator"
	"ninja/lexer"
	"ninja/object"
	"ninja/parser"
	"ninja/repl"
	"os"
)

//go:embed version.txt
var version string

func main() {
	exec := flag.StringP("exec", "e", "", "Runs the given code.")
	_ = flag.BoolP("ast", "a", false, "Return AST structure")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Version: %s. \nUsage: ninja [flags] [program file] [arguments]\n\nAvailable flags:\n", version)

		flag.PrintDefaults()
	}

	flag.Parse()

	if len(os.Args) == 1 {
		runRepl(os.Stdin, os.Stdout)
		return
	}

	if len(*exec) > 0 {
		execCode(*exec, os.Stdout)
		return
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err == nil {
		execCode(string(file), os.Stdout)
		return
	}
}

func runRepl(in io.Reader, out io.Writer) {
	replProgram := repl.NewRepel(out, in)
	replProgram.Verbose(true)
	replProgram.Version(version)

	replProgram.Start()
}

func execCode(input string, writer io.Writer) {
	env := object.NewEnvironment()
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		printParserErrors(p.Errors(), writer)
		return
	}

	fmt.Println(program.String())

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Fprintf(writer, evaluated.Inspect())
	}
}

func printParserErrors(errors []string, writer io.Writer) {
	fmt.Fprintf(writer, "🔥 Fire at core! parser errors:")
	for _, msg := range errors {
		fmt.Fprintf(writer, "\t %s\n", msg)
	}
}
