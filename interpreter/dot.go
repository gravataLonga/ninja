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
	// obj := i.evaluate(v.Object)
	// call, ok := v.Right.(*ast.Identifier)
	// if !ok {
	//	return object.NewErrorFormat("we expect to be a identifier on right of dot operation. Got: %t", v.Right)
	// }

	//objCallable, ok := obj.(object.CallableMethod)
	//if !ok {
	//	return object.NewErrorFormat("object must implement callable.")
	// }
	//fmt.Println(objCallable, call)
	//method, ok := call.Function.(*ast.Identifier)
	//if !ok {
	//	return object.NewErrorFormat("method name isn't a identifier")
	//}

	//args := i.evaluateExpressions(call.Arguments)

	// return objCallable.Call(method.Value, args...)
	return object.NULL
}
