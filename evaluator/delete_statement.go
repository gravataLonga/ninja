package evaluator

import (
	"ninja/ast"
	"ninja/object"
)

func evalDelete(left ast.Expression, index object.Object, env *object.Environment) object.Object {

	ident, ok := left.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("DeleteStatement.left must be a identifier. Got: %T", left)
	}

	value, ok := env.Get(ident.Value)
	if !ok {
		return object.NewErrorFormat("DeleteStatement.left %s identifier not found.", ident.Value)
	}

	switch value.(type) {
	case *object.Array:
		arr, _ := value.(*object.Array)
		indexInteger, ok := index.(*object.Integer)
		if !ok {
			return object.NewErrorFormat("DeleteStatement.index must be a Integer. Got: %T", index)
		}
		arr.Elements = removeIndexFromArray(arr.Elements, indexInteger.Value)
		env.Set(ident.Value, arr)
	case *object.Hash:
		hash, _ := value.(*object.Hash)
		hashable, ok := index.(object.Hashable)
		if !ok {
			return object.NewErrorFormat("DeleteStatement.index must be a Hashable. Got: %T", index)
		}
		delete(hash.Pairs, hashable.HashKey())

		env.Set(ident.Value, hash)
	default:
		return object.NewErrorFormat("DeleteStatement.left only work with array or hash object. Got: %T", value)
	}

	return nil
}

// removeIndexFromArray is slow operations, we need better way?
func removeIndexFromArray(slice []object.Object, s int64) []object.Object {

	copy(slice[s:], slice[s+1:])      // Shift a[i+1:] left one index.
	slice[len(slice)-1] = object.NULL // Erase last element (write zero value).
	slice = slice[:len(slice)-1]      // Truncate slice.

	return slice
}
