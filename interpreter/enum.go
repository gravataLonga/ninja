package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitEnum(v *ast.EnumStatement) (result object.Object) {
	enum := &object.Enum{Branches: map[string]object.Object{}}
	for o, v := range v.Branches {
		enum.Branches[o] = i.evaluate(v)
	}

	ident, ok := v.Identifier.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("expected identifier. got: %s", v.Identifier)
	}

	i.env.Set(ident.Value, enum)
	return enum
}

func (i *Interpreter) VisitScopeOperatorExpression(v *ast.ScopeOperatorExpression) (result object.Object) {
	access, ok := v.AccessIdentifier.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("expected access identifier. got: %s", v.AccessIdentifier)
	}

	property, ok := v.PropertyIdentifier.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("expected property identifier. got: %s", v.PropertyIdentifier)
	}

	obj, ok := i.env.Get(access.Value)
	if !ok {
		return object.NewErrorFormat("identifier not found: " + access.Value)
	}

	enum, ok := obj.(*object.Enum)
	if !ok {
		return object.NewErrorFormat("identifier must be accessible with :: got: %s", v)
	}

	brancheValue, ok := enum.Branches[property.Value]
	if !ok {
		return object.NewErrorFormat("identifier %s don't exists on enum object", property.Value)
	}

	return brancheValue
}
