package evaluator

import (
	"math"
	"ninja/object"
)

func evalIntegerInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {

	leftObject := left.(*object.Integer)
	rightObject := right.(*object.Integer)

	leftVal := leftObject.Value
	rightVal := rightObject.Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "/":
		left := float64(leftVal)
		right := float64(rightVal)
		total := left / right

		// @todo check if we need a ELIPSON var or if there any way of doing
		if math.Round(total)-total <= 0.00000000000001 {
			return &object.Integer{Value: int64(total)}
		}

		return &object.Float{Value: left / right}
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
