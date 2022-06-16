package evaluator

import "ninja/object"

func evalFloatOrIntegerInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftValFloat, okLeftFloat := left.(*object.Float)
	rightValFloat, okRightFloat := right.(*object.Float)

	leftValInteger, okLeftInteger := left.(*object.Integer)
	rightValInteger, okRightInteger := right.(*object.Integer)

	if okLeftFloat && okRightFloat {
		return evalFloatInfixExpression(operator, leftValFloat, rightValFloat)
	}

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

	return newError("unknown operator: -%s", right.Type())
}
