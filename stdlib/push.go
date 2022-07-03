package stdlib

import (
	"ninja/object"
)

func Push(args ...object.Object) object.Object {

	err := object.Check(
		"push", args,
		object.ExactArgs(2),
		object.WithTypes(object.ARRAY_OBJ),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	newElements := make([]object.Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}
