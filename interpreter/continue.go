package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitContinue(v *ast.ContinueStatement) (result object.Object) {
	return &object.Continue{}
}
