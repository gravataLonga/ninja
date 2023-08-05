package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
	"os"
	"strings"
)

func (i *Interpreter) VisitImportExpr(v *ast.Import) (result object.Object) {

	resultFilename := i.evaluate(v.Filename)

	filename, ok := resultFilename.(*object.String)
	if !ok {
		return object.NULL
	}

	readFile, err := os.Open(filename.Value)

	if err != nil {
		return object.NewErrorFormat("IO Error: error reading file '%s': %s %s", filename.Value, err, v.Token)
	}

	l := lexer.New(readFile)
	p := parser.New(l)
	programs := p.ParseProgram()

	if len(p.Errors()) > 0 {
		strErros := []string{}
		for _, e := range p.Errors() {
			strErros = append(strErros, e)
		}
		return object.NewErrorFormat("%s: %s", filename.Value, strings.Join(strErros, "\n"))
	}

	result = i.execute(programs)

	if result == nil {
		return nil
	}

	errorStr, ok := result.(*object.Error)
	if ok {
		return object.NewErrorFormat("%s: %s", filename.Value, errorStr.Message)
	}

	// Only return if last item of imported file have "return"
	if len(programs.Statements) > 0 {
		stmts := programs.Statements
		stmt := stmts[len(stmts)-1]

		_, ok = stmt.(*ast.ReturnStatement)
		if ok {
			return result
		}
	}

	return object.NULL
}
