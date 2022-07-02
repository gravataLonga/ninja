package stdlib

import (
	"ninja/object"
	"ninja/typing"
)

func Args(args ...object.Object) object.Object {

	err := typing.Check(
		"args", args,
		typing.ExactArgs(0),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	elements := make([]object.Object, len(object.Arguments))
	for i, arg := range object.Arguments {
		elements[i] = &object.String{Value: arg}
	}
	return &object.Array{Elements: elements}
}
