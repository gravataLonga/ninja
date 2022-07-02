package stdlib

import (
	"ninja/object"
	"ninja/typing"
)

func Rest(args ...object.Object) object.Object {

	err := typing.Check(
		"first", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}

	return object.NULL
}
