package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitBooleanExpr(v *ast.Boolean) (result object.Object) {
	if v.Value {
		return object.TRUE
	}
	return object.FALSE
}
