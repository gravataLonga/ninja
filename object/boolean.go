package object

import (
	"fmt"
)

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BooleanObj }
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

func (s *Boolean) Call(method string, args ...Object) Object {
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

		return &String{Value: BooleanObj}
	}
	return NewErrorFormat("method %s not exists on string object.", method)
}
