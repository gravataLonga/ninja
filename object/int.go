package object

import (
	"fmt"
	"ninja/ast"
	"strconv"
)

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (i *Integer) Compare(right Object) int8 {
	if obj, ok := right.(*Integer); ok {
		switch {
		case i.Value < obj.Value:
			return -1
		case i.Value > obj.Value:
			return 1
		default:
			return 0
		}
	}
	return -1
}

func (s *Integer) Call(objectCall *ast.ObjectCall, method string, env *Environment, args ...Object) Object {
	switch method {
	case "type":
		err := Check(
			"int.type",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}
		return &String{Value: INTEGER_OBJ}
	case "string":
		err := Check(
			"int.string",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}
		return &String{Value: strconv.FormatInt(s.Value, 10)}
	case "float":
		err := Check(
			"int.float",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}
		return &Float{Value: float64(s.Value)}
	case "abs":
		err := Check(
			"int.abs",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}
		var absT int64 = s.Value
		if s.Value < 0 {
			absT = -s.Value
		}
		return &Integer{Value: absT}
	}
	return NewErrorFormat("method %s not exists on integer object.", method)

}
