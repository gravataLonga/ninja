package repl

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
	"github.com/gravataLonga/ninja/interpreter"

	// "github.com/gravataLonga/ninja/evaluator"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
	"io"
	"os"
	"os/user"
	"strings"

	color "github.com/fatih/color"
)

//go:embed ninja-logo.png
var logoImage []byte

const PROMPT = ">>> "

const NINJA_LICENSE = "Ninja Language - MIT LICENSE - Version: %s\n"

type Repl struct {
	out     io.Writer
	in      io.Reader
	scan    *bufio.Scanner
	env     *object.Environment
	version string
}

var colorName = map[string]*color.Color{
	"normal":  color.New(color.FgWhite),
	"program": color.New(color.FgWhite, color.Bold),
	"brand":   color.New(color.FgHiBlue, color.Bold),
	"error":   color.New(color.FgRed),
}

func NewRepel(out io.Writer, in io.Reader) *Repl {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	return &Repl{out: out, in: in, scan: scanner, env: env}
}

func (r *Repl) Version(vs string) {
	r.version = strings.Replace(vs, "\n", "", 0)
}

func (r *Repl) Output(levelOutput string, format string, a ...interface{}) {
	c, ok := colorName[levelOutput]
	if !ok {
		c = colorName["normal"]
	}

	c.Fprintf(r.out, format, a...)
}

func (r *Repl) Start() {
	user2, err := user.Current()
	if err != nil {
		panic(err)
	}

	r.printSplashScreen()
	r.printSplashLicense()

	r.Output("program", "Hi %s! This is Ninja Programming Language\n", user2.Username)
	r.Output("program", "Feel free to type in commands\n")
	r.Output("program", "If found an error, open issue at github.com/gravataLonga/ninja\n")

	for {
		r.Output("normal", PROMPT)
		scanned := r.scan.Scan()
		if !scanned {
			return
		}

		line := r.scan.Text()
		l := lexer.New(strings.NewReader(line))
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			r.printParserErrors(p.Errors())
			continue
		}

		i := interpreter.New(os.Stdout, r.env)

		evaluated := i.Interpreter(program)

		if _, ok := evaluated.(*object.Error); ok {
			r.Output("error", evaluated.Inspect())
			r.Output("program", "\n")
			continue
		}

		if evaluated != nil {
			r.Output("program", evaluated.Inspect())
			r.Output("program", "\n")
		}
	}
}

func (r *Repl) printSplashLicense() {
	r.Output("brand", NINJA_LICENSE, r.version)
}

func (r *Repl) printSplashScreen() {
	fmt.Fprint(r.out, "\n\n")
	fmt.Fprint(r.out, createSpashScreen())
	fmt.Fprint(r.out, "\n\n")
}

func (r *Repl) printParserErrors(errors []string) {
	r.Output("error", "We got some parser errors.\n")
	r.Output("error", "\tparser errors:\n")
	for _, msg := range errors {
		r.Output("error", "\t\t"+msg+"\n")
	}
}

func (r *Repl) printSemanticErrors(errors []string) {
	r.Output("error", "We got some semantic errors.\n")
	r.Output("error", "\tsemantic errors:\n")
	for _, msg := range errors {
		r.Output("error", "\t\t"+msg+"\n")
	}
}

func createSpashScreen() string {

	file, err := os.CreateTemp("", "repl_logo")
	file.Write(logoImage)
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	flags := aic_package.DefaultFlags()

	flags.Width = 50
	flags.Colored = true
	flags.CustomMap = " .-+$*"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}

	// Conversion for an image
	asciiArt, err := aic_package.Convert(file.Name(), flags)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%v\n", asciiArt)
}
