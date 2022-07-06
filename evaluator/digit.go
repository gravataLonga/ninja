package evaluator

import "ninja/object"

func evalFloatOrIntegerInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftValFloat, _ := left.(*object.Float)
	rightValFloat, _ := right.(*object.Float)

	leftValInteger, okLeftInteger := left.(*object.Integer)
	rightValInteger, okRightInteger := right.(*object.Integer)

	if okLeftInteger && okRightInteger {
		return evalIntegerInfixExpression(operator, leftValInteger, rightValInteger)
	}

	if okLeftInteger {
		leftValFloat = &object.Float{
			Value: float64(leftValInteger.Value),
		}
	}

	if okRightInteger {
		rightValFloat = &object.Float{
			Value: float64(rightValInteger.Value),
		}
	}

	return evalFloatInfixExpression(operator, leftValFloat, rightValFloat)
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() == object.INTEGER_OBJ {
		value := right.(*object.Integer).Value
		return &object.Integer{Value: -value}
	}

	if right.Type() == object.FLOAT_OBJ {
		value := right.(*object.Float).Value
		return &object.Float{Value: -value}
	}

	return object.NewErrorFormat("unknown operator: -%s", right.Type())
}

func evalIncrementExpression(right object.Object) object.Object {
	if right.Type() == object.INTEGER_OBJ {
		value := right.(*object.Integer).Value
		return &object.Integer{Value: value + 1}
	}

	if right.Type() == object.FLOAT_OBJ {
		value := right.(*object.Float).Value
		return &object.Float{Value: value + 1.0}
	}

	return object.NewErrorFormat("unknown object type %s for operator %s", right.Type(), "++")
}

func evalDecrementExpression(right object.Object) object.Object {
	if right.Type() == object.INTEGER_OBJ {
		value := right.(*object.Integer).Value
		return &object.Integer{Value: value - 1}
	}

	if right.Type() == object.FLOAT_OBJ {
		value := right.(*object.Float).Value
		return &object.Float{Value: value - 1.0}
	}

	return object.NewErrorFormat("unknown object type %s for operator %s", right.Type(), "--")
}
