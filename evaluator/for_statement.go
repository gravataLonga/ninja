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

	condition := evalConditionForLoop(node.Condition, env)
	for object.IsTruthy(condition) {

		result = Eval(node.Body, env)
		if result != nil {
			if result.Type() == object.RETURN_VALUE_OBJ {
				return result
			}

			if result.Type() == object.BREAK_VALUE_OBJ {
				return nil
			}
		}
		if node.Iteration != nil {
			Eval(node.Iteration, env)
		}
		condition = evalConditionForLoop(node.Condition, env)
	}

	return result
}

func evalConditionForLoop(nodeCondition ast.Expression, env *object.Environment) object.Object {
	if nodeCondition != nil {
		return Eval(nodeCondition, env)
	}
	return object.TRUE
}
