package object

import (
	"fmt"
	"ninja/ast"
)

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

func (s *Boolean) Call(objectCall *ast.ObjectCall, method string, env *Environment, args ...Object) Object {
	switch method {
	case "type":
		err := Check(
			"bool.type",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}
		
		return &String{Value: BOOLEAN_OBJ}
	}
	return NewErrorFormat("method %s not exists on string object.", method)
}
