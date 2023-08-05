package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitIntegerExpr(v *ast.IntegerLiteral) (result object.Object) {
	return &object.Integer{Value: v.Value}
}

func (i *Interpreter) VisitFloatExpr(v *ast.FloatLiteral) (result object.Object) {
	return &object.Float{Value: v.Value}
}
