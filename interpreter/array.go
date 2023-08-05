package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitArrayExpr(v *ast.ArrayLiteral) (result object.Object) {
	elements := i.evaluateExpressions(v.Elements)
	if len(elements) == 1 && object.IsError(elements[0]) {
		return elements[0]
	}
	return &object.Array{Elements: elements}
}

// removeIndexFromArray is slow operations, we need better way?
func removeIndexFromArray(slice []object.Object, s int64) []object.Object {

	copy(slice[s:], slice[s+1:])      // Shift a[i+1:] left one index.
	slice[len(slice)-1] = object.NULL // Erase last element (write zero value).
	slice = slice[:len(slice)-1]      // Truncate slice.

	return slice
}
