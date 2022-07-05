package stdlib

import (
	"ninja/object"
)

// First get item from array object
func First(args ...object.Object) object.Object {

	err := object.Check(
		"first", args,
		object.ExactArgs(1),
		object.WithTypes(object.ARRAY_OBJ),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	cloneable, ok := args[0].(object.Cloneable)
	if !ok {
		return object.NewErrorFormat("object isn't cloneable.")
	}

	arrClone := cloneable.Clone()
	arr := arrClone.(*object.Array)

	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return object.NULL
}
