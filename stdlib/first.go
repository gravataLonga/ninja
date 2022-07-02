package stdlib

import (
	"ninja/object"
	"ninja/typing"
)

func First(args ...object.Object) object.Object {

	err := typing.Check(
		"first", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	);

	if err != nil {
		return object.NewError(err.Error())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return object.NULL
}
