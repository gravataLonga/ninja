package object

import (
	"bytes"
	"ninja/ast"
	"strings"
)

type FunctionLiteral struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *FunctionLiteral) Type() ObjectType { return FUNCTION_OBJ }
func (f *FunctionLiteral) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("function")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
