package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitReturn(v *ast.ReturnStatement) (result object.Object) {
	if v.ReturnValue == nil {
		return &object.ReturnValue{Value: object.NULL}
	}

	result = i.evaluate(v.ReturnValue)
	if object.IsError(result) {
		return
	}

	return &object.ReturnValue{Value: result}
}
