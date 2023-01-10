package interpreter

import (
	"errors"
	"fmt"
	"github.com/gravataLonga/ninja/object"
)

func indexExpression(left, index object.Object) object.Object {
	switch left.Type() {
	case object.ARRAY_OBJ:
		obj, err := indexArrayExpression(left.(*object.Array), index)
		if err != nil {
			return object.NewErrorFormat(err.Error())
		}
		return obj
	case object.HASH_OBJ:
		obj, err := indexHashExpression(left.(*object.Hash), index)
		if err != nil {
			return object.NewErrorFormat(err.Error())
		}
		return obj
	case object.STRING_OBJ:
		obj, err := indexStringExpression(left.(*object.String), index)
		if err != nil {
			return object.NewErrorFormat(err.Error())
		}
		return obj
	}

	return object.NewErrorFormat("index operator not supported: %s", left.Type())
}

func indexArrayExpression(left *object.Array, index object.Object) (object.Object, error) {
	if _, ok := index.(*object.Integer); !ok {
		return nil, errors.New(fmt.Sprintf("index operator not supported: %s", left.Type()))
	}
	max := len(left.Elements)
	idx := index.(*object.Integer).Value

	if idx < 0 || idx >= int64(max) {
		return object.NULL, nil
	}

	return left.Elements[idx], nil
}

func indexHashExpression(left *object.Hash, index object.Object) (object.Object, error) {
	key, ok := index.(object.Hashable)
	if !ok {
		return nil, errors.New(fmt.Sprintf("unusable as hash key: %s", index.Type()))
	}

	pair, ok := left.Pairs[key.HashKey()]
	if !ok {
		return object.NULL, nil
	}

	return pair.Value, nil
}

func indexStringExpression(left *object.String, index object.Object) (object.Object, error) {
	idx, ok := index.(*object.Integer)
	if !ok {
		return nil, errors.New(fmt.Sprintf("index isnt integer: %s", index.Type()))
	}

	rn := []rune(left.Value)

	if idx.Value < 0 || int64(len(rn)) <= idx.Value {
		return object.NULL, nil
	}

	return &object.String{
		Value: string([]rune(left.Value)[idx.Value]),
	}, nil
}
