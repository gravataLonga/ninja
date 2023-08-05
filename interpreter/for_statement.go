package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitFor(v *ast.ForStatement) (result object.Object) {
	if v.InitialCondition != nil {
		i.execute(v.InitialCondition)
	}

	// @todo test this better
	i.EnterLoop()
	condition := i.interpreterConditionForLoop(v.Condition)
	for object.IsTruthy(condition) {

		result = i.execute(v.Body)
		if result != nil {
			if object.IsReturn(result) {
				i.ExitLoop()
				return
			}

			if result.Type() == object.BREAK_VALUE_OBJ {
				i.ExitLoop()
				return nil
			}

			if object.IsError(result) {
				i.ExitLoop()
				return
			}
		}
		if v.Iteration != nil {
			i.execute(v.Iteration)
		}
		condition = i.interpreterConditionForLoop(v.Condition)
	}
	// @todo test this better
	i.ExitLoop()

	return result
}

// interpreterConditionForLoop @todo check better way.
func (i *Interpreter) interpreterConditionForLoop(v ast.Expression) object.Object {
	if v == nil {
		return object.TRUE
	}
	return i.evaluate(v)
}
