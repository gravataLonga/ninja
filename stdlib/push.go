package stdlib

import "ninja/object"

func Push(args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.NewErrorFormat("wrong number of arguments. got=%d, want=2",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return object.NewErrorFormat("argument to `push` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	newElements := make([]object.Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}
