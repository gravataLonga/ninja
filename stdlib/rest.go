package stdlib

import "ninja/object"

func Rest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewErrorFormat("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return object.NewErrorFormat("argument to `rest` must be ARRAY, got %s",
			args[0].Type())
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
