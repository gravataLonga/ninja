package main

import (
	"fmt"
	"io/ioutil"
	"ninja/evaluator"
	"ninja/lexer"
	"ninja/object"
	"ninja/parser"
	"ninja/repl"
	"os"
	user2 "os/user"

	flag "github.com/spf13/pflag"
)

func main() {
	exec := flag.StringP("exec", "e", "", "Runs the given code.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: ninja [flags] [program file] [arguments]\n\nAvailable flags:\n")

		flag.PrintDefaults()
	}

	flag.Parse()

	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	if len(*exec) > 0 {
		execCode(*exec)
		return
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err == nil {
		execCode(string(file))
		return
	}

}

func runRepl() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Ninja Programming Language\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}

func execCode(input string) {
	env := object.NewEnvironment()
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		printParserErrors(p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}

func printParserErrors(errors []string) {
	fmt.Println("🔥 Fire at core!")
	fmt.Println(" parser errors:")
	for _, msg := range errors {
		fmt.Printf("\t %s\n", msg)
	}
}
