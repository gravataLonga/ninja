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
	}
	return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func infixPlusExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.INTEGER_OBJ:
		switch right.Type() {
		case object.INTEGER_OBJ:
			return &object.Integer{Value: left.(*object.Integer).Value + right.(*object.Integer).Value}, nil
		case object.FLOAT_OBJ:
			v := float64(left.(*object.Integer).Value)
			return &object.Float{Value: v + right.(*object.Float).Value}, nil
		}
	case object.FLOAT_OBJ:
		switch right.Type() {
		case object.FLOAT_OBJ:
			return &object.Float{Value: left.(*object.Float).Value + right.(*object.Float).Value}, nil
		case object.INTEGER_OBJ:
			v := float64(right.(*object.Integer).Value)
			return &object.Float{Value: left.(*object.Float).Value + v}, nil
		}

	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "+", right.Type()))
}

func infixMinusExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.INTEGER_OBJ:
		switch right.Type() {
		case object.INTEGER_OBJ:
			return &object.Integer{Value: left.(*object.Integer).Value - right.(*object.Integer).Value}, nil
		case object.FLOAT_OBJ:
			left := float64(left.(*object.Integer).Value)
			return &object.Float{Value: left - right.(*object.Float).Value}, nil
		}
	case object.FLOAT_OBJ:
		switch right.Type() {
		case object.FLOAT_OBJ:
			return &object.Float{Value: left.(*object.Float).Value - right.(*object.Float).Value}, nil
		case object.INTEGER_OBJ:
			right := float64(right.(*object.Integer).Value)
			return &object.Float{Value: left.(*object.Float).Value - right}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "-", right.Type()))
}

func infixMulExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.INTEGER_OBJ:
		switch right.Type() {
		case object.INTEGER_OBJ:
			return &object.Integer{Value: left.(*object.Integer).Value * right.(*object.Integer).Value}, nil
		case object.FLOAT_OBJ:
			left := float64(left.(*object.Integer).Value)
			return &object.Float{Value: left * right.(*object.Float).Value}, nil

		}
	case object.FLOAT_OBJ:
		switch right.Type() {
		case object.FLOAT_OBJ:
			return &object.Float{Value: left.(*object.Float).Value * right.(*object.Float).Value}, nil
		case object.INTEGER_OBJ:
			right := float64(right.(*object.Integer).Value)
			return &object.Float{Value: left.(*object.Float).Value * right}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "*", right.Type()))
}

func infixDivExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.INTEGER_OBJ:
		switch right.Type() {
		case object.INTEGER_OBJ:
			return &object.Integer{Value: left.(*object.Integer).Value / right.(*object.Integer).Value}, nil
		case object.FLOAT_OBJ:
			left := float64(left.(*object.Integer).Value)
			return &object.Float{Value: left / right.(*object.Float).Value}, nil
		}
	case object.FLOAT_OBJ:
		switch right.Type() {
		case object.FLOAT_OBJ:
			return &object.Float{Value: left.(*object.Float).Value / right.(*object.Float).Value}, nil
		case object.INTEGER_OBJ:
			right := float64(right.(*object.Integer).Value)
			return &object.Float{Value: left.(*object.Float).Value / right}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "/", right.Type()))
}

func infixModExpression(left, right object.Object) (object.Object, error) {
	switch left.Type() {
	case object.INTEGER_OBJ:
		switch right.Type() {
		case object.INTEGER_OBJ:
			return &object.Integer{Value: left.(*object.Integer).Value % right.(*object.Integer).Value}, nil
		case object.FLOAT_OBJ:
			left := float64(left.(*object.Integer).Value)
			return &object.Float{Value: math.Mod(left, right.(*object.Float).Value)}, nil
		}
	case object.FLOAT_OBJ:
		switch right.Type() {
		case object.FLOAT_OBJ:
			return &object.Float{Value: math.Mod(left.(*object.Float).Value, right.(*object.Float).Value)}, nil
		case object.INTEGER_OBJ:
			right := float64(right.(*object.Integer).Value)
			return &object.Float{Value: math.Mod(left.(*object.Float).Value, right)}, nil
		}

	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), "*", right.Type()))
}
