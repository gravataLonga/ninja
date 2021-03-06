package evaluator

import "github.com/gravataLonga/ninja/object"

func evalInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	switch {
	case operator == "&&":
		return nativeBoolToBooleanObject(object.IsTruthy(left) && object.IsTruthy(right))
	case operator == "||":
		return nativeBoolToBooleanObject(object.IsTruthy(left) || object.IsTruthy(right))
	case object.IsString(left) && object.IsString(right):
		return evalStringInfixExpression(operator, left, right)
	case object.IsNumber(left) && object.IsNumber(right):
		return evalFloatOrIntegerInfixExpression(operator, left, right)
	case operator == "==":
		// todo in future we can compare each array
		if object.IsArray(left) || object.IsArray(right) {
			return nativeBoolToBooleanObject(false)
		}

		if object.IsHash(left) || object.IsHash(right) {
			return nativeBoolToBooleanObject(false)
		}

		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		// todo in future we can compare each array
		if object.IsArray(left) || object.IsArray(right) {
			return nativeBoolToBooleanObject(false)
		}

		if object.IsHash(left) || object.IsHash(right) {
			return nativeBoolToBooleanObject(false)
		}
		return nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return object.NewErrorFormat("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
