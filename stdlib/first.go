package stdlib

import "ninja/object"

func First(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewErrorFormat("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return object.NewErrorFormat("argument to `first` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return object.NULL
}
