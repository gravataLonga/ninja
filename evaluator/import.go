package evaluator

import (
	"ninja/ast"
	"ninja/lexer"
	"ninja/object"
	"ninja/parser"
	"os"
	"strings"
)

func evalImport(node ast.Node, env *object.Environment) object.Object {
	astImport, ok := node.(*ast.Import)
	if !ok {
		return object.NewErrorFormat("evalImport isnt type of ast.Import. Got: %t", node)
	}

	resultFilename := Eval(astImport.Filename, env)

	filename, ok := resultFilename.(*object.String)
	if !ok {
		return object.NULL
	}

	readFile, err := os.Open(filename.Value)

	if err != nil {
		return object.NewErrorFormat("IO Error: error reading file '%s': %s", filename.Value, err)
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

	result := Eval(programs, env)

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
