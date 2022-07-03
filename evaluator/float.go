package evaluator

import (
	"math"
	"ninja/object"
)

func evalFloatInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {

	leftObject := left.(*object.Float)
	rightObject := right.(*object.Float)

	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "%":
		return &object.Float{Value: math.Mod(leftVal, rightVal)}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftObject.Compare(rightObject) == -1)
	case ">":
		return nativeBoolToBooleanObject(leftObject.Compare(rightObject) == 1)
	case "==":
		return nativeBoolToBooleanObject(leftObject.Compare(rightObject) == 0)
	case "!=":
		return nativeBoolToBooleanObject(leftObject.Compare(rightObject) != 0)
	case "<=":
		return nativeBoolToBooleanObject(leftObject.Compare(rightObject) <= 0)
	case ">=":
		return nativeBoolToBooleanObject(leftObject.Compare(rightObject) >= 0)
	default:
		return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
