package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitStringExpr(v *ast.StringLiteral) (result object.Object) {
	return &object.String{Value: v.Value}
}
