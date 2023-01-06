package evaluator

import "github.com/gravataLonga/ninja/object"

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ArrayObj && index.Type() == object.IntegerObj:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.HashObj:
		return evalHashIndexExpression(left, index)
	case left.Type() == object.StringObj:
		return evalStringIndexExpression(left, index)
	default:
		return object.NewErrorFormat("index operator not supported: %s", left.Type())
	}
}
