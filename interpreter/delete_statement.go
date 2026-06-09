package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitDelete(v *ast.DeleteStatement) (result object.Object) {
	ident, ok := v.Left.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("DeleteStatement.left %s identifier not found. %s", v.Left, v.Token.HumanLocation())
	}

	value, ok := i.env.Get(ident.Value)
	if !ok {
		return object.NewErrorFormat("DeleteStatement.left %s identifier not found. %s", ident.Value, v.Token.HumanLocation())
	}

	index := i.evaluate(v.Index)

	switch value.(type) {
	case *object.Array:
		arr, _ := value.(*object.Array)
		if !object.IsInteger(index) {
			return object.NewErrorFormat("DeleteStatement.index must be a Integer. Got: %T %s", index, v.Token.HumanLocation())
		}
		index, _ := index.(*object.Integer)
		total := int64(len(arr.Elements))
		if total < index.Value {
			return object.NewErrorFormat("DeleteStatement.index must be equal or less than the total of the items. Got: %T %s", index, v.Token.HumanLocation())
		}
		if index.Value < 0 {
			return object.NewErrorFormat("DeleteStatement.index must be equal or greater than the 0. Got: %T %s", index, v.Token.HumanLocation())
		}
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
		return object.NewErrorFormat("DeleteStatement.left only work with array or hash object. Got: %T %s", value, v.Token.HumanLocation())
	}

	return nil
}
