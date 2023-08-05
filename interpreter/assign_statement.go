package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitVarStmt(v *ast.VarStatement) (result object.Object) {
	result = i.evaluate(v.Value)
	if object.IsError(result) {
		return
	}
	i.env.Set(v.Name.Value, result)
	return nil
}

func (i *Interpreter) VisitAssignStmt(v *ast.AssignStatement) (result object.Object) {
	ident, ok := v.Left.(*ast.Identifier)
	if ok {
		left := ident.Value
		i.env.Set(left, i.evaluate(v.Right))
		return nil
	}

	expr, ok := v.Left.(*ast.ExpressionStatement)
	if !ok {
		return nil
	}

	idx, ok := expr.Expression.(*ast.IndexExpression)
	if !ok {
		return nil
	}

	ident, ok = idx.Left.(*ast.Identifier)
	if !ok {
		return nil
	}

	left := ident.Value

	obj, ok := i.env.Get(left)
	if !ok {
		return nil
	}

	if object.IsArray(obj) {
		arr, _ := obj.(*object.Array)
		index := i.evaluate(idx.Index)
		indexIntegerObject, ok := index.(*object.Integer)
		if !ok {
			return nil
		}

		indexInteger := int(indexIntegerObject.Value)
		lenElements := len(arr.Elements)

		if indexInteger <= -1 {
			return object.NewErrorFormat("index out of range, got %d not positive index", indexInteger)
		}

		if lenElements < indexInteger {
			return object.NewErrorFormat("index out of range, got %d but array has only %d elements", indexInteger, lenElements)
		}

		if indexInteger > lenElements-1 {
			lenElements = lenElements + 1
		}

		elements := make([]object.Object, lenElements)
		copy(elements, arr.Elements)
		elements[indexInteger] = i.evaluate(v.Right)
		arr.Elements = elements
		i.env.Set(left, arr)
	}

	if object.IsHash(obj) {
		hashObject, _ := obj.(*object.Hash)

		objIndex := i.evaluate(idx.Index)
		h, ok := objIndex.(object.Hashable)
		if !ok {
			return object.NewErrorFormat("expected index to be hashable")
		}
		hashObject.Pairs[h.HashKey()] = object.HashPair{Key: objIndex, Value: i.evaluate(v.Right)}
	}

	return nil
}

func (i *Interpreter) VisitIdentExpr(v *ast.Identifier) (result object.Object) {
	value, ok := i.env.Get(v.Value)
	if !ok {
		return object.NewErrorFormat("identifier not found: %s %s", v.Value, v.Token)
		// return object.NULL
	}
	return value
}
