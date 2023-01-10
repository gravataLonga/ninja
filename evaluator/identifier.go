package evaluator

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/stdlib"
)

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {

	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := stdlib.Builtins[node.Value]; ok {
		return builtin
	}

	if globalVariable, ok := object.GlobalEnvironment.Get(node.Value); ok {
		return globalVariable
	}

	return object.NewErrorFormat("identifier not found: %s %s", node.Value, node.Token)
}

func evalAssignStatement(node *ast.AssignStatement, env *object.Environment) object.Object {
	switch node.Name.(type) {
	case *ast.Identifier:
		return evalAssignIdentifier(node, env)
	case *ast.IndexExpression:
		return evalAssignIndexIdentifier(node, env)
		break
	default:
		return object.NewErrorFormat("node.Name is not type of identifier. Got %T", node.Name)
	}

	return nil
}

func evalAssignIdentifier(node *ast.AssignStatement, env *object.Environment) object.Object {
	identifier, ok := node.Name.(*ast.Identifier)

	// Unecessary check..
	if !ok {
		return object.NewErrorFormat("node.Name is not type of identifier. Got %T %s", node.Name, node.Token)
	}

	_, ok = env.Get(identifier.Value)
	if !ok {
		return object.NewErrorFormat("identifier not found: %s %s", identifier.Value, node.Token)
	}

	val := Eval(node.Value, env)
	if object.IsError(val) {
		return val
	}
	env.Set(identifier.Value, val)
	return nil
}

func evalAssignIndexIdentifier(node *ast.AssignStatement, env *object.Environment) object.Object {
	indexIdentifier, ok := node.Name.(*ast.IndexExpression)

	if !ok {
		// Unecessary check..
		if !ok {
			return object.NewErrorFormat("node.Name is not type of IndexExpression. Got %T", node.Name)
		}
	}

	objIdentifier, _ := env.Get(indexIdentifier.Left.String())
	value := Eval(node.Value, env)

	switch objIdentifier.(type) {
	case *object.Hash:
		hashObject, _ := objIdentifier.(*object.Hash)

		objIndex := Eval(indexIdentifier.Index, env)
		h, ok := objIndex.(object.Hashable)
		if !ok {
			return object.NewErrorFormat("expected index to be hashable")
		}
		hashObject.Pairs[h.HashKey()] = object.HashPair{Key: objIndex, Value: value}
	case *object.Array:
		arrayObject, _ := objIdentifier.(*object.Array)

		objIndex := Eval(indexIdentifier.Index, env)
		objectIndexInteger, ok := objIndex.(*object.Integer)

		if !ok {
			return object.NewErrorFormat("node.Index is not type of Integer. Got %T", objIndex)
		}

		if objectIndexInteger.Value < 0 {
			return object.NewErrorFormat("index out of range, got %d not positive index", objectIndexInteger.Value)
		}

		l := int64(len(arrayObject.Elements))
		if l == objectIndexInteger.Value {
			arrayObject.Elements = append(arrayObject.Elements, value)
			return nil
		}

		if l < objectIndexInteger.Value {
			return object.NewErrorFormat("index out of range, got %d but array has only %d elements", objectIndexInteger.Value, l)
		}

		arrayObject.Elements[objectIndexInteger.Value] = value
	}

	return nil
}
