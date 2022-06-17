package evaluator

import (
	"ninja/ast"
	"ninja/object"
)

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {

	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return object.NewErrorFormat("identifier not found: " + node.Value)
}

func evalReassignmentVarStatement(node *ast.ReassignmentVarStatement, env *object.Environment) object.Object {
	_, ok := env.Get(node.Name.Value)
	if !ok {
		return object.NewErrorFormat("identifier not found: %s", node.Name.Value)
	}

	val := Eval(node.Value, env)
	if object.IsError(val) {
		return val
	}
	env.Set(node.Name.Value, val)
	return nil
}
