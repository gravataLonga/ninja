package evaluator

import (
	"ninja/ast"
	"ninja/object"
)

func evalBlockStatement(
	block *ast.BlockStatement,
	env *object.Environment,
) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result == nil {
			continue
		}

		rt := result.Type()
		if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ || rt == object.BREAK_VALUE_OBJ {
			return result
		}
	}
	return result
}
