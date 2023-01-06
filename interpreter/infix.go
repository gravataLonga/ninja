package interpreter

import (
	"errors"
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
	"math"
)

func infixExpression(v *ast.InfixExpression, operator string, left object.Object, right object.Object) object.Object {
	switch operator {
	case "+":
		obj, err := infixPlusExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "-":
		obj, err := infixMinusExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "*":
		obj, err := infixMulExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "/":
		obj, err := infixDivExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "%":
		obj, err := infixModExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "**":
		obj, err := infixPowExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "|":
		obj, err := infixOrBitExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "&":
		obj, err := infixAndBitExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "^":
		obj, err := infixXorExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "<<":
		obj, err := infixShiftLeftExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case ">>":
		obj, err := infixShiftRightExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "==":
		obj, err := infixEqualExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "!=":
		obj, err := infixNotEqualExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "<":
		obj, err := infixLessExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case ">":
		obj, err := infixGreaterExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "<=":
		obj, err := infixLessOrEqualExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case ">=":
		obj, err := infixGreaterOrEqualExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "&&":
		obj, err := infixAndExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case "||":
		obj, err := infixOrExpression(left, right)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	}
	return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func infixPlusExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			return &object.Integer{Value: left.(*object.Integer).Value + right.(*object.Integer).Value}, nil
		case object.FloatObj:
			v := float64(left.(*object.Integer).Value)
			return &object.Float{Value: v + right.(*object.Float).Value}, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			return &object.Float{Value: left.(*object.Float).Value + right.(*object.Float).Value}, nil
		case object.IntegerObj:
			v := float64(right.(*object.Integer).Value)
			return &object.Float{Value: left.(*object.Float).Value + v}, nil
		}
	case object.StringObj:
		if right.Type() == object.StringObj {
			return &object.String{Value: left.(*object.String).Value + right.(*object.String).Value}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "+", right.Type()))
}

func infixMinusExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			return &object.Integer{Value: left.(*object.Integer).Value - right.(*object.Integer).Value}, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			return &object.Float{Value: left - right.(*object.Float).Value}, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			return &object.Float{Value: left.(*object.Float).Value - right.(*object.Float).Value}, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			return &object.Float{Value: left.(*object.Float).Value - right}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "-", right.Type()))
}

func infixMulExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			return &object.Integer{Value: left.(*object.Integer).Value * right.(*object.Integer).Value}, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			return &object.Float{Value: left * right.(*object.Float).Value}, nil

		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			return &object.Float{Value: left.(*object.Float).Value * right.(*object.Float).Value}, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			return &object.Float{Value: left.(*object.Float).Value * right}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "*", right.Type()))
}

func infixDivExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			return &object.Integer{Value: left.(*object.Integer).Value / right.(*object.Integer).Value}, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			return &object.Float{Value: left / right.(*object.Float).Value}, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			return &object.Float{Value: left.(*object.Float).Value / right.(*object.Float).Value}, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			return &object.Float{Value: left.(*object.Float).Value / right}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "/", right.Type()))
}

func infixModExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			return &object.Integer{Value: left.(*object.Integer).Value % right.(*object.Integer).Value}, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			return &object.Float{Value: math.Mod(left, right.(*object.Float).Value)}, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			return &object.Float{Value: math.Mod(left.(*object.Float).Value, right.(*object.Float).Value)}, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			return &object.Float{Value: math.Mod(left.(*object.Float).Value, right)}, nil
		}

	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "%", right.Type()))
}

func infixPowExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			result := math.Pow(float64(left.(*object.Integer).Value), float64(right.(*object.Integer).Value))
			return &object.Integer{Value: int64(result)}, nil
		case object.FloatObj:
			return &object.Float{Value: math.Pow(float64(left.(*object.Integer).Value), right.(*object.Float).Value)}, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			return &object.Float{Value: math.Pow(left.(*object.Float).Value, right.(*object.Float).Value)}, nil
		case object.IntegerObj:
			return &object.Float{Value: math.Pow(left.(*object.Float).Value, float64(right.(*object.Integer).Value))}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "**", right.Type()))
}

func infixOrBitExpression(left, right object.Object) (object.Object, error) {
	if left.Type() != object.IntegerObj || right.Type() != object.IntegerObj {
		return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "|", right.Type()))
	}

	return &object.Integer{Value: left.(*object.Integer).Value | right.(*object.Integer).Value}, nil
}

func infixAndBitExpression(left, right object.Object) (object.Object, error) {
	if left.Type() != object.IntegerObj || right.Type() != object.IntegerObj {
		return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "&", right.Type()))
	}

	return &object.Integer{Value: left.(*object.Integer).Value & right.(*object.Integer).Value}, nil
}

func infixXorExpression(left, right object.Object) (object.Object, error) {
	if left.Type() != object.IntegerObj || right.Type() != object.IntegerObj {
		return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "^", right.Type()))
	}

	return &object.Integer{Value: left.(*object.Integer).Value ^ right.(*object.Integer).Value}, nil
}

func infixEqualExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			if left.(*object.Integer).Value == right.(*object.Integer).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			if left == right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		default:
			return object.FALSE, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			if left.(*object.Float).Value == right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			if left.(*object.Float).Value == right {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		default:
			return object.FALSE, nil
		}
	case object.StringObj:
		switch right.Type() {
		case object.StringObj:
			if left.(*object.String).Value == right.(*object.String).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		default:
			return object.FALSE, nil
		}
	case object.BooleanObj:
		switch right.Type() {
		case object.BooleanObj:
			if left.(*object.Boolean).Value == right.(*object.Boolean).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		default:
			return object.FALSE, nil
		}
	case object.ArrayObj:
		fallthrough
	case object.HashObj:
		return object.FALSE, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "==", right.Type()))
}

func infixNotEqualExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.StringObj:
		switch right.Type() {
		case object.StringObj:
			if left.(*object.String).Value != right.(*object.String).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		default:
			return object.TRUE, nil
		}
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			if left.(*object.Integer).Value != right.(*object.Integer).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			if left != right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		default:
			return object.TRUE, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			if left.(*object.Float).Value != right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			if left.(*object.Float).Value != right {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		default:
			return object.TRUE, nil

		}
	case object.BooleanObj:
		switch right.Type() {
		case object.BooleanObj:
			if left.(*object.Boolean).Value != right.(*object.Boolean).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		default:
			return object.TRUE, nil
		}
	case object.HashObj:
		fallthrough
	case object.ArrayObj:
		return object.TRUE, nil

	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "!=", right.Type()))
}

func infixLessExpression(left, right object.Object) (object.Object, error) {
	if !object.IsNumber(left) || !object.IsNumber(right) {
		return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "<", right.Type()))
	}

	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			if left.(*object.Integer).Value < right.(*object.Integer).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			if left < right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			if left.(*object.Float).Value < right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			if left.(*object.Float).Value < right {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "<", right.Type()))
}

func infixGreaterExpression(left, right object.Object) (object.Object, error) {
	if !object.IsNumber(left) || !object.IsNumber(right) {
		return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), ">", right.Type()))
	}

	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			if left.(*object.Integer).Value > right.(*object.Integer).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			if left > right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			if left.(*object.Float).Value > right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			if left.(*object.Float).Value > right {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), ">", right.Type()))
}

func infixLessOrEqualExpression(left, right object.Object) (object.Object, error) {
	if !object.IsNumber(left) || !object.IsNumber(right) {
		return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "<=", right.Type()))
	}

	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			if left.(*object.Integer).Value <= right.(*object.Integer).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			if left <= right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			if left.(*object.Float).Value <= right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			if left.(*object.Float).Value <= right {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "<=", right.Type()))
}

func infixGreaterOrEqualExpression(left, right object.Object) (object.Object, error) {
	if !object.IsNumber(left) || !object.IsNumber(right) {
		return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), ">=", right.Type()))
	}

	switch left.Type() {
	case object.IntegerObj:
		switch right.Type() {
		case object.IntegerObj:
			if left.(*object.Integer).Value >= right.(*object.Integer).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.FloatObj:
			left := float64(left.(*object.Integer).Value)
			if left >= right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		}
	case object.FloatObj:
		switch right.Type() {
		case object.FloatObj:
			if left.(*object.Float).Value >= right.(*object.Float).Value {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		case object.IntegerObj:
			right := float64(right.(*object.Integer).Value)
			if left.(*object.Float).Value >= right {
				return object.TRUE, nil
			}
			return object.FALSE, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), ">", right.Type()))
}

func infixAndExpression(left, right object.Object) (object.Object, error) {
	if object.IsTruthy(left) && object.IsTruthy(right) {
		return object.TRUE, nil
	}
	return object.FALSE, nil
}

func infixOrExpression(left, right object.Object) (object.Object, error) {
	if object.IsTruthy(left) || object.IsTruthy(right) {
		return object.TRUE, nil
	}
	return object.FALSE, nil
}

func infixShiftLeftExpression(left, right object.Object) (object.Object, error) {
	left, okLeft := left.(*object.Integer)
	right, okRight := right.(*object.Integer)

	if !okLeft || !okRight {
		return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "<<", right.Type()))
	}

	return &object.Integer{Value: left.(*object.Integer).Value << right.(*object.Integer).Value}, nil
}

func infixShiftRightExpression(left, right object.Object) (object.Object, error) {
	left, okLeft := left.(*object.Integer)
	right, okRight := right.(*object.Integer)

	if !okLeft || !okRight {
		return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), ">>", right.Type()))
	}

	return &object.Integer{Value: left.(*object.Integer).Value >> right.(*object.Integer).Value}, nil
}
