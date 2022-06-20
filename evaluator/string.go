package evaluator

import "ninja/object"

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return &object.Boolean{Value: leftVal == rightVal}
	case "!=":
		return &object.Boolean{Value: leftVal != rightVal}
	}

	return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())

}

func evalStringIndexExpression(str, index object.Object) object.Object {
	stringObject := str.(*object.String)

	idx, ok := index.(*object.Integer)
	if !ok {
		return object.NewErrorFormat("index isnt integer: %s", index.Type())
	}

	rn := []rune(stringObject.Value)

	if idx.Value <= -1 {
		return object.NULL
	}

	if int64(len(rn))-1 < idx.Value {
		return object.NULL
	}

	return &object.String{
		Value: string([]rune(stringObject.Value)[idx.Value]),
	}
}
