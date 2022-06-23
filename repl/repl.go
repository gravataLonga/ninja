package repl

import (
	"bufio"
	"io"
	"ninja/evaluator"
	"ninja/lexer"
	"ninja/object"
	"ninja/parser"
	"os/user"
	"strings"

	color "github.com/fatih/color"
)

const PROMPT = ">> "

const NINJA_LICENSE = " License 2022 - Built by Jonathan Fontes - Version: %s"

const NINJA_SPLASH = `


$$\   $$\ $$\                         
$$$\  $$ |\__|                        
$$$$\ $$ |$$\ $$$$$$$\  $$\  $$$$$$\  
$$ $$\$$ |$$ |$$  __$$\ \__| \____$$\ 
$$ \$$$$ |$$ |$$ |  $$ |$$\  $$$$$$$ |
$$ |\$$$ |$$ |$$ |  $$ |$$ |$$  __$$ |
$$ | \$$ |$$ |$$ |  $$ |$$ |\$$$$$$$ |
\__|  \__|\__|\__|  \__|$$ | \_______|
                  $$\   $$ |          
                  \$$$$$$  |          
                   \______/           

`

type repl struct {
	out     io.Writer
	in      io.Reader
	scan    *bufio.Scanner
	env     *object.Environment
	verbose bool
	version string
}

var colorName = map[string]*color.Color{
	"normal":  color.New(color.FgWhite),
	"program": color.New(color.FgWhite, color.Bold),
	"brand":   color.New(color.FgCyan, color.Bold),
	"error":   color.New(color.FgRed),
}

func NewRepel(out io.Writer, in io.Reader) *repl {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	return &repl{out: out, in: in, verbose: false, scan: scanner, env: env}
}

func (r *repl) Verbose(state bool) {
	r.verbose = state
}

func (r *repl) Version(vs string) {
	r.version = strings.Replace(vs, "\n", "", 0)
}

func (r *repl) Output(levelOutput string, format string, a ...interface{}) {
	c, ok := colorName[levelOutput]
	if !ok {
		c = colorName["normal"]
	}

	c.Fprintf(r.out, format, a...)
}

func (r *repl) Start() {
	user2, err := user.Current()
	if err != nil {
		panic(err)
	}

	r.printSplashLicense()
	r.printSplashScreen()

	r.Output("normal", "Hi %s! This is Ninja Programming Language", user2.Username)
	r.Output("normal", "Feel free to type in commands\n")

	for {
		r.Output("normal", PROMPT)
		scanned := r.scan.Scan()
		if !scanned {
			return
		}

		line := r.scan.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			r.printParserErrors(p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, r.env)

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

func (r *repl) printSplashLicense() {
	r.Output("brand", NINJA_LICENSE, r.version)
}

func (r *repl) printSplashScreen() {
	r.Output("brand", NINJA_SPLASH)
}

func (r *repl) printParserErrors(errors []string) {
	r.Output("error", "We got some parser errors.\n")
	r.Output("error", "\tparser errors:\n")
	for _, msg := range errors {
		r.Output("error", "\t\t"+msg+"\n")
	}
}
