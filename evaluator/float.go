package evaluator

import "ninja/object"

func evalFloatInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: roundFloat(leftVal+rightVal, 10)}
	case "-":
		return &object.Float{Value: roundFloat(leftVal-rightVal, 10)}
	case "*":
		return &object.Float{Value: roundFloat(leftVal*rightVal, 10)}
	case "/":
		return &object.Float{Value: roundFloat(leftVal/rightVal, 10)}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
