package evaluator

import (
	"ninja/ast"
	"ninja/object"
)

func evalForStatement(
	node *ast.ForStatement,
	env *object.Environment,
) object.Object {

	var result object.Object

	if node.InitialCondition != nil {
		Eval(node.InitialCondition, env)
	}

	for {
		condition := Eval(node.Condition, env)

		conditionIsTrue, ok := condition.(*object.Boolean)
		if !ok {
			break
		}

		if !conditionIsTrue.Value {
			break
		}

		if node.Iteration != nil {
			Eval(node.Iteration, env)
		}

		result = Eval(node.Body, env)
	}

	return result
}
