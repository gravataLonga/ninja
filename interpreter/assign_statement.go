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
		if depth, ok := i.locals[ident]; ok {
			i.env.SetAt(depth, ident.Value, i.evaluate(v.Right))
			return nil
		}

		rs := i.evaluate(v.Right)
		if object.IsError(rs) {
			return rs
		}
		i.env.Assign(left, rs)
		return nil
	}

	expr, ok := v.Left.(*ast.ExpressionStatement)
	if !ok {
		return nil
	}

	if dot, ok := expr.Expression.(*ast.Dot); ok {
		obj := i.evaluate(dot.Object)

		hash, ok := obj.(*object.Hash)
		if !ok {
			return object.NewErrorFormat("cannot assign property %q on %s %s", dot.Right.Value, obj.Type(), v.Token.HumanLocation())
		}

		key := &object.String{Value: dot.Right.Value}
		hash.Pairs[key.HashKey()] = object.HashPair{Key: key, Value: i.evaluate(v.Right)}
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

	if !object.IsArray(obj) && !object.IsHash(obj) {
		return object.NewErrorFormat("cannot assign property %q on %s", ident.Value, obj.Type())
	}

	if object.IsArray(obj) {
		arr, _ := obj.(*object.Array)
		index := i.evaluate(idx.Index)
		if object.IsError(index) {
			return index
		}

		indexIntegerObject, ok := index.(*object.Integer)
		if !ok {
			return nil
		}

		indexInteger := int(indexIntegerObject.Value)
		lenElements := len(arr.Elements)

		if indexInteger <= -1 {
			return object.NewErrorFormat("index out of range, got %d not positive index %s", indexInteger, v.Token.HumanLocation())
		}

		if lenElements < indexInteger {
			return object.NewErrorFormat("index out of range, got %d but array has only %d elements %s", indexInteger, lenElements, v.Token.HumanLocation())
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
		if object.IsError(objIndex) {
			return objIndex
		}
		h, ok := objIndex.(object.Hashable)
		if !ok {
			return object.NewErrorFormat("expected index to be hashable")
		}
		hashObject.Pairs[h.HashKey()] = object.HashPair{Key: objIndex, Value: i.evaluate(v.Right)}
	}

	return nil
}

func (i *Interpreter) VisitIdentExpr(v *ast.Identifier) (result object.Object) {
	if depth, ok := i.locals[v]; ok {
		return i.env.GetAt(depth, v.Value)
	}

	value, ok := i.env.Get(v.Value)
	if !ok {
		return object.NewErrorFormat("identifier not found: %s %s", v.Value, v.Token.HumanLocation())
		// return object.NULL
	}
	return value
}
