package interpreter

import (
	"errors"
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func prefixExpression(v *ast.PrefixExpression, obj object.Object) object.Object {
	switch obj.Type() {
	case object.STRING_OBJ:
		obj, err := prefixStringExpression(v.Operator, obj)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case object.INTEGER_OBJ:
		obj, err := prefixIntegerExpression(v.Operator, obj)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case object.FLOAT_OBJ:
		obj, err := prefixFloatExpression(v.Operator, obj)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case object.BOOLEAN_OBJ:
		obj, err := prefixBooleanExpression(v.Operator, obj)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case object.ARRAY_OBJ:
		obj, err := prefixArrayExpression(v.Operator, obj)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	case object.HASH_OBJ:
		obj, err := prefixHashExpression(v.Operator, obj)
		if err != nil {
			return object.NewErrorFormat("%s %s", err, v.Token)
		}
		return obj
	}
	return object.NewErrorFormat("unknown operator: %s%s %s", v.Operator, obj.Type(), v.Token)
}

func prefixStringExpression(operator string, obj object.Object) (object.Object, error) {
	//value := obj.(*object.String).Right
	switch operator {
	case "!":
		return &object.Boolean{Value: !object.IsTruthy(obj)}, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s%s", operator, obj.Type()))
}

func prefixIntegerExpression(operator string, obj object.Object) (object.Object, error) {
	value := obj.(*object.Integer).Value
	switch operator {
	case "-":
		return &object.Integer{Value: -value}, nil
	case "++":
		value = value + 1
		return &object.Integer{Value: value}, nil
	case "--":
		value = value - 1
		return &object.Integer{Value: value}, nil
	case "!":
		return &object.Boolean{Value: !object.IsTruthy(obj)}, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s%s", operator, obj.Type()))
}

func prefixFloatExpression(operator string, obj object.Object) (object.Object, error) {
	value := obj.(*object.Float).Value
	switch operator {
	case "-":
		return &object.Float{Value: -value}, nil
	case "++":
		value = value + 1.0
		return &object.Float{Value: value}, nil
	case "--":
		value = value - 1.0
		return &object.Float{Value: value}, nil
	case "!":
		return &object.Boolean{Value: !object.IsTruthy(obj)}, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s%s", operator, obj.Type()))
}

func prefixBooleanExpression(operator string, obj object.Object) (object.Object, error) {
	// value := obj.(*object.Boolean).Right
	switch operator {
	case "!":
		return &object.Boolean{Value: !object.IsTruthy(obj)}, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s%s", operator, obj.Type()))
}

func prefixArrayExpression(operator string, obj object.Object) (object.Object, error) {
	// value := obj.(*object.Boolean).Right
	switch operator {
	case "!":
		return &object.Boolean{Value: !object.IsTruthy(obj)}, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s%s", operator, obj.Type()))
}

func prefixHashExpression(operator string, obj object.Object) (object.Object, error) {
	// value := obj.(*object.Boolean).Right
	switch operator {
	case "!":
		return &object.Boolean{Value: !object.IsTruthy(obj)}, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown operator: %s%s", operator, obj.Type()))
}