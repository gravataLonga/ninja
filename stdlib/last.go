package stdlib

import "ninja/object"

func Last(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewErrorFormat("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return object.NewErrorFormat("argument to `last` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return object.NULL
}
