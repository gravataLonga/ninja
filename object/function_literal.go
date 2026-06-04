package object

import (
	"bytes"
	"strings"
)

type FunctionLiteral struct {
	Parameters []FunctionParameter
	Body       FunctionBody
	Closure    *Environment
}

func (f *FunctionLiteral) Type() ObjectType { return FUNCTION_OBJ }
func (f *FunctionLiteral) Inspect() string {
	var out bytes.Buffer
	params := make([]string, len(f.Parameters))
	for i, p := range f.Parameters {
		params[i] = p.String()
	}
	out.WriteString("function(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
