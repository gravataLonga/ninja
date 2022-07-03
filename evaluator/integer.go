package evaluator

import "ninja/object"

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
		return &object.Integer{Value: leftVal / rightVal}
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
