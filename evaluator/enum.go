package evaluator

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func evalEnumStatement(node *ast.EnumStatement, env *object.Environment) object.Object {
	enum := &object.Enum{Branches: map[string]object.Object{}}
	for i, v := range node.Branches {
		enum.Branches[i] = Eval(v, env)
	}

	ident, ok := node.Identifier.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("expected identifier. got: %s", node.Identifier)
	}

	env.Set(ident.Value, enum)
	return enum
}

func evalScopeOperatorExpression(node *ast.ScopeOperatorExpression, env *object.Environment) object.Object {
	access, ok := node.AccessIdentifier.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("expected access identifier. got: %s", node.AccessIdentifier)
	}

	property, ok := node.PropertyIdentifier.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("expected property identifier. got: %s", node.PropertyIdentifier)
	}

	v, ok := env.Get(access.Value)
	if !ok {
		return object.NewErrorFormat("identifier not found: " + access.Value)
	}

	enum, ok := v.(*object.Enum)
	if !ok {
		return object.NewErrorFormat("identifier must be accessible with :: got: %s", v)
	}

	brancheValue, ok := enum.Branches[property.Value]
	if !ok {
		return object.NewErrorFormat("identifier %s don't exists on enum object", property.Value)
	}

	return brancheValue
}
