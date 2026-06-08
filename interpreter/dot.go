package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitDotExpr(v *ast.Dot) (result object.Object) {
	obj := i.evaluate(v.Object)
	hash, ok := obj.(*object.Hash)
	if !ok {
		return object.NewErrorFormat("cannot access property %q on %s", v.Right.Value, obj.Type())
	}

	key := &object.String{Value: v.Right.Value}
	pair, ok := hash.Pairs[key.HashKey()]
	if !ok {
		return object.NULL
	}
	return pair.Value
}
