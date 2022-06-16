package evaluator

import "ninja/object"

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "++":
		return evalIncrementPrefixOperatorExpression(right)
	case "--":
		return evalDecrementPrefixOperatorExpression(right)
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}
