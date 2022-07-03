package stdlib

import (
	"ninja/object"
)

func First(args ...object.Object) object.Object {

	err := object.Check(
		"first", args,
		object.ExactArgs(1),
		object.WithTypes(object.ARRAY_OBJ),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return object.NULL
}
