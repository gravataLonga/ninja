package stdlib

import (
	"ninja/object"
)

func Last(args ...object.Object) object.Object {

	err := object.Check(
		"last", args,
		object.ExactArgs(1),
		object.WithTypes(object.ARRAY_OBJ),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return object.NULL
}
