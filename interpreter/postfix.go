package interpreter

import (
	"errors"
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func postfixExpression(v *ast.PostfixExpression, obj object.Object) object.Object {
	switch obj.Type() {
	case object.INTEGER_OBJ:
		obj, err := postfixIntegerExpression(v.Operator, obj)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case object.FLOAT_OBJ:
		obj, err := postfixFloatExpression(v.Operator, obj)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	}
	return nil
}

func postfixIntegerExpression(operator string, obj object.Object) (object.Object, error) {
	value := obj.(*object.Integer).Value
	switch operator {
	case "++":
		value = value + 1
		return &object.Integer{Value: value}, nil
	case "--":
		value = value - 1
		return &object.Integer{Value: value}, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s%s", obj.Type(), operator))
}

func postfixFloatExpression(operator string, obj object.Object) (object.Object, error) {
	value := obj.(*object.Float).Value
	switch operator {
	case "++":
		value = value + 1
		return &object.Float{Value: value}, nil
	case "--":
		value = value - 1
		return &object.Float{Value: value}, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s%s", obj.Type(), operator))
}
