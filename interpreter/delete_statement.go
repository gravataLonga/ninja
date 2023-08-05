package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitDelete(v *ast.DeleteStatement) (result object.Object) {
	ident, ok := v.Left.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("DeleteStatement.left must be a identifier. Got: %T", v.Left)
	}

	value, ok := i.env.Get(ident.Value)
	if !ok {
		return object.NewErrorFormat("DeleteStatement.left %s identifier not found.", ident.Value)
	}

	index := i.evaluate(v.Index)

	switch value.(type) {
	case *object.Array:
		arr, _ := value.(*object.Array)
		if !object.IsInteger(index) {
			return object.NewErrorFormat("DeleteStatement.index must be a Integer. Got: %T", index)
		}
		index, _ := index.(*object.Integer)
		arr.Elements = removeIndexFromArray(arr.Elements, index.Value)
		i.env.Set(ident.Value, arr)
	case *object.Hash:
		hash, _ := value.(*object.Hash)
		hashable, ok := index.(object.Hashable)
		if !ok {
			return object.NewErrorFormat("DeleteStatement.index must be a Hashable. Got: %T", index)
		}
		delete(hash.Pairs, hashable.HashKey())

		i.env.Set(ident.Value, hash)
	default:
		return object.NewErrorFormat("DeleteStatement.left only work with array or hash object. Got: %T", value)
	}

	return nil
}
