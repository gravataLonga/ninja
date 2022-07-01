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

func (s *Integer) Call(objectCall *ast.ObjectCall, method string, env *Environment, args ...Object) Object {
	switch method {
	case "type":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method type not accept any arguments. got: %s", argStr)
		}
		return &String{Value: INTEGER_OBJ}
	case "string":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method string not accept any arguments. got: %s", argStr)
		}
		return &String{Value: strconv.FormatInt(s.Value, 10)}
	case "float":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method float not accept any arguments. got: %s", argStr)
		}
		return &Float{Value: float64(s.Value)}
	case "abs":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method abs not accept any arguments. got: %s", argStr)
		}
		var absT int64 = s.Value
		if s.Value < 0 {
			absT = -s.Value
		}
		return &Integer{Value: absT}
	}
	return NewErrorFormat("method %s not exists on integer object.", method)

}
